package debitnote

import "time"

// DebitNoteData datos principales de la nota débito
type DebitNoteData struct {
	Number            string
	CUDE              string
	IssueDate         time.Time
	IssueTime         string
	DebitNoteTypeCode string // "92" = Nota débito que referencia factura electrónica
	Note              string
	BillingReference  *BillingReferenceData
	Supplier          PartyData
	Customer          PartyData
	Lines             []DebitNoteLineData
	Totals            TotalsData
	DianExtensions    DianExtensionsData
}

// BillingReferenceData referencia a la factura que se ajusta
type BillingReferenceData struct {
	InvoiceID string
	UUID      string
	IssueDate time.Time
}

// PartyData datos de una parte
type PartyData struct {
	PersonType                 string
	ID                         string
	DV                         string
	DocumentType               string
	Name                       string
	TaxLevelCode               string
	TaxSchemeID                string
	TaxSchemeName              string
	IndustryClassificationCode string
	Address                    AddressData
	Contact                    ContactData
}

// AddressData datos de dirección
type AddressData struct {
	ID                   string
	CityName             string
	PostalZone           string
	CountrySubentity     string
	CountrySubentityCode string
	AddressLine          string
	CountryCode          string
	CountryName          string
}

// ContactData datos de contacto
type ContactData struct {
	Telephone string
	Email     string
}

// DebitNoteLineData línea de nota débito
type DebitNoteLineData struct {
	ID                  string
	Description         string
	ProductCode         string
	Quantity            float64
	UnitCode            string
	UnitPrice           float64
	LineExtensionAmount float64
	Taxes               []TaxData
}

// TaxData datos de impuesto
type TaxData struct {
	TaxSchemeID   string
	TaxSchemeName string
	TaxableAmount float64
	Percent       float64
	TaxAmount     float64
}

// TotalsData totales de la nota débito
type TotalsData struct {
	LineExtensionAmount float64
	TaxExclusiveAmount  float64
	TaxInclusiveAmount  float64
	PayableAmount       float64
	Taxes               []TaxData
}

// DianExtensionsData extensiones DIAN
type DianExtensionsData struct {
	InvoiceAuthorization string
	AuthPeriodStart      time.Time
	AuthPeriodEnd        time.Time
	Prefix               string
	RangeFrom            int64
	RangeTo              int64
	ProviderID           string
	ProviderDV           string
	ProviderDocType      string
	SoftwareID           string
	SoftwareSecurityCode string
	QRCode               string
}

// ============================================================================
// Template Types
// ============================================================================

// DebitNoteTemplateData datos para template
type DebitNoteTemplateData struct {
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
	DebitNoteNumber    string
	CUDE               string
	Environment        string
	IssueDate          string
	IssueTime          string
	DebitNoteTypeCode  string
	Note               string
	CurrencyCode       string
	LineCount          int

	// Billing Reference
	BillingReference *BillingReferenceTemplateData

	// Parties
	Supplier PartyTemplateData
	Customer PartyTemplateData

	// Delivery (opcional)
	Delivery *DeliveryTemplateData

	// Allowances/Charges (opcional)
	AllowanceCharges []AllowanceChargeTemplateData

	// Totals
	LineExtensionAmount string
	TaxExclusiveAmount  string
	TaxInclusiveAmount  string
	PayableAmount       string

	// Lines
	DebitNoteLines []DebitNoteLineTemplateData
	TaxTotals      []TaxTotalTemplateData
}

// BillingReferenceTemplateData referencia a factura
type BillingReferenceTemplateData struct {
	ID        string
	UUID      string
	IssueDate string
}

// PartyTemplateData parte (supplier/customer)
type PartyTemplateData struct {
	AdditionalAccountID        string
	PartyName                  string
	IndustryClassificationCode string
	Address                    AddressTemplateData
	TaxScheme                  TaxSchemeTemplateData
	LegalEntity                LegalEntityTemplateData
	Contact                    ContactTemplateData
}

// AddressTemplateData dirección
type AddressTemplateData struct {
	ID                   string
	CityName             string
	PostalZone           string
	CountrySubentity     string
	CountrySubentityCode string
	Line                 string
	CountryCode          string
	CountryName          string
}

// TaxSchemeTemplateData esquema tributario
type TaxSchemeTemplateData struct {
	RegistrationName    string
	CompanyID           string
	CompanyIDSchemeID   string
	CompanyIDSchemeName string
	TaxLevelCode        string
	ID                  string
	Name                string
}

// LegalEntityTemplateData entidad legal
type LegalEntityTemplateData struct {
	RegistrationName            string
	CompanyID                   string
	CompanyIDSchemeID           string
	CompanyIDSchemeName         string
	CorporateRegistrationScheme string
}

// ContactTemplateData contacto
type ContactTemplateData struct {
	Telephone string
	Email     string
}

// DebitNoteLineTemplateData línea de nota débito
type DebitNoteLineTemplateData struct {
	ID                    string
	UnitCode              string
	Quantity              string
	LineExtensionAmount   string
	FreeOfChargeIndicator string
	CurrencyID            string
	Item                  ItemTemplateData
	Price                 PriceTemplateData
}

// ItemTemplateData item/producto
type ItemTemplateData struct {
	Description      string
	StandardItemID   ItemIDTemplateData
	AdditionalItemID ItemIDTemplateData
}

// ItemIDTemplateData identificación de item
type ItemIDTemplateData struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// PriceTemplateData precio
type PriceTemplateData struct {
	Amount       string
	BaseQuantity string
}

// TaxTotalTemplateData total de impuestos
type TaxTotalTemplateData struct {
	TaxAmount    string
	CurrencyID   string
	TaxSubtotals []TaxSubtotalTemplateData
}

// TaxSubtotalTemplateData subtotal de impuesto
type TaxSubtotalTemplateData struct {
	TaxableAmount string
	TaxAmount     string
	CurrencyID    string
	Percent       string
	TaxCategory   TaxCategoryTemplateData
}

// TaxCategoryTemplateData categoría de impuesto
type TaxCategoryTemplateData struct {
	ID   string
	Name string
}
