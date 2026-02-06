package supportdocument

import "time"

// SupportDocumentData datos principales del documento soporte
type SupportDocumentData struct {
	// Identificación
	Number             string
	CUDS               string // Código Único de Documento Soporte
	IssueDate          time.Time
	IssueTime          string
	OperationType      string // Tipo de operación (compra nacional, importación, etc)
	
	// Referencias a documentos externos (facturas de proveedor)
	BillingReferences []BillingReferenceData
	
	// Partes (Comprador = Supplier, Proveedor = Customer en UBL)
	Buyer    PartyData // La empresa que compra (emisor del documento soporte)
	Supplier PartyData // El proveedor (receptor)
	
	// Líneas
	Lines []SupportDocumentLineData
	
	// Totales
	Totals TotalsData
	
	// Extensiones DIAN
	DianExtensions DianExtensionsData
	
	// Notas adicionales
	Notes []string
}

// BillingReferenceData referencia a factura de proveedor
type BillingReferenceData struct {
	InvoiceID string    // Número de factura del proveedor
	UUID      string    // UUID si existe
	IssueDate time.Time // Fecha de la factura del proveedor
}

// PartyData datos de una parte (comprador o proveedor)
type PartyData struct {
	PersonType               string // 1=Persona jurídica, 2=Persona natural
	ID                       string // NIT o identificación
	DV                       string // Dígito de verificación
	DocumentType             string // 31=NIT, 13=CC, etc
	Name                     string // Razón social o nombre
	TaxLevelCode             string // Código de responsabilidad fiscal
	TaxSchemeID              string // 01=IVA
	TaxSchemeName            string // IVA
	IndustryClassificationCode string // Código CIIU
	Address                  AddressData
	Contact                  ContactData
}

// AddressData datos de dirección
type AddressData struct {
	ID                   string // Código de ciudad/municipio
	CityName             string
	PostalZone           string
	CountrySubentity     string // Departamento
	CountrySubentityCode string // Código departamento
	AddressLine          string
	Country              CountryData
}

// CountryData datos de país
type CountryData struct {
	Code string // CO
	Name string // Colombia
}

// ContactData datos de contacto
type ContactData struct {
	Telephone string
	Email     string
}

// SupportDocumentLineData línea del documento soporte
type SupportDocumentLineData struct {
	ID                    string
	Quantity              float64
	UnitCode              string // EA, KG, etc
	LineExtensionAmount   float64
	FreeOfChargeIndicator bool
	Item                  ItemData
	Price                 PriceData
	TaxTotals             []TaxTotalData
	WithholdingTaxTotals  []TaxTotalData // Retenciones
}

// ItemData datos del item/producto
type ItemData struct {
	Description      string
	StandardItemID   ItemIDData
	AdditionalItemID ItemIDData
}

// ItemIDData identificación de item
type ItemIDData struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// PriceData datos de precio
type PriceData struct {
	Amount       float64
	BaseQuantity float64
}

// TaxTotalData total de impuestos o retenciones
type TaxTotalData struct {
	TaxAmount    float64
	TaxSubtotals []TaxSubtotalData
}

// TaxSubtotalData subtotal de impuesto
type TaxSubtotalData struct {
	TaxableAmount float64
	TaxAmount     float64
	Percent       float64
	TaxCategory   TaxCategoryData
}

// TaxCategoryData categoría de impuesto
type TaxCategoryData struct {
	ID   string // 01=IVA, 04=INC, etc
	Name string
}

// TotalsData totales del documento
type TotalsData struct {
	LineExtensionAmount  float64
	TaxExclusiveAmount   float64
	TaxInclusiveAmount   float64
	AllowanceTotalAmount float64
	ChargeTotalAmount    float64
	PayableAmount        float64
}

// DianExtensionsData extensiones DIAN
type DianExtensionsData struct {
	InvoiceAuthorization string
	AuthorizationPeriod  AuthorizationPeriodData
	InvoiceControl       InvoiceControlData
	Provider             ProviderData
	Software             SoftwareData
	QRCode               string
}

// AuthorizationPeriodData período de autorización
type AuthorizationPeriodData struct {
	StartDate time.Time
	EndDate   time.Time
}

// InvoiceControlData control de numeración
type InvoiceControlData struct {
	Prefix string
	From   string
	To     string
}

// ProviderData proveedor tecnológico
type ProviderData struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// SoftwareData software de facturación
type SoftwareData struct {
	ID           string
	SecurityCode string
}

// ============================================================================
// Template Types
// ============================================================================

// SupportDocumentTemplateData datos para template
type SupportDocumentTemplateData struct {
	// DIAN Extensions
	InvoiceAuthorization string
	AuthPeriodStartDate  string
	AuthPeriodEndDate    string
	Prefix               string
	From                 string
	To                   string
	ProviderID           string
	ProviderSchemeID     string
	ProviderSchemeName   string
	SoftwareID           string
	SecurityCode         string
	QRCode               string

	// Header
	ProfileExecutionID string
	SupportDocNumber   string
	CUDS               string
	Environment        string
	IssueDate          string
	IssueTime          string
	DocumentTypeCode   string // "05"
	CurrencyCode       string
	LineCount          int

	// Notes
	Notes []string

	// Billing References
	BillingReferences []BillingReferenceTemplateData

	// Parties
	Buyer    PartyTemplateData
	Supplier PartyTemplateData

	// Totals
	LineExtensionAmount  string
	TaxExclusiveAmount   string
	TaxInclusiveAmount   string
	AllowanceTotalAmount string
	ChargeTotalAmount    string
	PayableAmount        string

	// Lines
	SupportDocumentLines []SupportDocumentLineTemplateData
	TaxTotals            []TaxTotalTemplateData
	WithholdingTaxTotals []TaxTotalTemplateData
}

// BillingReferenceTemplateData referencia para template
type BillingReferenceTemplateData struct {
	InvoiceID string
	UUID      string
	IssueDate string
}

// PartyTemplateData datos de parte para template
type PartyTemplateData struct {
	PersonType               string
	ID                       string
	DV                       string
	DocumentType             string
	Name                     string
	TaxLevelCode             string
	TaxSchemeID              string
	TaxSchemeName            string
	IndustryClassificationCode string
	Address                  AddressTemplateData
	Contact                  ContactTemplateData
}

// AddressTemplateData dirección para template
type AddressTemplateData struct {
	ID                   string
	CityName             string
	PostalZone           string
	CountrySubentity     string
	CountrySubentityCode string
	AddressLine          string
	CountryCode          string
	CountryName          string
}

// ContactTemplateData contacto para template
type ContactTemplateData struct {
	Telephone string
	Email     string
}

// SupportDocumentLineTemplateData línea para template
type SupportDocumentLineTemplateData struct {
	ID                    string
	UnitCode              string
	Quantity              string
	LineExtensionAmount   string
	FreeOfChargeIndicator string
	CurrencyID            string
	Item                  ItemTemplateData
	Price                 PriceTemplateData
}

// ItemTemplateData item para template
type ItemTemplateData struct {
	Description      string
	StandardItemID   ItemIDTemplateData
	AdditionalItemID ItemIDTemplateData
}

// ItemIDTemplateData ID de item para template
type ItemIDTemplateData struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// PriceTemplateData precio para template
type PriceTemplateData struct {
	Amount       string
	BaseQuantity string
}

// TaxTotalTemplateData total de impuestos para template
type TaxTotalTemplateData struct {
	TaxAmount    string
	CurrencyID   string
	TaxSubtotals []TaxSubtotalTemplateData
}

// TaxSubtotalTemplateData subtotal para template
type TaxSubtotalTemplateData struct {
	TaxableAmount string
	TaxAmount     string
	CurrencyID    string
	Percent       string
	TaxCategory   TaxCategoryTemplateData
}

// TaxCategoryTemplateData categoría para template
type TaxCategoryTemplateData struct {
	ID   string
	Name string
}
