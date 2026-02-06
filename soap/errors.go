package soap

import "fmt"

// SOAPError error personalizado para operaciones SOAP
type SOAPError struct {
	Operation string // Nombre de la operación (SendBillSync, GetStatus, etc.)
	Code      string // Código de error
	Message   string // Mensaje de error
	Err       error  // Error original
}

// Error implementa la interfaz error
func (e *SOAPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %s (caused by: %v)", e.Operation, e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Operation, e.Code, e.Message)
}

// Unwrap permite usar errors.Unwrap
func (e *SOAPError) Unwrap() error {
	return e.Err
}

// NewSOAPError crea un nuevo error SOAP
func NewSOAPError(operation, code, message string, err error) *SOAPError {
	return &SOAPError{
		Operation: operation,
		Code:      code,
		Message:   message,
		Err:       err,
	}
}

// Errores comunes predefinidos
var (
	ErrInvalidRequest      = "INVALID_REQUEST"
	ErrSecurityHeader      = "SECURITY_HEADER_ERROR"
	ErrHTTPTransport       = "HTTP_TRANSPORT_ERROR"
	ErrResponseParsing     = "RESPONSE_PARSING_ERROR"
	ErrUnexpectedResponse  = "UNEXPECTED_RESPONSE"
	ErrDIANRejection       = "DIAN_REJECTION"
	ErrCertificateLoad     = "CERTIFICATE_LOAD_ERROR"
	ErrTimeout             = "TIMEOUT"
)

// IsSOAPError verifica si un error es un SOAPError
func IsSOAPError(err error) bool {
	_, ok := err.(*SOAPError)
	return ok
}

// GetSOAPError extrae el SOAPError de un error
func GetSOAPError(err error) *SOAPError {
	if soapErr, ok := err.(*SOAPError); ok {
		return soapErr
	}
	return nil
}
