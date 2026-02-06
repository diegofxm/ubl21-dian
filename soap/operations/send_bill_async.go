package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendBillAsync envía una factura de forma asíncrona a DIAN
//
// Este método envía el documento y retorna un TrackId para consultar
// el estado después usando GetStatus. Es el método recomendado para producción.
//
// Parámetros:
//   - req: Datos de la factura (FileName, ContentFile en base64)
//
// Retorna:
//   - SendBillAsyncResponse con TrackId (XmlDocumentKey) para consultar estado
//   - error si falla la comunicación
func SendBillAsync(transport Transport, certPath, keyPath, url, action string, req *types.SendBillAsyncRequest) (*types.SendBillAsyncResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendBillAsync: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendBillAsync: failed to generate security header: %w", err)
	}

	body := envelope.BuildSendBillAsyncBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendBillAsync: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("SendBillAsync: %w", err)
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("SendBillAsync: failed to parse response: %w", err)
	}

	if soapResp.Body.SendBillAsyncResponse == nil {
		return nil, fmt.Errorf("SendBillAsync: unexpected response format")
	}

	return response.ToSendBillAsyncResponse(&soapResp.Body.SendBillAsyncResponse.Result), nil
}
