package applicationresponse

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// ApplicationResponseXML representa un ApplicationResponse UBL 2.1 con tags XML
type ApplicationResponseXML struct {
	XMLName xml.Name `xml:"ApplicationResponse"`

	// Namespaces
	Xmlns             string `xml:"xmlns,attr"`
	XmlnsCAC          string `xml:"xmlns:cac,attr"`
	XmlnsCBC          string `xml:"xmlns:cbc,attr"`
	XmlnsCCTS         string `xml:"xmlns:ccts,attr"`
	XmlnsDS           string `xml:"xmlns:ds,attr"`
	XmlnsExt          string `xml:"xmlns:ext,attr"`
	XmlnsSts          string `xml:"xmlns:sts,attr"`
	XmlnsXades        string `xml:"xmlns:xades,attr"`
	XmlnsXades141     string `xml:"xmlns:xades141,attr"`
	XmlnsXsi          string `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation string `xml:"xsi:schemaLocation,attr"`

	// UBLExtensions (para firma XAdES de DIAN)
	UBLExtensions types.UBLExtensions `xml:"ext:UBLExtensions"`

	// Identificación básica
	UBLVersionID       types.CBCElement  `xml:"cbc:UBLVersionID"`
	CustomizationID    types.CBCElement  `xml:"cbc:CustomizationID"`
	ProfileID          types.CBCElement  `xml:"cbc:ProfileID"`
	ProfileExecutionID types.CBCElement  `xml:"cbc:ProfileExecutionID"`
	ID                 types.CBCElement  `xml:"cbc:ID"`
	UUID               types.UUIDElement `xml:"cbc:UUID"`
	IssueDate          types.CBCElement  `xml:"cbc:IssueDate"`
	IssueTime          types.CBCElement  `xml:"cbc:IssueTime"`

	// Notas
	Note []types.CBCElement `xml:"cbc:Note,omitempty"`

	// Partes (DIAN como emisor, empresa como receptor)
	SenderParty   SenderPartyXML   `xml:"cac:SenderParty"`
	ReceiverParty ReceiverPartyXML `xml:"cac:ReceiverParty"`

	// Respuesta del documento
	DocumentResponse DocumentResponseXML `xml:"cac:DocumentResponse"`
}

// SenderPartyXML emisor del ApplicationResponse (DIAN)
type SenderPartyXML struct {
	PartyTaxScheme PartyTaxSchemeXML `xml:"cac:PartyTaxScheme"`
}

// ReceiverPartyXML receptor del ApplicationResponse (empresa)
type ReceiverPartyXML struct {
	PartyTaxScheme PartyTaxSchemeXML `xml:"cac:PartyTaxScheme"`
}

// PartyTaxSchemeXML información fiscal de una parte
type PartyTaxSchemeXML struct {
	RegistrationName types.CBCElement         `xml:"cbc:RegistrationName"`
	CompanyID        types.IDElement          `xml:"cbc:CompanyID"`
	TaxLevelCode     types.TaxLevelCodeElement `xml:"cbc:TaxLevelCode,omitempty"`
	TaxScheme        TaxSchemeXML             `xml:"cac:TaxScheme"`
}

// TaxSchemeXML esquema de impuestos
type TaxSchemeXML struct {
	ID   types.CBCElement `xml:"cbc:ID"`
	Name types.CBCElement `xml:"cbc:Name"`
}

// DocumentResponseXML respuesta del documento
type DocumentResponseXML struct {
	Response          ResponseXML          `xml:"cac:Response"`
	DocumentReference DocumentReferenceXML `xml:"cac:DocumentReference"`
}

// ResponseXML respuesta de validación
type ResponseXML struct {
	ResponseCode types.CBCElement   `xml:"cbc:ResponseCode"`
	Description  []types.CBCElement `xml:"cbc:Description,omitempty"`
}

// DocumentReferenceXML referencia al documento validado
type DocumentReferenceXML struct {
	ID                    types.CBCElement        `xml:"cbc:ID"`
	UUID                  types.UUIDElement       `xml:"cbc:UUID"`
	IssueDate             types.CBCElement        `xml:"cbc:IssueDate"`
	DocumentTypeCode      *types.CBCElement       `xml:"cbc:DocumentTypeCode,omitempty"`
	DocumentType          *types.CBCElement       `xml:"cbc:DocumentType,omitempty"`
	ResultOfVerification  *ResultOfVerificationXML `xml:"cac:ResultOfVerification,omitempty"`
}

// ResultOfVerificationXML resultado de la verificación
type ResultOfVerificationXML struct {
	ValidatorID            types.CBCElement `xml:"cbc:ValidatorID"`
	ValidationResultCode   types.CBCElement `xml:"cbc:ValidationResultCode"`
	ValidationDate         types.CBCElement `xml:"cbc:ValidationDate"`
	ValidationTime         types.CBCElement `xml:"cbc:ValidationTime"`
	ValidateProcess        *types.CBCElement `xml:"cbc:ValidateProcess,omitempty"`
	ValidateTool           *types.CBCElement `xml:"cbc:ValidateTool,omitempty"`
	ValidateToolVersion    *types.CBCElement `xml:"cbc:ValidateToolVersion,omitempty"`
}
