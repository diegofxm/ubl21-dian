package response

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// SOAPEnvelope estructura para parsear respuesta SOAP
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

// SOAPBody body del SOAP
type SOAPBody struct {
	SendBillSyncResponse            *SendBillSyncResponseXML            `xml:"SendBillSyncResponse"`
	SendBillAsyncResponse           *SendBillAsyncResponseXML           `xml:"SendBillAsyncResponse"`
	SendTestSetAsyncResponse        *SendTestSetAsyncResponseXML        `xml:"SendTestSetAsyncResponse"`
	GetStatusResponse               *GetStatusResponseXML               `xml:"GetStatusResponse"`
	GetStatusZipResponse            *GetStatusZipResponseXML            `xml:"GetStatusZipResponse"`
	GetXmlByDocumentKeyResponse     *GetXmlByDocumentKeyResponseXML     `xml:"GetXmlByDocumentKeyResponse"`
	SendBillAttachmentAsyncResponse *SendBillAttachmentAsyncResponseXML `xml:"SendBillAttachmentAsyncResponse"`
	SendEventResponse               *SendEventResponseXML               `xml:"SendEventResponse"`
	SendNominaSyncResponse          *SendNominaSyncResponseXML          `xml:"SendNominaSyncResponse"`
	GetStatusEventResponse          *GetStatusEventResponseXML          `xml:"GetStatusEventResponse"`
	GetNumberingRangeResponse       *GetNumberingRangeResponseXML       `xml:"GetNumberingRangeResponse"`
	GetReferenceNotesResponse       *GetReferenceNotesResponseXML       `xml:"GetReferenceNotesResponse"`
	GetDocumentInfoResponse         *GetDocumentInfoResponseXML         `xml:"GetDocumentInfoResponse"`
	GetAcquirerResponse             *GetAcquirerResponseXML             `xml:"GetAcquirerResponse"`
	GetExchangeEmailsResponse       *GetExchangeEmailsResponseXML       `xml:"GetExchangeEmailsResponse"`
	Fault                           *SOAPFault                          `xml:"Fault"`
}

// SOAPFault error SOAP
type SOAPFault struct {
	Code   string `xml:"Code>Value"`
	Reason string `xml:"Reason>Text"`
	Detail string `xml:"Detail"`
}

// SendBillSyncResponseXML respuesta XML de SendBillSync
type SendBillSyncResponseXML struct {
	Result ResponseXML `xml:"SendBillSyncResult"`
}

// SendBillAsyncResponseXML respuesta XML de SendBillAsync
type SendBillAsyncResponseXML struct {
	Result ResponseXML `xml:"SendBillAsyncResult"`
}

// SendTestSetAsyncResponseXML respuesta XML de SendTestSetAsync
type SendTestSetAsyncResponseXML struct {
	Result ResponseXML `xml:"SendTestSetAsyncResult"`
}

// GetStatusResponseXML respuesta XML de GetStatus
type GetStatusResponseXML struct {
	Result ResponseXML `xml:"GetStatusResult"`
}

// GetStatusZipResponseXML respuesta XML de GetStatusZip
type GetStatusZipResponseXML struct {
	ZipKey        string `xml:"GetStatusZipResult>ZipKey"`
	ContentFile   string `xml:"GetStatusZipResult>ContentFile"`
	StatusCode    string `xml:"GetStatusZipResult>StatusCode"`
	StatusMessage string `xml:"GetStatusZipResult>StatusMessage"`
}

// GetXmlByDocumentKeyResponseXML respuesta XML de GetXmlByDocumentKey
type GetXmlByDocumentKeyResponseXML struct {
	XmlBase64Bytes string `xml:"GetXmlByDocumentKeyResult>XmlBase64Bytes"`
	StatusCode     string `xml:"GetXmlByDocumentKeyResult>StatusCode"`
	StatusMessage  string `xml:"GetXmlByDocumentKeyResult>StatusMessage"`
}

// SendBillAttachmentAsyncResponseXML respuesta XML de SendBillAttachmentAsync
type SendBillAttachmentAsyncResponseXML struct {
	Result ResponseXML `xml:"SendBillAttachmentAsyncResult"`
}

// SendEventResponseXML respuesta XML de SendEvent
type SendEventResponseXML struct {
	Result ResponseXML `xml:"SendEventResult"`
}

// SendNominaSyncResponseXML respuesta XML de SendNominaSync
type SendNominaSyncResponseXML struct {
	Result ResponseXML `xml:"SendNominaSyncResult"`
}

// GetStatusEventResponseXML respuesta XML de GetStatusEvent
type GetStatusEventResponseXML struct {
	Result ResponseXML `xml:"GetStatusEventResult"`
}

// GetNumberingRangeResponseXML respuesta XML de GetNumberingRange
type GetNumberingRangeResponseXML struct {
	Result ResponseXML `xml:"GetNumberingRangeResult"`
}

// GetReferenceNotesResponseXML respuesta XML de GetReferenceNotes
type GetReferenceNotesResponseXML struct {
	Result ResponseXML `xml:"GetReferenceNotesResult"`
}

// GetDocumentInfoResponseXML respuesta XML de GetDocumentInfo
type GetDocumentInfoResponseXML struct {
	Result ResponseXML `xml:"GetDocumentInfoResult"`
}

// GetAcquirerResponseXML respuesta XML de GetAcquirer
type GetAcquirerResponseXML struct {
	Result ResponseXML `xml:"GetAcquirerResult"`
}

// GetExchangeEmailsResponseXML respuesta XML de GetExchangeEmails
type GetExchangeEmailsResponseXML struct {
	Result ResponseXML `xml:"GetExchangeEmailsResult"`
}

// ResponseXML estructura común de respuesta
type ResponseXML struct {
	IsValid           string            `xml:"IsValid"`
	StatusCode        string            `xml:"StatusCode"`
	StatusDescription string            `xml:"StatusDescription"`
	StatusMessage     string            `xml:"StatusMessage"`
	ErrorMessage      []ErrorMessageXML `xml:"ErrorMessage>DianResponse"`
	XmlDocumentKey    string            `xml:"XmlDocumentKey"`
	XmlBase64Bytes    string            `xml:"XmlBase64Bytes"`
}

// ErrorMessageXML mensaje de error XML
type ErrorMessageXML struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

// Parse parsea la respuesta SOAP
func Parse(xmlData []byte) (*SOAPEnvelope, error) {
	var envelope SOAPEnvelope

	// Limpiar namespaces para facilitar el parsing
	xmlStr := string(xmlData)
	xmlStr = strings.ReplaceAll(xmlStr, "s:", "")
	xmlStr = strings.ReplaceAll(xmlStr, "a:", "")
	xmlStr = strings.ReplaceAll(xmlStr, "b:", "")
	xmlStr = strings.ReplaceAll(xmlStr, "i:", "")

	err := xml.Unmarshal([]byte(xmlStr), &envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SOAP response: %w", err)
	}

	// Verificar si hay un SOAP Fault
	if envelope.Body.Fault != nil {
		return nil, fmt.Errorf("SOAP Fault: %s - %s", envelope.Body.Fault.Code, envelope.Body.Fault.Reason)
	}

	return &envelope, nil
}
