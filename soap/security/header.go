package security

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	xmlpkg "github.com/diegofxm/ubl21-dian/xml"
)

// NOTA: Los templates se cargan desde disco en runtime porque go:embed no soporta paths relativos con ../
const (
	securityHeaderTemplatePath = "ubl21-dian/soap/templates/security/security_header.tmpl"
	signedInfoTemplatePath     = "ubl21-dian/soap/templates/security/signed_info.tmpl"
	toElementTemplatePath      = "ubl21-dian/soap/templates/security/to_element.tmpl"
)

// Header genera el WS-Security header para SOAP
type Header struct {
	privateKey  *rsa.PrivateKey
	certificate *x509.Certificate
	toURL       string
	action      string
	timestamp   time.Time
	// IDs únicos generados para este request
	idTimestamp              string
	idBinarySecurityToken    string
	idSignature              string
	idKeyInfo                string
	idSecurityTokenReference string
	idTo                     string
}

// NewHeader crea un nuevo security header
func NewHeader(certPath, keyPath, toURL, action string) (*Header, error) {
	// Leer certificado y clave privada
	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}

	var certificate *x509.Certificate
	var privateKey *rsa.PrivateKey

	// Decodificar todos los bloques PEM
	rest := certPEM
	for {
		var block *pem.Block
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}

		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse certificate: %w", err)
			}
			if certificate == nil {
				certificate = cert
			}

		case "RSA PRIVATE KEY":
			key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
			}
			privateKey = key

		case "PRIVATE KEY":
			key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse private key: %w", err)
			}
			rsaKey, ok := key.(*rsa.PrivateKey)
			if !ok {
				return nil, errors.New("private key is not RSA")
			}
			privateKey = rsaKey
		}
	}

	if certificate == nil {
		return nil, errors.New("no certificate found in PEM file")
	}
	if privateKey == nil {
		return nil, errors.New("no private key found in PEM file")
	}

	// Generar IDs únicos para este request
	uniqueID := generateUniqueID()
	
	return &Header{
		privateKey:               privateKey,
		certificate:              certificate,
		toURL:                    toURL,
		action:                   action,
		timestamp:                time.Now().UTC(),
		idTimestamp:              fmt.Sprintf("TS-%s", uniqueID),
		idBinarySecurityToken:    fmt.Sprintf("X509-%s", uniqueID),
		idSignature:              fmt.Sprintf("SIG-%s", uniqueID),
		idKeyInfo:                fmt.Sprintf("KI-%s", uniqueID),
		idSecurityTokenReference: fmt.Sprintf("STR-%s", uniqueID),
		idTo:                     fmt.Sprintf("ID-%s", uniqueID),
	}, nil
}

// Generate genera el XML del security header usando templates
func (sh *Header) Generate() (string, error) {
	// 1. Timestamp (TimeToLive de 60000 segundos como PHP)
	created := sh.timestamp.Format("2006-01-02T15:04:05Z")
	expires := sh.timestamp.Add(60000 * time.Second).Format("2006-01-02T15:04:05Z")

	// 2. Certificado en base64
	certB64 := base64.StdEncoding.EncodeToString(sh.certificate.Raw)

	// 3. Calcular digest del wsa:To usando template
	tmplTo, err := template.ParseFiles(toElementTemplatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse to template: %w", err)
	}

	var toBuffer bytes.Buffer
	if err := tmplTo.Execute(&toBuffer, map[string]string{
		"ToID":  sh.idTo,
		"ToURL": sh.toURL,
	}); err != nil {
		return "", fmt.Errorf("failed to execute to template: %w", err)
	}

	toElement := toBuffer.String()
	// CRÍTICO: Usar Exclusive C14N con InclusiveNamespaces="soap wcf" como especifica el Transform
	toC14N, err := xmlpkg.CanonicalizeExclusive([]byte(toElement), []string{"soap", "wcf"})
	if err != nil {
		return "", fmt.Errorf("failed to canonicalize To element: %w", err)
	}

	toDigest := sha256.Sum256(toC14N)
	toDigestB64 := base64.StdEncoding.EncodeToString(toDigest[:])

	// 4. Construir SignedInfo usando template
	tmplSignedInfo, err := template.ParseFiles(signedInfoTemplatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse signedInfo template: %w", err)
	}

	var signedInfoBuffer bytes.Buffer
	if err := tmplSignedInfo.Execute(&signedInfoBuffer, map[string]string{
		"ToID":        sh.idTo,
		"DigestValue": toDigestB64,
	}); err != nil {
		return "", fmt.Errorf("failed to execute signedinfo template: %w", err)
	}

	signedInfo := signedInfoBuffer.String()

	// 5. Canonicalizar SignedInfo con Exclusive C14N
	signedInfoC14N, err := xmlpkg.CanonicalizeExclusive([]byte(signedInfo), []string{"wsa", "soap", "wcf"})
	if err != nil {
		return "", fmt.Errorf("failed to canonicalize SignedInfo: %w", err)
	}

	// 6. Firmar SignedInfo
	signedInfoHash := sha256.Sum256(signedInfoC14N)
	signature, err := rsa.SignPKCS1v15(rand.Reader, sh.privateKey, crypto.SHA256, signedInfoHash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	// 7. Construir el security header final usando template
	tmplHeader, err := template.ParseFiles(securityHeaderTemplatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse security header template: %w", err)
	}

	var headerBuffer bytes.Buffer
	if err := tmplHeader.Execute(&headerBuffer, map[string]string{
		"TimestampID":              sh.idTimestamp,
		"Created":                  created,
		"Expires":                  expires,
		"BinarySecurityTokenID":    sh.idBinarySecurityToken,
		"Certificate":              certB64,
		"SignatureID":              sh.idSignature,
		"ToID":                     sh.idTo,
		"DigestValue":              toDigestB64,
		"SignatureValue":           signatureB64,
		"KeyInfoID":                sh.idKeyInfo,
		"SecurityTokenReferenceID": sh.idSecurityTokenReference,
		"Action":                   sh.action,
		"ToURL":                    sh.toURL,
	}); err != nil {
		return "", fmt.Errorf("failed to execute header template: %w", err)
	}

	return headerBuffer.String(), nil
}

// generateUniqueID genera un ID único basado en timestamp y hash aleatorio
func generateUniqueID() string {
	timestamp := time.Now().UnixNano()
	
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return fmt.Sprintf("%d", timestamp)
	}
	
	combined := fmt.Sprintf("%d%s", timestamp, hex.EncodeToString(randomBytes))
	hash := sha256.Sum256([]byte(combined))
	
	return hex.EncodeToString(hash[:])[:16]
}
