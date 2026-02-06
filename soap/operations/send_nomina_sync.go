package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendNominaSync envía nómina electrónica de forma síncrona
//
// Método para envío de documentos de nómina electrónica.
// Módulo separado del sistema de facturación.
//
// Parámetros:
//   - req: Datos de la nómina (FileName, ContentFile en base64)
//
// Retorna:
//   - SendNominaSyncResponse con validación completa
//   - error si falla la comunicación o validación
func SendNominaSync(transport Transport, certPath, keyPath, url, action string, req *types.SendNominaSyncRequest) (*types.SendNominaSyncResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendNominaSync: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendNominaSync: failed to generate security header: %w", err)
	}

	body := envelope.BuildSendNominaSyncBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendNominaSync: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("SendNominaSync: failed to parse response: %w", err)
	}

	if soapResp.Body.SendNominaSyncResponse == nil {
		return nil, fmt.Errorf("SendNominaSync: unexpected response format")
	}

	return response.ToSendNominaSyncResponse(&soapResp.Body.SendNominaSyncResponse.Result), nil
}
