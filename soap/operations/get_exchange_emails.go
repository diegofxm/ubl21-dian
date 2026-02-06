package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetExchangeEmails consulta los correos configurados para intercambio de documentos
//
// Retorna la lista de emails autorizados para envío/recepción de documentos electrónicos.
//
// Parámetros:
//   - req: NIT del emisor
//
// Retorna:
//   - GetExchangeEmailsResponse con lista de emails
//   - error si falla la comunicación
func GetExchangeEmails(transport Transport, certPath, keyPath, url, action string, req *types.GetExchangeEmailsRequest) (*types.GetExchangeEmailsResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetExchangeEmails: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetExchangeEmails: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetExchangeEmailsBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetExchangeEmails: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("GetExchangeEmails: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetExchangeEmails: failed to parse response: %w", err)
	}

	if soapResp.Body.GetExchangeEmailsResponse == nil {
		return nil, fmt.Errorf("GetExchangeEmails: unexpected response format")
	}

	return response.ToGetExchangeEmailsResponse(&soapResp.Body.GetExchangeEmailsResponse.Result), nil
}
