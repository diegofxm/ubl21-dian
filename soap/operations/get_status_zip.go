package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetStatusZip consulta el estado y descarga el ZIP con ApplicationResponse
//
// Similar a GetStatus pero retorna el ApplicationResponse completo en formato ZIP.
// Útil para obtener la respuesta firmada por DIAN en su formato original.
//
// Parámetros:
//   - req: TrackId (XmlDocumentKey) recibido del envío asíncrono
//
// Retorna:
//   - GetStatusZipResponse con ZIP en base64 (ContentFile)
//   - error si falla la comunicación
func GetStatusZip(transport Transport, certPath, keyPath, url, action string, req *types.GetStatusZipRequest) (*types.GetStatusZipResponse, error) {
	// 1. Crear security header
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetStatusZip: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetStatusZip: failed to generate security header: %w", err)
	}

	// 2. Crear body
	body := envelope.BuildGetStatusZipBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetStatusZip: failed to build request body")
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
		return nil, fmt.Errorf("GetStatusZip: failed to parse response: %w", err)
	}

	if soapResp.Body.GetStatusZipResponse == nil {
		return nil, fmt.Errorf("GetStatusZip: unexpected response format")
	}

	return response.ToGetStatusZipResponse(soapResp.Body.GetStatusZipResponse), nil
}
