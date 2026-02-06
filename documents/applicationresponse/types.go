package applicationresponse

import "time"

// ApplicationResponseData datos principales del ApplicationResponse
type ApplicationResponseData struct {
	// Identificación
	ID                 string
	CUDE               string
	IssueDate          time.Time
	IssueTime          string
	ProfileExecutionID string // "1" = Producción, "2" = Habilitación

	// Partes
	SenderParty   PartyData
	ReceiverParty PartyData

	// Respuesta
	ResponseCode string   // "02" = Validado, otros códigos para errores
	Descriptions []string // Mensajes descriptivos

	// Referencia al documento validado
	DocumentReference DocumentReferenceData

	// Notas adicionales
	Notes []string
}

// PartyData datos de una parte (DIAN o empresa)
type PartyData struct {
	RegistrationName string // Nombre o razón social
	CompanyID        string // NIT o identificación
	SchemeID         string // Esquema de identificación (31=NIT, etc)
	SchemeName       string // Nombre del esquema
	TaxLevelCode     string // Código de responsabilidad fiscal (O-13, O-47, etc)
	TaxSchemeID      string // ID del esquema de impuestos (01=IVA)
	TaxSchemeName    string // Nombre del esquema de impuestos
}

// DocumentReferenceData referencia al documento validado
type DocumentReferenceData struct {
	ID               string    // Número del documento (factura, nota, etc)
	UUID             string    // CUFE/CUDE del documento
	IssueDate        time.Time // Fecha de emisión del documento
	DocumentTypeCode string    // Código de tipo de documento (01=Factura, 91=NC, 92=ND)
	DocumentType     string    // Tipo de documento en texto

	// Resultado de verificación
	ValidationResult *ValidationResultData
}

// ValidationResultData resultado de la verificación DIAN
type ValidationResultData struct {
	ValidatorID          string    // "Unidad Especial Dirección de Impuestos y Aduanas Nacionales"
	ValidationResultCode string    // "02" = Validado
	ValidationDate       time.Time // Fecha de validación
	ValidationTime       string    // Hora de validación (HH:MM:SS-05:00)
	ValidateProcess      string    // Proceso de validación (opcional)
	ValidateTool         string    // Herramienta de validación (opcional)
	ValidateToolVersion  string    // Versión de herramienta (opcional)
}

// ============================================================================
// Template Types
// ============================================================================

// ApplicationResponseTemplateData datos para template
type ApplicationResponseTemplateData struct {
	// DIAN Extensions (si aplica)
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

	// Header
	ProfileExecutionID string
	ID                 string
	CUDE               string
	Environment        string
	IssueDate          string
	IssueTime          string

	// Notes
	Notes []string

	// Parties
	SenderParty   PartyTemplateData
	ReceiverParty PartyTemplateData

	// Response
	ResponseCode string
	Descriptions []string

	// Document Reference
	DocumentReference DocumentReferenceTemplateData
}

// PartyTemplateData datos de parte para template
type PartyTemplateData struct {
	RegistrationName string
	CompanyID        string
	SchemeID         string
	SchemeName       string
	TaxLevelCode     string
	TaxSchemeID      string
	TaxSchemeName    string
}

// DocumentReferenceTemplateData referencia a documento para template
type DocumentReferenceTemplateData struct {
	ID               string
	UUID             string
	IssueDate        string
	DocumentTypeCode string
	DocumentType     string

	// Validation Result
	ValidationResult *ValidationResultTemplateData
}

// ValidationResultTemplateData resultado de validación para template
type ValidationResultTemplateData struct {
	ValidatorID          string
	ValidationResultCode string
	ValidationDate       string
	ValidationTime       string
	ValidateProcess      string
	ValidateTool         string
	ValidateToolVersion  string
}
