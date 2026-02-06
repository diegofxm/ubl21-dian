package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetReferenceNotes consulta las notas crédito/débito relacionadas a un documento
//
// Permite obtener todas las notas de ajuste vinculadas a una factura específica.
//
// Parámetros:
//   - req: CUFE/CUDE del documento
//
// Retorna:
//   - GetReferenceNotesResponse con lista de notas relacionadas
//   - error si falla la comunicación
func GetReferenceNotes(transport Transport, certPath, keyPath, url, action string, req *types.GetReferenceNotesRequest) (*types.GetReferenceNotesResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetReferenceNotes: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetReferenceNotes: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetReferenceNotesBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetReferenceNotes: failed to build request body")
	}

	// 3. Crear envelope
	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	// 4. Enviar request
	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, fmt.Errorf("GetReferenceNotes: %w", err)
	}

	// 5. Parsear respuesta
	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetReferenceNotes: failed to parse response: %w", err)
	}

	if soapResp.Body.GetReferenceNotesResponse == nil {
		return nil, fmt.Errorf("GetReferenceNotes: unexpected response format")
	}

	return response.ToGetReferenceNotesResponse(&soapResp.Body.GetReferenceNotesResponse.Result), nil
}
