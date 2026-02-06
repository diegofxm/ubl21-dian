package attached

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// AttachedDocumentXML representa un AttachedDocument UBL 2.1 con tags XML
type AttachedDocumentXML struct {
	XMLName xml.Name `xml:"AttachedDocument"`

	// Namespaces
	Xmlns         string `xml:"xmlns,attr"`
	XmlnsCAC      string `xml:"xmlns:cac,attr"`
	XmlnsCBC      string `xml:"xmlns:cbc,attr"`
	XmlnsCCTS     string `xml:"xmlns:ccts,attr"`
	XmlnsExt      string `xml:"xmlns:ext,attr"`
	XmlnsXades    string `xml:"xmlns:xades,attr"`
	XmlnsXades141 string `xml:"xmlns:xades141,attr"`
	XmlnsDS       string `xml:"xmlns:ds,attr"`

	// UBLExtensions (para la firma)
	UBLExtensions types.UBLExtensions `xml:"ext:UBLExtensions"`

	// Identificación básica
	UBLVersionID       types.CBCElement `xml:"cbc:UBLVersionID"`
	CustomizationID    types.IDElement  `xml:"cbc:CustomizationID"`
	ProfileID          types.CBCElement `xml:"cbc:ProfileID"`
	ProfileExecutionID types.CBCElement `xml:"cbc:ProfileExecutionID"`
	ID                 types.CBCElement `xml:"cbc:ID"`
	IssueDate          types.CBCElement `xml:"cbc:IssueDate"`
	IssueTime          types.CBCElement `xml:"cbc:IssueTime"`
	DocumentType       types.CBCElement `xml:"cbc:DocumentType"`
	ParentDocumentID   types.CBCElement `xml:"cbc:ParentDocumentID"`

	// Partes
	SenderParty   SenderPartyXML    `xml:"cac:SenderParty"`
	ReceiverParty *ReceiverPartyXML `xml:"cac:ReceiverParty,omitempty"`

	// Attachment: Factura firmada completa en CDATA
	Attachment AttachmentXML `xml:"cac:Attachment"`

	// ParentDocumentLineReference: ApplicationResponse en CDATA
	ParentDocumentLineReference ParentDocumentLineReferenceXML `xml:"cac:ParentDocumentLineReference"`
}

// SenderPartyXML emisor del AttachedDocument
type SenderPartyXML struct {
	PartyTaxScheme types.PartyTaxSchemeXML `xml:"cac:PartyTaxScheme"`
}

// ReceiverPartyXML receptor del AttachedDocument
type ReceiverPartyXML struct {
	PartyTaxScheme types.PartyTaxSchemeXML `xml:"cac:PartyTaxScheme"`
}

// AttachmentXML contiene la factura firmada en CDATA
type AttachmentXML struct {
	ExternalReference ExternalReferenceXML `xml:"cac:ExternalReference"`
}

// ExternalReferenceXML referencia externa con CDATA
type ExternalReferenceXML struct {
	MimeCode     types.CBCElement   `xml:"cbc:MimeCode"`
	EncodingCode types.CBCElement   `xml:"cbc:EncodingCode"`
	Description  types.CDATAElement `xml:"cbc:Description"`
}

// ParentDocumentLineReferenceXML referencia al documento padre con ApplicationResponse
type ParentDocumentLineReferenceXML struct {
	LineID            types.CBCElement     `xml:"cbc:LineID"`
	DocumentReference DocumentReferenceXML `xml:"cac:DocumentReference"`
}

// DocumentReferenceXML referencia al documento
type DocumentReferenceXML struct {
	ID                   types.CBCElement        `xml:"cbc:ID"`
	UUID                 types.UUIDElement       `xml:"cbc:UUID"`
	IssueDate            types.CBCElement        `xml:"cbc:IssueDate"`
	DocumentType         types.CBCElement        `xml:"cbc:DocumentType"`
	Attachment           AttachmentXML           `xml:"cac:Attachment"`
	ResultOfVerification ResultOfVerificationXML `xml:"cac:ResultOfVerification"`
}

// ResultOfVerificationXML resultado de verificación de DIAN
type ResultOfVerificationXML struct {
	ValidatorID          types.CBCElement `xml:"cbc:ValidatorID"`
	ValidationResultCode types.CBCElement `xml:"cbc:ValidationResultCode"`
	ValidationDate       types.CBCElement `xml:"cbc:ValidationDate"`
	ValidationTime       types.CBCElement `xml:"cbc:ValidationTime"`
}
