package types

// GetStatusRequest request para consultar estado
type GetStatusRequest struct {
	TrackId string // ID de seguimiento retornado por DIAN
}

// GetStatusResponse respuesta de GetStatus
type GetStatusResponse struct {
	Response
}

// GetStatusZipRequest request para descargar ZIP de respuesta
type GetStatusZipRequest struct {
	TrackId string
}

// GetStatusZipResponse respuesta de GetStatusZip
type GetStatusZipResponse struct {
	ZipKey        string
	ContentFile   string // ZIP en base64
	StatusCode    string
	StatusMessage string
}

// GetStatusEventRequest request para consultar estado de evento
type GetStatusEventRequest struct {
	TrackId string
}

// GetStatusEventResponse respuesta de GetStatusEvent
type GetStatusEventResponse struct {
	Response
}

// GetXmlByDocumentKeyRequest request para obtener XML por CUFE
type GetXmlByDocumentKeyRequest struct {
	TrackId string
}

// GetXmlByDocumentKeyResponse respuesta de GetXmlByDocumentKey
type GetXmlByDocumentKeyResponse struct {
	XmlBase64Bytes string
	StatusCode     string
	StatusMessage  string
}

// GetReferenceNotesRequest request para obtener notas de referencia
type GetReferenceNotesRequest struct {
	DocumentKey string // CUFE/CUDE del documento
}

// GetReferenceNotesResponse respuesta de GetReferenceNotes
type GetReferenceNotesResponse struct {
	Notes         []ReferenceNote
	StatusCode    string
	StatusMessage string
}

// ReferenceNote nota de referencia (crédito/débito)
type ReferenceNote struct {
	DocumentKey string
	IssueDate   string
	Type        string // CreditNote, DebitNote
}

// GetDocumentInfoRequest request para obtener información de documento
type GetDocumentInfoRequest struct {
	DocumentKey string // CUFE/CUDE del documento
}

// GetDocumentInfoResponse respuesta de GetDocumentInfo
type GetDocumentInfoResponse struct {
	DocumentKey   string
	IssueDate     string
	Status        string
	Events        []DocumentEvent
	StatusCode    string
	StatusMessage string
}

// DocumentEvent evento de documento
type DocumentEvent struct {
	EventType   string
	EventDate   string
	Description string
}
