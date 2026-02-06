package model

import "github.com/diegofxm/ubl21-dian/documents/common/types"

// DianExtensions contiene las extensiones específicas de DIAN (común para Invoice, CreditNote, DebitNote)
type DianExtensions struct {
	InvoiceControl        InvoiceControl               `xml:"sts:InvoiceControl"`
	InvoiceSource         InvoiceSource                `xml:"sts:InvoiceSource"`
	SoftwareProvider      SoftwareProvider             `xml:"sts:SoftwareProvider"`
	SoftwareSecurityCode  SoftwareSecurityCodeElement  `xml:"sts:SoftwareSecurityCode"`
	AuthorizationProvider AuthorizationProvider        `xml:"sts:AuthorizationProvider"`
	QRCode                types.CBCElement             `xml:"sts:QRCode"`
}

// InvoiceControl información de control de la factura/nota
type InvoiceControl struct {
	InvoiceAuthorization types.CBCElement   `xml:"sts:InvoiceAuthorization"`
	AuthorizationPeriod  AuthorizationPeriod `xml:"sts:AuthorizationPeriod"`
	AuthorizedInvoices   AuthorizedInvoices  `xml:"sts:AuthorizedInvoices"`
}

// AuthorizationPeriod período de autorización
type AuthorizationPeriod struct {
	StartDate types.CBCElement `xml:"cbc:StartDate"`
	EndDate   types.CBCElement `xml:"cbc:EndDate"`
}

// AuthorizedInvoices documentos autorizados
type AuthorizedInvoices struct {
	Prefix types.CBCElement `xml:"sts:Prefix"`
	From   types.CBCElement `xml:"sts:From"`
	To     types.CBCElement `xml:"sts:To"`
}

// InvoiceSource fuente del documento
type InvoiceSource struct {
	IdentificationCode types.IdentificationCodeElement `xml:"cbc:IdentificationCode"`
}

// SoftwareProvider proveedor de software
type SoftwareProvider struct {
	ProviderID types.IDElement   `xml:"sts:ProviderID"`
	SoftwareID SoftwareIDElement `xml:"sts:SoftwareID"`
}

// SoftwareIDElement elemento con atributos para software ID
type SoftwareIDElement struct {
	SchemeAgencyID   string `xml:"schemeAgencyID,attr,omitempty"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr,omitempty"`
	Value            string `xml:",chardata"`
}

// SoftwareSecurityCodeElement elemento con atributos para código de seguridad
type SoftwareSecurityCodeElement struct {
	SchemeAgencyID   string `xml:"schemeAgencyID,attr,omitempty"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr,omitempty"`
	Value            string `xml:",chardata"`
}

// AuthorizationProvider proveedor de autorización (siempre DIAN)
type AuthorizationProvider struct {
	AuthorizationProviderID types.IDElement `xml:"sts:AuthorizationProviderID"`
}

// BillingReferenceXML referencia a factura (común para CreditNote y DebitNote)
type BillingReferenceXML struct {
	InvoiceDocumentReference InvoiceDocumentReferenceXML `xml:"cac:InvoiceDocumentReference"`
}

// InvoiceDocumentReferenceXML referencia al documento de factura
type InvoiceDocumentReferenceXML struct {
	ID        types.CBCElement  `xml:"cbc:ID"`
	UUID      types.UUIDElement `xml:"cbc:UUID"`
	IssueDate types.CBCElement  `xml:"cbc:IssueDate"`
}
