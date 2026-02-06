package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// GetXmlByDocumentKey descarga el XML de un documento por CUFE/CUDE
//
// Permite recuperar el documento firmado desde DIAN usando su CUFE (Código Único
// de Factura Electrónica) o CUDE (Código Único de Documento Electrónico).
//
// Parámetros:
//   - req: TrackId o CUFE/CUDE del documento
//
// Retorna:
//   - GetXmlByDocumentKeyResponse con XML completo en base64
//   - error si falla la comunicación o documento no existe
func GetXmlByDocumentKey(transport Transport, certPath, keyPath, url, action string, req *types.GetXmlByDocumentKeyRequest) (*types.GetXmlByDocumentKeyResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("GetXmlByDocumentKey: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("GetXmlByDocumentKey: failed to generate security header: %w", err)
	}

	body := envelope.BuildGetXmlByDocumentKeyBody(req)
	if body == "" {
		return nil, fmt.Errorf("GetXmlByDocumentKey: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("GetXmlByDocumentKey: failed to parse response: %w", err)
	}

	if soapResp.Body.GetXmlByDocumentKeyResponse == nil {
		return nil, fmt.Errorf("GetXmlByDocumentKey: unexpected response format")
	}

	return response.ToGetXmlByDocumentKeyResponse(soapResp.Body.GetXmlByDocumentKeyResponse), nil
}
