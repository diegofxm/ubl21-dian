package envelope

import (
	"bytes"
	"text/template"

	"github.com/diegofxm/ubl21-dian/soap/types"
)

// Paths de templates para cada operación
const (
	sendBillSyncBodyPath            = "ubl21-dian/soap/templates/operations/send_bill_sync_body.tmpl"
	sendBillAsyncBodyPath           = "ubl21-dian/soap/templates/operations/send_bill_async_body.tmpl"
	sendTestSetAsyncBodyPath        = "ubl21-dian/soap/templates/operations/send_bill_async_body.tmpl" // Reutiliza async
	sendBillAttachmentAsyncBodyPath = "ubl21-dian/soap/templates/operations/send_bill_attachment_async_body.tmpl"
	sendNominaSyncBodyPath          = "ubl21-dian/soap/templates/operations/send_nomina_sync_body.tmpl"
	sendEventUpdateStatusBodyPath   = "ubl21-dian/soap/templates/operations/send_event_update_status_body.tmpl"
	getStatusBodyPath               = "ubl21-dian/soap/templates/operations/get_status_body.tmpl"
	getStatusZipBodyPath            = "ubl21-dian/soap/templates/operations/get_status_zip_body.tmpl"
	getStatusEventBodyPath          = "ubl21-dian/soap/templates/operations/get_status_event_body.tmpl"
	getXmlByDocumentKeyBodyPath     = "ubl21-dian/soap/templates/operations/get_xml_by_document_key_body.tmpl"
	getNumberingRangeBodyPath       = "ubl21-dian/soap/templates/operations/get_numbering_range_body.tmpl"
	getReferenceNotesBodyPath       = "ubl21-dian/soap/templates/operations/get_reference_notes_body.tmpl"
	getDocumentInfoBodyPath         = "ubl21-dian/soap/templates/operations/get_document_info_body.tmpl"
	getAcquirerBodyPath             = "ubl21-dian/soap/templates/operations/get_acquirer_body.tmpl"
	getExchangeEmailsBodyPath       = "ubl21-dian/soap/templates/operations/get_exchange_emails_body.tmpl"
)

// BuildSendBillSyncBody construye el body para SendBillSync
func BuildSendBillSyncBody(req *types.SendBillSyncRequest) string {
	tmpl, err := template.ParseFiles(sendBillSyncBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildSendBillAsyncBody construye el body para SendBillAsync
func BuildSendBillAsyncBody(req *types.SendBillAsyncRequest) string {
	tmpl, err := template.ParseFiles(sendBillAsyncBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildSendTestSetAsyncBody construye el body para SendTestSetAsync
func BuildSendTestSetAsyncBody(req *types.SendTestSetAsyncRequest) string {
	tmpl, err := template.ParseFiles(sendTestSetAsyncBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
		"TestSetId":   req.TestSetId,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildSendBillAttachmentAsyncBody construye el body para SendBillAttachmentAsync
func BuildSendBillAttachmentAsyncBody(req *types.SendBillAttachmentAsyncRequest) string {
	tmpl, err := template.ParseFiles(sendBillAttachmentAsyncBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildSendNominaSyncBody construye el body para SendNominaSync
func BuildSendNominaSyncBody(req *types.SendNominaSyncRequest) string {
	tmpl, err := template.ParseFiles(sendNominaSyncBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildSendEventBody construye el body para SendEvent
func BuildSendEventBody(req *types.SendEventRequest) string {
	tmpl, err := template.ParseFiles(sendEventUpdateStatusBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"FileName":    req.FileName,
		"ContentFile": req.ContentFile,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetStatusBody construye el body para GetStatus
func BuildGetStatusBody(req *types.GetStatusRequest) string {
	tmpl, err := template.ParseFiles(getStatusBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"TrackId": req.TrackId,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetStatusZipBody construye el body para GetStatusZip
func BuildGetStatusZipBody(req *types.GetStatusZipRequest) string {
	tmpl, err := template.ParseFiles(getStatusZipBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"TrackId": req.TrackId,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetStatusEventBody construye el body para GetStatusEvent
func BuildGetStatusEventBody(req *types.GetStatusEventRequest) string {
	tmpl, err := template.ParseFiles(getStatusEventBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"TrackId": req.TrackId,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetXmlByDocumentKeyBody construye el body para GetXmlByDocumentKey
func BuildGetXmlByDocumentKeyBody(req *types.GetXmlByDocumentKeyRequest) string {
	tmpl, err := template.ParseFiles(getXmlByDocumentKeyBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"TrackId": req.TrackId,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetNumberingRangeBody construye el body para GetNumberingRange
func BuildGetNumberingRangeBody(req *types.GetNumberingRangeRequest) string {
	tmpl, err := template.ParseFiles(getNumberingRangeBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"AccountCode":  req.NIT,
		"AccountCodeT": req.NIT,
		"SoftwareCode": req.SoftwareID,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetReferenceNotesBody construye el body para GetReferenceNotes
func BuildGetReferenceNotesBody(req *types.GetReferenceNotesRequest) string {
	tmpl, err := template.ParseFiles(getReferenceNotesBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"DocumentKey": req.DocumentKey,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetDocumentInfoBody construye el body para GetDocumentInfo
func BuildGetDocumentInfoBody(req *types.GetDocumentInfoRequest) string {
	tmpl, err := template.ParseFiles(getDocumentInfoBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"DocumentKey": req.DocumentKey,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetAcquirerBody construye el body para GetAcquirer
func BuildGetAcquirerBody(req *types.GetAcquirerRequest) string {
	tmpl, err := template.ParseFiles(getAcquirerBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"AccountCode":  req.NIT,
		"AccountCodeT": req.NIT,
	}); err != nil {
		return ""
	}

	return buffer.String()
}

// BuildGetExchangeEmailsBody construye el body para GetExchangeEmails
func BuildGetExchangeEmailsBody(req *types.GetExchangeEmailsRequest) string {
	tmpl, err := template.ParseFiles(getExchangeEmailsBodyPath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"AccountCode":  req.NIT,
		"AccountCodeT": req.NIT,
	}); err != nil {
		return ""
	}

	return buffer.String()
}
