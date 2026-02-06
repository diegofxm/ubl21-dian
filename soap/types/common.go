package types

import "time"

// Config configuración del cliente SOAP
type Config struct {
	Environment Environment
	Certificate string        // Ruta al certificado PEM
	PrivateKey  string        // Ruta a la clave privada PEM (opcional si está en Certificate)
	Timeout     time.Duration // Timeout para requests HTTP
}

// Environment representa el ambiente de DIAN
type Environment string

const (
	// Produccion ambiente de producción
	Produccion Environment = "produccion"
	// Habilitacion ambiente de habilitación/pruebas
	Habilitacion Environment = "habilitacion"
)

// Response respuesta genérica de DIAN
type Response struct {
	IsValid           bool
	StatusCode        string
	StatusDescription string
	StatusMessage     string
	ErrorMessages     []ErrorMessage
	XmlDocumentKey    string // TrackId para consultas GetStatus
	XmlBase64Bytes    string // XML firmado en base64
	ZipKey            string // TrackId para SendBillAsync (lotes)
}

// ErrorMessage mensaje de error de DIAN
type ErrorMessage struct {
	Code        string
	Description string
}
