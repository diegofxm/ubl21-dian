package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetNumberingRange consulta los rangos de numeración autorizados para un NIT
//
// Esta operación permite verificar los rangos de numeración activos
// asignados por la DIAN para facturación electrónica.
//
// Parámetros:
//   - req: NIT del emisor y SoftwareID
//
// Retorna:
//   - GetNumberingRangeResponse con lista de rangos activos
//   - error si falla la comunicación
func GetNumberingRange(transport Transport, certPath, keyPath, url, action string, req *types.GetNumberingRangeRequest) (*types.GetNumberingRangeResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetNumberingRange: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetNumberingRange: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetNumberingRangeBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetNumberingRange: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("GetNumberingRange: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetNumberingRange: failed to parse response: %w", err)
	}

	if soapResp.Body.GetNumberingRangeResponse == nil {
		return nil, fmt.Errorf("GetNumberingRange: unexpected response format")
	}

	return response.ToGetNumberingRangeResponse(&soapResp.Body.GetNumberingRangeResponse.Result), nil
}
