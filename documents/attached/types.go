package attached

import "time"

// PartyData datos de una parte (emisor o receptor)
type PartyData struct {
	RegistrationName string // Nombre o razón social
	CompanyID        string // NIT o identificación
	SchemeID         string // Dígito de verificación
	SchemeName       string // Tipo de documento (31=NIT, 13=CC, etc)
	TaxLevelCode     string // Código de responsabilidad fiscal (O-13, O-47, etc)
	TaxSchemeID      string // ID del esquema de impuestos (01=IVA)
	TaxSchemeName    string // Nombre del esquema de impuestos
}

// ApplicationResponseData datos del ApplicationResponse de DIAN
type ApplicationResponseData struct {
	InvoiceID            string    // Número de la factura
	CUFE                 string    // CUFE de la factura
	IssueDate            time.Time // Fecha de emisión de la factura
	ResponseXML          string    // XML completo del ApplicationResponse
	ValidationResultCode string    // Código de resultado (02 = Validado)
	ValidationDate       string    // Fecha de validación (YYYY-MM-DD)
	ValidationTime       string    // Hora de validación (HH:MM:SS-07:00)
}
