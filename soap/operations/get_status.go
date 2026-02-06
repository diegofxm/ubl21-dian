package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetStatus consulta el estado de un documento por TrackId
//
// Este método es CRÍTICO para verificar si DIAN aceptó o rechazó un documento
// enviado de forma asíncrona. Debe llamarse después de SendBillAsync o SendTestSetAsync.
//
// Parámetros:
//   - req: TrackId (XmlDocumentKey) recibido del envío asíncrono
//
// Retorna:
//   - GetStatusResponse con IsValid, StatusCode, ApplicationResponse final en XmlBase64Bytes
//   - error si falla la comunicación
func GetStatus(transport Transport, certPath, keyPath, url, action string, req *types.GetStatusRequest) (*types.GetStatusResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetStatus: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetStatus: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetStatusBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetStatus: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetStatus: failed to parse response: %w", err)
	}

	if soapResp.Body.GetStatusResponse == nil {
		return nil, fmt.Errorf("GetStatus: unexpected response format")
	}

	return response.ToGetStatusResponse(&soapResp.Body.GetStatusResponse.Result), nil
}
