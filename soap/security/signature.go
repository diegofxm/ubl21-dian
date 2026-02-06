package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	xmlpkg "github.com/diegofxm/ubl21-dian/xml"
)

// SignatureData contiene los datos necesarios para crear una firma XML
type SignatureData struct {
	ToID         string
	ToURL        string
	DigestValue  string
	SignatureValue string
}

// ComputeDigest calcula el digest SHA256 de un elemento XML canonicalizado
func ComputeDigest(xmlElement []byte, inclusiveNamespaces []string) (string, error) {
	// Canonicalizar con Exclusive C14N
	canonical, err := xmlpkg.CanonicalizeExclusive(xmlElement, inclusiveNamespaces)
	if err != nil {
		return "", fmt.Errorf("failed to canonicalize: %w", err)
	}

	// Calcular SHA256
	digest := sha256.Sum256(canonical)
	return base64.StdEncoding.EncodeToString(digest[:]), nil
}

// SignData firma datos usando RSA-SHA256
func SignData(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	// Calcular hash SHA256
	hash := sha256.Sum256(data)
	
	// Firmar con RSA
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// CanonicalizeSigned canonicaliza el SignedInfo para firma
func CanonicalizeSigned(signedInfoXML []byte) ([]byte, error) {
	return xmlpkg.CanonicalizeExclusive(signedInfoXML, []string{"wsa", "soap", "wcf"})
}
