package signature

import "encoding/xml"

// Signature estructura de firma XMLDSig
type Signature struct {
	XMLName        xml.Name       `xml:"ds:Signature"`
	Xmlns          string         `xml:"xmlns:ds,attr"`
	ID             string         `xml:"Id,attr"`
	SignedInfo     SignedInfo     `xml:"ds:SignedInfo"`
	SignatureValue SignatureValue `xml:"ds:SignatureValue"`
	KeyInfo        KeyInfo        `xml:"ds:KeyInfo"`
	Object         *Object        `xml:"ds:Object,omitempty"`
}

// SignedInfo información firmada
type SignedInfo struct {
	XMLName                xml.Name               `xml:"ds:SignedInfo"`
	CanonicalizationMethod CanonicalizationMethod `xml:"ds:CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"ds:SignatureMethod"`
	Reference              []Reference            `xml:"ds:Reference"`
}

// CanonicalizationMethod método de canonicalización
type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// SignatureMethod método de firma
type SignatureMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// Reference referencia a un elemento
type Reference struct {
	ID           string       `xml:"Id,attr,omitempty"`
	URI          string       `xml:"URI,attr"`
	Type         string       `xml:"Type,attr,omitempty"`
	Transforms   *Transforms  `xml:"ds:Transforms,omitempty"`
	DigestMethod DigestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

// Transforms transformaciones
type Transforms struct {
	Transform []Transform `xml:"ds:Transform"`
}

// Transform transformación
type Transform struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// DigestMethod método de digest
type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// SignatureValue valor de la firma
type SignatureValue struct {
	ID    string `xml:"Id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// KeyInfo información de la clave
type KeyInfo struct {
	ID       string    `xml:"Id,attr,omitempty"`
	X509Data X509Data  `xml:"ds:X509Data"`
	KeyValue *KeyValue `xml:"ds:KeyValue,omitempty"`
}

// X509Data datos del certificado X509
type X509Data struct {
	X509Certificate string `xml:"ds:X509Certificate"`
}

// KeyValue valor de la clave pública
type KeyValue struct {
	RSAKeyValue *RSAKeyValue `xml:"ds:RSAKeyValue,omitempty"`
}

// RSAKeyValue clave RSA
type RSAKeyValue struct {
	Modulus  string `xml:"ds:Modulus"`
	Exponent string `xml:"ds:Exponent"`
}

// Object objeto de la firma
type Object struct {
	ID                   string               `xml:"Id,attr,omitempty"`
	QualifyingProperties QualifyingProperties `xml:"xades:QualifyingProperties"`
}

// QualifyingProperties propiedades calificadas XAdES
type QualifyingProperties struct {
	Xmlns            string           `xml:"xmlns:xades,attr"`
	ID               string           `xml:"Id,attr,omitempty"`
	Target           string           `xml:"Target,attr"`
	SignedProperties SignedProperties `xml:"xades:SignedProperties"`
}

// SignedPropertiesContainer contenedor de propiedades firmadas
type SignedPropertiesContainer struct {
	SignedProperties SignedProperties
}

// SignedProperties propiedades firmadas
type SignedProperties struct {
	ID                          string                       `xml:"Id,attr"`
	SignedSignatureProperties   SignedSignatureProperties    `xml:"xades:SignedSignatureProperties"`
	SignedDataObjectProperties  *SignedDataObjectProperties  `xml:"xades:SignedDataObjectProperties,omitempty"`
}

// SignedSignatureProperties propiedades de firma firmadas
type SignedSignatureProperties struct {
	SigningTime               string                    `xml:"xades:SigningTime"`
	SigningCertificate        SigningCertificate        `xml:"xades:SigningCertificate"`
	SignaturePolicyIdentifier SignaturePolicyIdentifier `xml:"xades:SignaturePolicyIdentifier"`
	SignerRole                SignerRole                `xml:"xades:SignerRole"`
}

// SigningCertificate certificado de firma
type SigningCertificate struct {
	Cert Cert `xml:"xades:Cert"`
}

// Cert certificado
type Cert struct {
	CertDigest   CertDigest   `xml:"xades:CertDigest"`
	IssuerSerial IssuerSerial `xml:"xades:IssuerSerial"`
}

// CertDigest digest del certificado
type CertDigest struct {
	DigestMethod DigestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

// IssuerSerial emisor y serial del certificado
type IssuerSerial struct {
	X509IssuerName   string `xml:"ds:X509IssuerName"`
	X509SerialNumber string `xml:"ds:X509SerialNumber"`
}

// SignaturePolicyIdentifier identificador de política de firma
type SignaturePolicyIdentifier struct {
	SignaturePolicyId SignaturePolicyId `xml:"xades:SignaturePolicyId"`
}

// SignaturePolicyId ID de política de firma
type SignaturePolicyId struct {
	SigPolicyId   SigPolicyId   `xml:"xades:SigPolicyId"`
	SigPolicyHash SigPolicyHash `xml:"xades:SigPolicyHash"`
}

// SigPolicyId identificador de política
type SigPolicyId struct {
	Identifier string `xml:"xades:Identifier"`
}

// SigPolicyHash hash de política
type SigPolicyHash struct {
	DigestMethod DigestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

// SignerRole rol del firmante
type SignerRole struct {
	ClaimedRoles ClaimedRoles `xml:"xades:ClaimedRoles"`
}

// ClaimedRoles roles reclamados
type ClaimedRoles struct {
	ClaimedRole string `xml:"xades:ClaimedRole"`
}

// SignedDataObjectProperties propiedades de objetos de datos firmados
type SignedDataObjectProperties struct {
	DataObjectFormat DataObjectFormat `xml:"xades:DataObjectFormat"`
}

// DataObjectFormat formato de objeto de datos
type DataObjectFormat struct {
	ObjectReference string `xml:"ObjectReference,attr"`
	MimeType        string `xml:"xades:MimeType"`
	Encoding        string `xml:"xades:Encoding"`
}
