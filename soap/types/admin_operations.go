package types

// GetNumberingRangeRequest request para obtener rangos de numeración
type GetNumberingRangeRequest struct {
	NIT        string // NIT del emisor (AccountCode)
	SoftwareID string // ID del software registrado
}

// GetNumberingRangeResponse respuesta de GetNumberingRange
type GetNumberingRangeResponse struct {
	Ranges        []NumberingRange
	StatusCode    string
	StatusMessage string
}

// NumberingRange rango de numeración autorizado
type NumberingRange struct {
	Prefix     string
	From       int64
	To         int64
	DateFrom   string
	DateTo     string
	Resolution string
}

// GetAcquirerRequest request para obtener información del adquiriente
type GetAcquirerRequest struct {
	NIT                  string // NIT del emisor (AccountCode)
	IdentificationNumber string // NIT/Cédula del adquiriente
}

// GetAcquirerResponse respuesta de GetAcquirer
type GetAcquirerResponse struct {
	IdentificationNumber string
	Name                 string
	Email                string
	Address              string
	StatusCode           string
	StatusMessage        string
}

// GetExchangeEmailsRequest request para obtener correos de intercambio
type GetExchangeEmailsRequest struct {
	NIT string // NIT del emisor
}

// GetExchangeEmailsResponse respuesta de GetExchangeEmails
type GetExchangeEmailsResponse struct {
	Emails        []string
	StatusCode    string
	StatusMessage string
}
