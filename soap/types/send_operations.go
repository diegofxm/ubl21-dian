package types

// SendBillSyncRequest request para envío síncrono de factura
type SendBillSyncRequest struct {
	FileName    string // Nombre del archivo ZIP (ej: "FES-990000001.zip")
	ContentFile string // Contenido del ZIP en base64
}

// SendBillSyncResponse respuesta de SendBillSync
type SendBillSyncResponse struct {
	Response
}

// SendBillAsyncRequest request para envío asíncrono de factura
type SendBillAsyncRequest struct {
	FileName    string
	ContentFile string
}

// SendBillAsyncResponse respuesta de SendBillAsync
type SendBillAsyncResponse struct {
	Response
}

// SendTestSetAsyncRequest request para envío a set de pruebas
type SendTestSetAsyncRequest struct {
	FileName    string
	ContentFile string
	TestSetId   string // ID del set de pruebas de DIAN
}

// SendTestSetAsyncResponse respuesta de SendTestSetAsync
type SendTestSetAsyncResponse struct {
	Response
}

// SendBillAttachmentAsyncRequest request para envío de documento adjunto
type SendBillAttachmentAsyncRequest struct {
	FileName    string
	ContentFile string
}

// SendBillAttachmentAsyncResponse respuesta de SendBillAttachmentAsync
type SendBillAttachmentAsyncResponse struct {
	Response
}

// SendEventRequest request para envío de evento (acuse, rechazo, etc.)
type SendEventRequest struct {
	FileName    string
	ContentFile string
}

// SendEventResponse respuesta de SendEvent
type SendEventResponse struct {
	Response
}

// SendNominaSyncRequest request para envío de nómina
type SendNominaSyncRequest struct {
	FileName    string
	ContentFile string
}

// SendNominaSyncResponse respuesta de SendNominaSync
type SendNominaSyncResponse struct {
	Response
}
