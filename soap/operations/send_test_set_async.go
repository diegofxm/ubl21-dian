package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendTestSetAsync envía una factura al set de pruebas de DIAN
//
// Este método se usa durante el proceso de certificación ante DIAN.
// Es obligatorio antes de poder usar el ambiente de producción.
//
// Parámetros:
//   - req: Datos de la factura (FileName, ContentFile, TestSetId)
//
// Retorna:
//   - SendTestSetAsyncResponse con resultado de validación del set de pruebas
//   - error si falla la comunicación o validación
func SendTestSetAsync(transport Transport, certPath, keyPath, url, action string, req *types.SendTestSetAsyncRequest) (*types.SendTestSetAsyncResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendTestSetAsync: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendTestSetAsync: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildSendTestSetAsyncBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendTestSetAsync: failed to build request body")
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
		return nil, fmt.Errorf("SendTestSetAsync: failed to parse response: %w", err)
	}

	if soapResp.Body.SendTestSetAsyncResponse == nil {
		return nil, fmt.Errorf("SendTestSetAsync: unexpected response format")
	}

	return response.ToSendTestSetAsyncResponse(&soapResp.Body.SendTestSetAsyncResponse.Result), nil
}
