package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetDocumentInfo consulta información completa de un documento por su CUFE/CUDE
//
// Retorna todos los detalles del documento incluyendo estado, fechas, y metadata.
//
// Parámetros:
//   - req: CUFE/CUDE del documento
//
// Retorna:
//   - GetDocumentInfoResponse con información completa
//   - error si falla la comunicación
func GetDocumentInfo(transport Transport, certPath, keyPath, url, action string, req *types.GetDocumentInfoRequest) (*types.GetDocumentInfoResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetDocumentInfo: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetDocumentInfo: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetDocumentInfoBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetDocumentInfo: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("GetDocumentInfo: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetDocumentInfo: failed to parse response: %w", err)
	}

	if soapResp.Body.GetDocumentInfoResponse == nil {
		return nil, fmt.Errorf("GetDocumentInfo: unexpected response format")
	}

	return response.ToGetDocumentInfoResponse(&soapResp.Body.GetDocumentInfoResponse.Result), nil
}
