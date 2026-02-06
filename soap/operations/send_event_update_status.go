package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendEventUpdateStatus envía un evento de documento (acuse, rechazo, aceptación)
//
// Los eventos permiten al receptor notificar al emisor sobre el estado del documento:
//   - Acuse de recibo (030)
//   - Aceptación expresa (032)
//   - Aceptación tácita (033)
//   - Rechazo (031)
//   - Reclamo (034)
//
// Parámetros:
//   - req: Datos del evento (FileName, ContentFile en base64)
//
// Retorna:
//   - SendEventResponse con TrackId del evento
//   - error si falla la comunicación
func SendEventUpdateStatus(transport Transport, certPath, keyPath, url, action string, req *types.SendEventRequest) (*types.SendEventResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendEventUpdateStatus: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendEventUpdateStatus: failed to generate security header: %w", err)
	}

	body := envelope.BuildSendEventBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendEventUpdateStatus: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("SendEventUpdateStatus: failed to parse response: %w", err)
	}

	if soapResp.Body.SendEventResponse == nil {
		return nil, fmt.Errorf("SendEventUpdateStatus: unexpected response format")
	}

	return response.ToSendEventResponse(&soapResp.Body.SendEventResponse.Result), nil
}
