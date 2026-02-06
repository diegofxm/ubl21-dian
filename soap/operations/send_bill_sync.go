package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendBillSync envía una factura de forma síncrona a DIAN
//
// Este método envía el documento y espera la respuesta completa de validación
// en el mismo request. Es útil para desarrollo y testing.
//
// Parámetros:
//   - req: Datos de la factura (FileName, ContentFile en base64)
//
// Retorna:
//   - SendBillSyncResponse con IsValid, StatusCode, XmlDocumentKey, XmlBase64Bytes
//   - error si falla la comunicación o DIAN rechaza
// Transport interface para evitar ciclo de importación
type Transport interface {
	Send(soapXML string) ([]byte, error)
}

func SendBillSync(transport Transport, certPath, keyPath, url, action string, req *types.SendBillSyncRequest) (*types.SendBillSyncResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendBillSync: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendBillSync: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildSendBillSyncBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendBillSync: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("SendBillSync: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("SendBillSync: failed to parse response: %w", err)
	}

	if soapResp.Body.SendBillSyncResponse == nil {
		return nil, fmt.Errorf("SendBillSync: unexpected response format")
	}

	return response.ToSendBillSyncResponse(&soapResp.Body.SendBillSyncResponse.Result), nil
}
