package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetAcquirer consulta información del adquiriente por NIT/Cédula
//
// Permite validar y obtener datos del comprador registrado en DIAN.
//
// Parámetros:
//   - req: NIT del emisor e IdentificationNumber del adquiriente
//
// Retorna:
//   - GetAcquirerResponse con datos del adquiriente
//   - error si falla la comunicación
func GetAcquirer(transport Transport, certPath, keyPath, url, action string, req *types.GetAcquirerRequest) (*types.GetAcquirerResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetAcquirer: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetAcquirer: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetAcquirerBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetAcquirer: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("GetAcquirer: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetAcquirer: failed to parse response: %w", err)
	}

	if soapResp.Body.GetAcquirerResponse == nil {
		return nil, fmt.Errorf("GetAcquirer: unexpected response format")
	}

	return response.ToGetAcquirerResponse(&soapResp.Body.GetAcquirerResponse.Result), nil
}
