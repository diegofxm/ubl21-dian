package signature

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	xmlpkg "github.com/diegofxm/ubl21-dian/xml"
	"software.sslmate.com/src/go-pkcs12"
)

var (
	ErrInvalidCertificate = errors.New("invalid certificate")
	ErrSignatureFailed    = errors.New("signature failed")
)

// Signer firma documentos XML con XAdES-BES
type Signer struct {
	privateKey  *rsa.PrivateKey
	certificate *x509.Certificate
	certChain   []*x509.Certificate
}

// NewSignerFromP12 crea un signer desde un archivo .p12
func NewSignerFromP12(p12Path, password string) (*Signer, error) {
	p12Data, err := os.ReadFile(p12Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read p12 file: %w", err)
	}

	privateKey, certificate, caCerts, err := pkcs12.DecodeChain(p12Data, password)
	if err != nil {
		var cert *x509.Certificate
		privateKey, cert, err = pkcs12.Decode(p12Data, password)
		if err != nil {
			return nil, fmt.Errorf("failed to decode p12: %w", err)
		}
		certificate = cert
		caCerts = nil
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not RSA")
	}

	certChain := []*x509.Certificate{certificate}
	if caCerts != nil {
		certChain = append(certChain, caCerts...)
	}

	return &Signer{
		privateKey:  rsaKey,
		certificate: certificate,
		certChain:   certChain,
	}, nil
}

// NewSignerFromSinglePEM crea un signer desde un solo archivo PEM
func NewSignerFromSinglePEM(pemPath string) (*Signer, error) {
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read PEM file: %w", err)
	}

	var certificate *x509.Certificate
	var privateKey *rsa.PrivateKey
	var certChain []*x509.Certificate

	rest := pemData
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
			certChain = append(certChain, cert)

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

	return &Signer{
		privateKey:  privateKey,
		certificate: certificate,
		certChain:   certChain,
	}, nil
}

// NewSignerFromPEM crea un signer desde archivos PEM separados
func NewSignerFromPEM(certPath, keyPath string) (*Signer, error) {
	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, ErrInvalidCertificate
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		return nil, errors.New("invalid private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		key, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
	}

	return &Signer{
		privateKey:  privateKey,
		certificate: cert,
		certChain:   []*x509.Certificate{cert},
	}, nil
}

// SignXML firma un documento XML con XAdES-BES
func (s *Signer) SignXML(xmlData []byte) ([]byte, error) {
	const (
		signatureID      = "xmldsig-88fbfc55-6dac-4b36-916d-4cfae7c0bb71"
		keyInfoID        = "xmldsig-88fbfc55-6dac-4b36-916d-4cfae7c0bb71-keyinfo"
		signedPropsID    = "xmldsig-88fbfc55-6dac-4b36-916d-4cfae7c0bb71-signedprops"
		signatureValueID = "xmldsig-88fbfc55-6dac-4b36-916d-4cfae7c0bb71-sigvalue"
		referenceID      = "xmldsig-88fbfc55-6dac-4b36-916d-4cfae7c0bb71-ref0"
	)

	// 1. Ordenar atributos del documento
	xmlDataSorted, err := sortAllAttributes(xmlData)
	if err != nil {
		return nil, fmt.Errorf("failed to sort document attributes: %w", err)
	}

	// 2. Canonicalizar documento y calcular digest
	documentC14N, err := xmlpkg.Canonicalize(xmlDataSorted)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize document: %w", err)
	}

	documentDigest := sha256.Sum256(documentC14N)
	documentDigestB64 := base64.StdEncoding.EncodeToString(documentDigest[:])

	// 3. Construir KeyInfo y SignedProperties usando TEMPLATES (con xmlns)
	keyInfoXML := buildKeyInfoTemplate(keyInfoID, s.certificate)
	signedPropsXML := buildSignedPropertiesTemplate(signedPropsID, s.certificate)

	// 4. Canonicalizar KeyInfo y SignedProperties (ya tienen xmlns correctos)
	keyInfoC14N, err := xmlpkg.Canonicalize(keyInfoXML)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize keyinfo: %w", err)
	}

	signedPropsC14N, err := xmlpkg.Canonicalize(signedPropsXML)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize signedprops: %w", err)
	}

	// 5. Calcular digests
	keyInfoDigest := sha256.Sum256(keyInfoC14N)
	keyInfoDigestB64 := base64.StdEncoding.EncodeToString(keyInfoDigest[:])

	signedPropsDigest := sha256.Sum256(signedPropsC14N)
	signedPropsDigestB64 := base64.StdEncoding.EncodeToString(signedPropsDigest[:])

	// 6. Crear SignedInfo con las 3 referencias
	signedInfo := SignedInfo{
		CanonicalizationMethod: CanonicalizationMethod{
			Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315",
		},
		SignatureMethod: SignatureMethod{
			Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
		},
		Reference: []Reference{
			{
				ID:  referenceID,
				URI: "",
				Transforms: &Transforms{
					Transform: []Transform{
						{Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature"},
					},
				},
				DigestMethod: DigestMethod{
					Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
				},
				DigestValue: documentDigestB64,
			},
			{
				URI: "#" + keyInfoID,
				DigestMethod: DigestMethod{
					Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
				},
				DigestValue: keyInfoDigestB64,
			},
			{
				Type: "http://uri.etsi.org/01903#SignedProperties",
				URI:  "#" + signedPropsID,
				DigestMethod: DigestMethod{
					Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
				},
				DigestValue: signedPropsDigestB64,
			},
		},
	}

	// 7. Serializar SignedInfo y ordenar atributos
	signedInfoXML, err := xml.Marshal(signedInfo)
	if err != nil {
		return nil, err
	}

	signedInfoXML, err = sortAllAttributes(signedInfoXML)
	if err != nil {
		return nil, err
	}

	// 8. Agregar prefijos ds: a SignedInfo
	signedInfoXML = []byte(regexp.MustCompile(`<(SignedInfo|CanonicalizationMethod|SignatureMethod|Reference|Transforms|Transform|DigestMethod|DigestValue)(\s|>)`).ReplaceAllString(
		string(signedInfoXML),
		"<ds:$1$2",
	))
	signedInfoXML = []byte(regexp.MustCompile(`</(SignedInfo|CanonicalizationMethod|SignatureMethod|Reference|Transforms|Transform|DigestMethod|DigestValue)>`).ReplaceAllString(
		string(signedInfoXML),
		"</ds:$1>",
	))

	// 9. Envolver SignedInfo solo con xmlns:ds para canonicalizar
	// CRÍTICO: SignedInfo NO debe heredar namespaces del documento raíz
	signedInfoWrapped := []byte(`<ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#">` + string(signedInfoXML) + `</ds:Signature>`)
	signedInfoC14NFull, err := xmlpkg.Canonicalize(signedInfoWrapped)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize signedinfo: %w", err)
	}

	// 10. Extraer SignedInfo canonicalizado
	startTag := []byte("<ds:SignedInfo")
	endTag := []byte("</ds:SignedInfo>")
	startIdx := bytes.Index(signedInfoC14NFull, startTag)
	endIdx := bytes.Index(signedInfoC14NFull, endTag)
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("failed to extract SignedInfo from canonicalized wrapper")
	}
	signedInfoC14N := signedInfoC14NFull[startIdx : endIdx+len(endTag)]

	// 11. Firmar SignedInfo
	signedInfoHash := sha256.Sum256(signedInfoC14N)
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, signedInfoHash[:])
	if err != nil {
		return nil, ErrSignatureFailed
	}

	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	// 12. Construir firma final usando el SignedInfo CANONICALIZADO (el que se firmó)
	// CRÍTICO: El SignedInfo en el XML final debe ser idéntico al que se firmó
	sigXML := []byte(
		`<ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#" Id="` + signatureID + `">` +
			string(signedInfoC14N) +
			`<ds:SignatureValue Id="` + signatureValueID + `">` + signatureB64 + `</ds:SignatureValue>` +
			string(keyInfoXML) +
			`<ds:Object><xades:QualifyingProperties xmlns:xades="http://uri.etsi.org/01903/v1.3.2#" Target="#` + signatureID + `">` +
			string(signedPropsXML) +
			`</xades:QualifyingProperties></ds:Object>` +
			`</ds:Signature>`,
	)

	// 13. Insertar firma en UBLExtensions
	signedXML, err := insertSignatureIntoUBLExtension(xmlData, sigXML)
	if err != nil {
		return nil, fmt.Errorf("failed to insert signature: %w", err)
	}

	return signedXML, nil
}

// insertSignatureIntoUBLExtension inserta la firma en UBLExtensions
func insertSignatureIntoUBLExtension(xmlData, signatureXML []byte) ([]byte, error) {
	xmlStr := string(xmlData)
	sigStr := string(signatureXML)

	closingTag := "</ext:UBLExtensions>"
	idx := strings.Index(xmlStr, closingTag)
	if idx == -1 {
		return nil, errors.New("UBLExtensions closing tag not found")
	}

	secondExtension := `<ext:UBLExtension><ext:ExtensionContent>` + sigStr + `</ext:ExtensionContent></ext:UBLExtension>`
	result := xmlStr[:idx] + secondExtension + xmlStr[idx:]

	return []byte(result), nil
}

// sortAllAttributes ordena todos los atributos XML alfabéticamente
func sortAllAttributes(xmlData []byte) ([]byte, error) {
	xmlStr := string(xmlData)

	re := regexp.MustCompile(`<([a-zA-Z0-9_:-]+)([^>]*?)(/?)>`)

	result := re.ReplaceAllStringFunc(xmlStr, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 4 {
			return match
		}

		tagName := submatches[1]
		attrsStr := submatches[2]
		selfClosing := submatches[3]

		attrs := parseAttributes(attrsStr)
		sort.Slice(attrs, func(i, j int) bool {
			return attrs[i].name < attrs[j].name
		})

		sortedAttrs := ""
		for _, attr := range attrs {
			sortedAttrs += fmt.Sprintf(` %s="%s"`, attr.name, attr.value)
		}

		return fmt.Sprintf("<%s%s%s>", tagName, sortedAttrs, selfClosing)
	})

	return []byte(result), nil
}

type attribute struct {
	name  string
	value string
}

func parseAttributes(attrsStr string) []attribute {
	var attrs []attribute

	attrRe := regexp.MustCompile(`([a-zA-Z0-9_:-]+)="([^"]*)"`)
	matches := attrRe.FindAllStringSubmatch(attrsStr, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			attrs = append(attrs, attribute{
				name:  match[1],
				value: match[2],
			})
		}
	}

	return attrs
}
