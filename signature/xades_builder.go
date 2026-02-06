package signature

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"time"
)

// buildKeyInfoTemplate construye el elemento KeyInfo con namespaces correctos usando template
func buildKeyInfoTemplate(keyInfoID string, cert *x509.Certificate) []byte {
	certB64 := base64.StdEncoding.EncodeToString(cert.Raw)
	
	xml := fmt.Sprintf(`<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#" Id="%s"><ds:X509Data><ds:X509Certificate>%s</ds:X509Certificate></ds:X509Data></ds:KeyInfo>`,
		keyInfoID, certB64)
	
	return []byte(xml)
}

// buildSignedPropertiesTemplate construye el elemento SignedProperties con namespaces correctos usando template
func buildSignedPropertiesTemplate(signedPropsID string, cert *x509.Certificate) []byte {
	// Calcular digest del certificado
	certDigest := sha256.Sum256(cert.Raw)
	certDigestB64 := base64.StdEncoding.EncodeToString(certDigest[:])
	
	// Tiempo de firma
	signingTime := time.Now().Format("2006-01-02T15:04:05-07:00")
	
	// Construir XML con namespaces declarados
	xml := fmt.Sprintf(`<xades:SignedProperties xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:xades="http://uri.etsi.org/01903/v1.3.2#" Id="%s"><xades:SignedSignatureProperties><xades:SigningTime>%s</xades:SigningTime><xades:SigningCertificate><xades:Cert><xades:CertDigest><ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"></ds:DigestMethod><ds:DigestValue>%s</ds:DigestValue></xades:CertDigest><xades:IssuerSerial><ds:X509IssuerName>%s</ds:X509IssuerName><ds:X509SerialNumber>%d</ds:X509SerialNumber></xades:IssuerSerial></xades:Cert></xades:SigningCertificate><xades:SignaturePolicyIdentifier><xades:SignaturePolicyId><xades:SigPolicyId><xades:Identifier>https://facturaelectronica.dian.gov.co/politicadefirma/v2/politicadefirmav2.pdf</xades:Identifier></xades:SigPolicyId><xades:SigPolicyHash><ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"></ds:DigestMethod><ds:DigestValue>dMoMvtcG5aIzgYo0tIsSQeVJBDnUnfSOfBpxXrmor0Y=</ds:DigestValue></xades:SigPolicyHash></xades:SignaturePolicyId></xades:SignaturePolicyIdentifier><xades:SignerRole><xades:ClaimedRoles><xades:ClaimedRole>supplier</xades:ClaimedRole></xades:ClaimedRoles></xades:SignerRole></xades:SignedSignatureProperties></xades:SignedProperties>`,
		signedPropsID,
		signingTime,
		certDigestB64,
		cert.Issuer.String(),
		cert.SerialNumber,
	)
	
	return []byte(xml)
}
