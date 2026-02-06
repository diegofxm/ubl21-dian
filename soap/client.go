package soap

import (
	"fmt"
	"time"

	"github.com/diegofxm/ubl21-dian/soap/operations"
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// Client cliente SOAP para DIAN con arquitectura modular
//
// Este cliente orquesta las llamadas a las 15 operaciones SOAP de DIAN,
// delegando la implementación a módulos especializados:
//   - operations/: Implementación de cada operación SOAP
//   - security/: Generación de headers WS-Security
//   - envelope/: Construcción de SOAP envelopes
//   - response/: Parsing de respuestas XML
//   - transport/: Comunicación HTTP/HTTPS con mTLS
type Client struct {
	config    *types.Config
	transport *Transport
	url       string
}

// NewClient crea un nuevo cliente SOAP configurado para DIAN
//
// Parámetros:
//   - config: Configuración con certificados, environment, timeout
//
// Retorna:
//   - *Client listo para usar
//   - error si falla la configuración de certificados o transporte
func NewClient(config *types.Config) (*Client, error) {
	if config.Timeout == 0 {
		config.Timeout = 180 * time.Second
	}

	url := GetURL(config.Environment)

	// Crear TLS config desde certificados
	tlsConfig, err := LoadClientTLSConfig(config.Certificate, config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create TLS config: %w", err)
	}

	transport := NewTransport(url, tlsConfig, config.Timeout)

	return &Client{
		config:    config,
		transport: transport,
		url:       url,
	}, nil
}

// ============================================================================
// GRUPO 1: OPERACIONES DE ENVÍO DE DOCUMENTOS
// ============================================================================

// SendBillSync envía una factura de forma síncrona
// Delega a operations.SendBillSync
func (c *Client) SendBillSync(req *types.SendBillSyncRequest) (*types.SendBillSyncResponse, error) {
	return operations.SendBillSync(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendBillSync, req)
}

// SendBillAsync envía una factura de forma asíncrona
// Delega a operations.SendBillAsync
func (c *Client) SendBillAsync(req *types.SendBillAsyncRequest) (*types.SendBillAsyncResponse, error) {
	return operations.SendBillAsync(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendBillAsync, req)
}

// SendTestSetAsync envía una factura al set de pruebas de DIAN
// Delega a operations.SendTestSetAsync
func (c *Client) SendTestSetAsync(req *types.SendTestSetAsyncRequest) (*types.SendTestSetAsyncResponse, error) {
	return operations.SendTestSetAsync(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendTestSetAsync, req)
}

// SendBillAttachmentAsync envía documentos soporte (anexos)
// Delega a operations.SendBillAttachmentAsync
func (c *Client) SendBillAttachmentAsync(req *types.SendBillAttachmentAsyncRequest) (*types.SendBillAttachmentAsyncResponse, error) {
	return operations.SendBillAttachmentAsync(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendBillAttachmentAsync, req)
}

// SendNominaSync envía nómina electrónica de forma síncrona
// Delega a operations.SendNominaSync
func (c *Client) SendNominaSync(req *types.SendNominaSyncRequest) (*types.SendNominaSyncResponse, error) {
	return operations.SendNominaSync(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendNominaSync, req)
}

// ============================================================================
// GRUPO 2: OPERACIONES DE CONSULTA DE ESTADO
// ============================================================================

// GetStatus consulta el estado de un documento por TrackId
// Delega a operations.GetStatus
func (c *Client) GetStatus(req *types.GetStatusRequest) (*types.GetStatusResponse, error) {
	return operations.GetStatus(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetStatus, req)
}

// GetStatusZip consulta el estado y descarga el ZIP con ApplicationResponse
// Delega a operations.GetStatusZip
func (c *Client) GetStatusZip(req *types.GetStatusZipRequest) (*types.GetStatusZipResponse, error) {
	return operations.GetStatusZip(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetStatusZip, req)
}

// GetStatusEvent consulta el estado de un evento de documento
// Delega a operations.GetStatusEvent
func (c *Client) GetStatusEvent(req *types.GetStatusEventRequest) (*types.GetStatusEventResponse, error) {
	return operations.GetStatusEvent(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetStatusEvent, req)
}

// ============================================================================
// GRUPO 3: OPERACIONES DE EVENTOS DE DOCUMENTOS
// ============================================================================

// SendEventUpdateStatus envía un evento de documento (acuse, rechazo, aceptación)
// Delega a operations.SendEventUpdateStatus
func (c *Client) SendEventUpdateStatus(req *types.SendEventRequest) (*types.SendEventResponse, error) {
	return operations.SendEventUpdateStatus(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionSendEventUpdateStatus, req)
}

// ============================================================================
// GRUPO 4: OPERACIONES DE CONSULTAS DE INFORMACIÓN
// ============================================================================

// GetNumberingRange consulta rangos de numeración autorizados
// Delega a operations.GetNumberingRange
func (c *Client) GetNumberingRange(req *types.GetNumberingRangeRequest) (*types.GetNumberingRangeResponse, error) {
	return operations.GetNumberingRange(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetNumberingRange, req)
}

// GetXmlByDocumentKey descarga el XML de un documento por CUFE/CUDE
// Delega a operations.GetXmlByDocumentKey
func (c *Client) GetXmlByDocumentKey(req *types.GetXmlByDocumentKeyRequest) (*types.GetXmlByDocumentKeyResponse, error) {
	return operations.GetXmlByDocumentKey(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetXmlByDocumentKey, req)
}

// GetReferenceNotes consulta notas crédito/débito asociadas a una factura
// Delega a operations.GetReferenceNotes
func (c *Client) GetReferenceNotes(req *types.GetReferenceNotesRequest) (*types.GetReferenceNotesResponse, error) {
	return operations.GetReferenceNotes(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetReferenceNotes, req)
}

// GetDocumentInfo consulta información completa de un documento
// Delega a operations.GetDocumentInfo
func (c *Client) GetDocumentInfo(req *types.GetDocumentInfoRequest) (*types.GetDocumentInfoResponse, error) {
	return operations.GetDocumentInfo(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetDocumentInfo, req)
}

// GetAcquirer consulta información del adquiriente (comprador)
// Delega a operations.GetAcquirer
func (c *Client) GetAcquirer(req *types.GetAcquirerRequest) (*types.GetAcquirerResponse, error) {
	return operations.GetAcquirer(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetAcquirer, req)
}

// GetExchangeEmails consulta correos de intercambio configurados
// Delega a operations.GetExchangeEmails
func (c *Client) GetExchangeEmails(req *types.GetExchangeEmailsRequest) (*types.GetExchangeEmailsResponse, error) {
	return operations.GetExchangeEmails(c.transport, c.config.Certificate, c.config.PrivateKey, c.url, ActionGetExchangeEmails, req)
}
