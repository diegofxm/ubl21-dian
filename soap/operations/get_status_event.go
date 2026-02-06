package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetStatusEvent consulta el estado de un evento de documento
//
// Permite verificar si un evento (acuse, rechazo, aceptación) fue procesado correctamente.
//
// Parámetros:
//   - req: TrackId del evento
//
// Retorna:
//   - GetStatusEventResponse con estado del evento
//   - error si falla la comunicación
func GetStatusEvent(transport Transport, certPath, keyPath, url, action string, req *types.GetStatusEventRequest) (*types.GetStatusEventResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetStatusEvent: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetStatusEvent: failed to generate security header: %w", err)
	}

	body := envelope.BuildGetStatusBody(&types.GetStatusRequest{TrackId: req.TrackId})
	if body == "" {
		return nil, fmt.Errorf("GetStatusEvent: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetStatusEvent: failed to parse response: %w", err)
	}

	if soapResp.Body.GetStatusEventResponse == nil {
		return nil, fmt.Errorf("GetStatusEvent: unexpected response format")
	}

	return response.ToGetStatusEventResponse(&soapResp.Body.GetStatusEventResponse.Result), nil
}
