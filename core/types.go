package core

import "time"

// Party representa una parte (emisor o receptor)
type Party struct {
	Name                       string
	TaxScheme                  TaxScheme
	PartyIdentification        PartyIdentification
	PartyTaxScheme             PartyTaxScheme
	PartyLegalEntity           PartyLegalEntity
	PhysicalLocation           *Address
	Contact                    *Contact
	AdditionalAccountID        string // Tipo de persona: 1=Persona Jurídica, 2=Persona Natural
	IndustryClassificationCode string // Código CIIU - Opcional
}

// PartyIdentification identificación de la parte
type PartyIdentification struct {
	ID              string // Número de identificación
	SchemeID        string // Dígito de Verificación (DV)
	SchemeName      string // Tipo de documento: 31=NIT, 13=Cédula, 22=CE, 41=Pasaporte, etc.
	SchemeVersionID string // Tipo de organización: 1=Persona Jurídica, 2=Persona Natural
}

// TaxScheme esquema tributario
type TaxScheme struct {
	ID   string // "01" = IVA, "04" = INC, etc.
	Name string
}

// PartyTaxScheme régimen tributario de la parte
type PartyTaxScheme struct {
	RegistrationName    string
	CompanyID           string
	TaxLevelCode        string // "O-13" = Responsable IVA, etc.
	RegistrationAddress *Address
	TaxScheme           TaxScheme
}

// PartyLegalEntity entidad legal
type PartyLegalEntity struct {
	RegistrationName            string
	CompanyID                   string
	CorporateRegistrationScheme *CorporateRegistrationScheme
}

// CorporateRegistrationScheme esquema de registro corporativo
type CorporateRegistrationScheme struct {
	ID   string
	Name string
}

// Address dirección
type Address struct {
	ID                   string
	CityName             string
	PostalZone           string
	CountrySubentity     string
	CountrySubentityCode string
	AddressLine          *AddressLine
	Country              Country
}

// AddressLine línea de dirección
type AddressLine struct {
	Line string
}

// Country país
type Country struct {
	IdentificationCode string // "CO" para Colombia
	Name               string
}

// Contact información de contacto
type Contact struct {
	Name           string
	Telephone      string
	ElectronicMail string
}

// MonetaryAmount monto monetario
type MonetaryAmount struct {
	Value      float64
	CurrencyID string // "COP"
}

// Quantity cantidad
type Quantity struct {
	Value    float64
	UnitCode string // "EA" = Each, "KGM" = Kilogram, etc.
}

// Item artículo/producto
type Item struct {
	Description                string
	StandardItemIdentification *StandardItemIdentification
	SellersItemIdentification  *SellersItemIdentification
	BrandName                  string
	ModelName                  string
}

// StandardItemIdentification identificación estándar del artículo
type StandardItemIdentification struct {
	ID         string
	SchemeID   string
	SchemeName string
}

// SellersItemIdentification identificación del vendedor
type SellersItemIdentification struct {
	ID string
}

// Price precio
type Price struct {
	PriceAmount  MonetaryAmount
	BaseQuantity *Quantity
}

// TaxTotal total de impuestos
type TaxTotal struct {
	TaxAmount    MonetaryAmount
	TaxSubtotals []TaxSubtotal
}

// TaxSubtotal subtotal de impuesto
type TaxSubtotal struct {
	TaxableAmount MonetaryAmount
	TaxAmount     MonetaryAmount
	TaxCategory   TaxCategory
}

// TaxCategory categoría de impuesto
type TaxCategory struct {
	Percent   float64
	TaxScheme TaxScheme
}

// AllowanceCharge descuento o cargo
type AllowanceCharge struct {
	ID                      string
	ChargeIndicator         bool // true = cargo, false = descuento
	AllowanceChargeReason   string
	MultiplierFactorNumeric float64
	Amount                  MonetaryAmount
	BaseAmount              MonetaryAmount
}

// PaymentMeans medio de pago
type PaymentMeans struct {
	ID               string
	PaymentMeansCode string // "10" = Efectivo, "48" = Transferencia, etc.
	PaymentDueDate   *time.Time
	PaymentID        string
}

// PaymentTerms términos de pago
type PaymentTerms struct {
	PaymentDueDate *time.Time
	Note           string
}
