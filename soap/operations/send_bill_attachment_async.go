package operations

import (
	"fmt"

	"github.com/diegofxm/ubl21-dian/soap/envelope"
	"github.com/diegofxm/ubl21-dian/soap/response"
	"github.com/diegofxm/ubl21-dian/soap/security"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// SendBillAttachmentAsync envía documentos soporte (anexos PDF, imágenes)
//
// Permite adjuntar archivos adicionales a facturas electrónicas.
//
// Parámetros:
//   - req: Datos del anexo (FileName, ContentFile en base64)
//
// Retorna:
//   - SendBillAttachmentAsyncResponse con TrackId
//   - error si falla la comunicación
func SendBillAttachmentAsync(transport Transport, certPath, keyPath, url, action string, req *types.SendBillAttachmentAsyncRequest) (*types.SendBillAttachmentAsyncResponse, error) {
	secHeader, err := security.NewHeader(certPath, keyPath, url, action)
	if err != nil {
		return nil, fmt.Errorf("SendBillAttachmentAsync: failed to create security header: %w", err)
	}

	securityXML, err := secHeader.Generate()
	if err != nil {
		return nil, fmt.Errorf("SendBillAttachmentAsync: failed to generate security header: %w", err)
	}

	body := envelope.BuildSendBillAttachmentAsyncBody(req)
	if body == "" {
		return nil, fmt.Errorf("SendBillAttachmentAsync: failed to build request body")
	}

	env := envelope.New(securityXML, body)
	soapXML := env.Build()

	respXML, err := transport.Send(soapXML)
	if err != nil {
		return nil, err
	}

	soapResp, err := response.Parse(respXML)
	if err != nil {
		return nil, fmt.Errorf("SendBillAttachmentAsync: failed to parse response: %w", err)
	}

	if soapResp.Body.SendBillAttachmentAsyncResponse == nil {
		return nil, fmt.Errorf("SendBillAttachmentAsync: unexpected response format")
	}

	return response.ToSendBillAttachmentAsyncResponse(&soapResp.Body.SendBillAttachmentAsyncResponse.Result), nil
}
