package invoice

import "time"

// SupplierData datos del emisor de la factura
type SupplierData struct {
	PersonType string // "1" = Persona Jurídica, "2" = Persona Natural
	Party      PartyData
}

// CustomerData datos del receptor de la factura
type CustomerData struct {
	PersonType string // "1" = Persona Jurídica, "2" = Persona Natural
	Party      PartyData
}

// PartyData datos de una parte (emisor o receptor)
type PartyData struct {
	ID               string
	DV               string // Dígito de verificación
	DocumentType     string // "31" = NIT, "13" = Cédula, "22" = CE, etc.
	Name             string
	TaxLevelCode     string // "O-13" = Responsable IVA, "O-47" = Régimen simple, etc.
	TaxSchemeID      string // "01" = IVA, "04" = INC, etc.
	TaxSchemeName    string
	ResolutionPrefix string // Prefijo de la resolución (ej: "SETP", "FESG")
	IndustryCodes    string // Códigos CIIU separados por ";" (ej: "2511;4330;4761;6202")
	Address          *AddressData
	Contact          *ContactData
}

// AddressData datos de dirección
type AddressData struct {
	ID                   string // Código del municipio (ej: "76520" para Palmira)
	CityName             string
	PostalZone           string
	CountrySubentity     string // Departamento/Estado
	CountrySubentityCode string // Código del departamento
	AddressLine          string
	CountryCode          string // "CO" para Colombia
	CountryName          string
}

// ContactData datos de contacto
type ContactData struct {
	Telephone string
	Email     string
}

// InvoiceLineData datos de una línea de factura
type InvoiceLineData struct {
	ID                  string
	Description         string
	ProductCode         string  // Código del producto para StandardItemIdentification
	Quantity            float64
	UnitCode            string  // "EA" = Each, "KGM" = Kilogram, etc.
	UnitPrice           float64
	LineExtensionAmount float64 // Cantidad * Precio unitario
	Taxes               []TaxData
}

// TaxData datos de un impuesto
type TaxData struct {
	TaxSchemeID   string  // "01" = IVA, "04" = INC, "03" = ICA
	TaxSchemeName string  // "IVA", "INC", "ICA"
	TaxableAmount float64 // Base gravable
	Percent       float64 // Porcentaje del impuesto
	TaxAmount     float64 // Monto del impuesto
}

// TotalsData totales de la factura
type TotalsData struct {
	LineExtensionAmount float64   // Suma de líneas sin impuestos
	TaxExclusiveAmount  float64   // Total sin impuestos
	TaxInclusiveAmount  float64   // Total con impuestos
	PayableAmount       float64   // Total a pagar
	Taxes               []TaxData // Totales de impuestos agrupados
}

// DianExtensionsData datos de las extensiones DIAN
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
	AuthProviderID       string
	AuthProviderDV       string
	QRCode               string
}

// ============================================================================
// Template Types - Tipos para generación de XML con templates
// ============================================================================

// InvoiceTemplateData contiene todos los datos para generar una factura con templates
type InvoiceTemplateData struct {
	// DianExtensions
	InvoiceAuthorization string
	AuthPeriodStartDate  string
	AuthPeriodEndDate    string
	Prefix               string
	From                 string
	To                   string
	ProviderID           string
	ProviderSchemeID     string // "4" = NIT, "1" = Cédula
	ProviderSchemeName   string // "31" = NIT, "13" = Cédula
	SoftwareID           string
	SecurityCode         string
	QRCode               string

	// Invoice Header
	ProfileExecutionID string
	InvoiceNumber      string
	CUFE               string
	Environment        string
	IssueDate          string
	IssueTime          string
	DueDate            string
	InvoiceTypeCode    string
	Notes              []string
	CurrencyCode       string
	LineCount          int

	// Parties
	Supplier PartyTemplateData
	Customer PartyTemplateData

	// Delivery (opcional)
	Delivery *DeliveryTemplateData

	// Payment
	PaymentMeans []PaymentMeansTemplateData

	// Monetary Totals
	PrepaidAmount        string
	LineExtensionAmount  string
	TaxExclusiveAmount   string
	TaxInclusiveAmount   string
	AllowanceTotalAmount string
	ChargeTotalAmount    string
	PayableAmount        string

	// Lines
	InvoiceLines []InvoiceLineTemplateData

	// Optional sections
	InvoicePeriod        *InvoicePeriodTemplateData
	PrepaidPayment       *PrepaidPaymentTemplateData
	TaxTotals            []TaxTotalTemplateData
	AllowanceCharges     []AllowanceChargeTemplateData
	WithholdingTaxTotals []WithholdingTaxTemplateData
	OrderReference       *OrderReferenceTemplateData
	BillingReference     *BillingReferenceTemplateData
}

// PartyTemplateData representa un Supplier o Customer
type PartyTemplateData struct {
	AdditionalAccountID        string
	PartyName                  string
	IndustryClassificationCode string
	Address                    AddressTemplateData
	TaxScheme                  TaxSchemeTemplateData
	LegalEntity                LegalEntityTemplateData
	Contact                    ContactTemplateData
}

// AddressTemplateData representa una dirección
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

// TaxSchemeTemplateData representa el esquema tributario
type TaxSchemeTemplateData struct {
	RegistrationName    string
	CompanyID           string
	CompanyIDSchemeID   string
	CompanyIDSchemeName string
	TaxLevelCode        string
	ID                  string
	Name                string
}

// LegalEntityTemplateData representa la entidad legal
type LegalEntityTemplateData struct {
	RegistrationName            string
	CompanyID                   string
	CompanyIDSchemeID           string
	CompanyIDSchemeName         string
	CorporateRegistrationScheme string
}

// ContactTemplateData representa información de contacto
type ContactTemplateData struct {
	Telephone string
	Email     string
}

// DeliveryTemplateData representa información de entrega
type DeliveryTemplateData struct {
	ActualDeliveryDate string
	Address            AddressTemplateData
}

// PaymentMeansTemplateData representa medios de pago
type PaymentMeansTemplateData struct {
	ID      string
	Code    string
	DueDate string
}

// InvoiceLineTemplateData representa una línea de factura
type InvoiceLineTemplateData struct {
	ID                    string
	UnitCode              string
	Quantity              string
	LineExtensionAmount   string
	FreeOfChargeIndicator string
	CurrencyID            string
	Item                  ItemTemplateData
	Price                 PriceTemplateData
}

// ItemTemplateData representa un item/producto
type ItemTemplateData struct {
	Description      string
	StandardItemID   ItemIDTemplateData
	AdditionalItemID ItemIDTemplateData
}

// ItemIDTemplateData representa identificación de item
type ItemIDTemplateData struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// PriceTemplateData representa precio
type PriceTemplateData struct {
	Amount       string
	BaseQuantity string
}

// TaxTotalTemplateData representa totales de impuestos
type TaxTotalTemplateData struct {
	TaxAmount    string
	CurrencyID   string
	TaxSubtotals []TaxSubtotalTemplateData
}

// TaxSubtotalTemplateData representa subtotal de un impuesto
type TaxSubtotalTemplateData struct {
	TaxableAmount string
	TaxAmount     string
	CurrencyID    string
	Percent       string
	TaxCategory   TaxCategoryTemplateData
}

// TaxCategoryTemplateData representa categoría de impuesto
type TaxCategoryTemplateData struct {
	ID         string // "01" = IVA, "04" = INC, "03" = ICA
	Name       string
	TaxScheme  string
}

// AllowanceChargeTemplateData representa descuentos o cargos
type AllowanceChargeTemplateData struct {
	ID                   string
	ChargeIndicator      string // "false" = descuento, "true" = cargo
	AllowanceChargeReason string
	Amount               string
	CurrencyID           string
	BaseAmount           string
}

// WithholdingTaxTemplateData representa retenciones
type WithholdingTaxTemplateData struct {
	TaxAmount    string
	CurrencyID   string
	TaxSubtotals []TaxSubtotalTemplateData
}

// OrderReferenceTemplateData representa referencia a orden de compra
type OrderReferenceTemplateData struct {
	ID string
}

// InvoicePeriodTemplateData representa período de facturación
type InvoicePeriodTemplateData struct {
	StartDate string
	EndDate   string
}

// PrepaidPaymentTemplateData representa pago anticipado
type PrepaidPaymentTemplateData struct {
	Amount string
}

// BillingReferenceTemplateData representa referencia a factura anterior (para notas)
type BillingReferenceTemplateData struct {
	InvoiceDocumentReference InvoiceDocumentReferenceTemplateData
}

// InvoiceDocumentReferenceTemplateData representa referencia a documento de factura
type InvoiceDocumentReferenceTemplateData struct {
	ID       string
	UUID     string
	IssueDate string
}
