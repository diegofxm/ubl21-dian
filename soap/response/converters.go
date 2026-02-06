package response

import (
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// ToSendBillSyncResponse convierte ResponseXML a SendBillSyncResponse
func ToSendBillSyncResponse(xmlResp *ResponseXML) *types.SendBillSyncResponse {
	resp := &types.SendBillSyncResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToSendBillAsyncResponse convierte ResponseXML a SendBillAsyncResponse
func ToSendBillAsyncResponse(xmlResp *ResponseXML) *types.SendBillAsyncResponse {
	resp := &types.SendBillAsyncResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToSendTestSetAsyncResponse convierte ResponseXML a SendTestSetAsyncResponse
func ToSendTestSetAsyncResponse(xmlResp *ResponseXML) *types.SendTestSetAsyncResponse {
	resp := &types.SendTestSetAsyncResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToGetStatusResponse convierte ResponseXML a GetStatusResponse
func ToGetStatusResponse(xmlResp *ResponseXML) *types.GetStatusResponse {
	resp := &types.GetStatusResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToGetStatusZipResponse convierte GetStatusZipResponseXML a GetStatusZipResponse
func ToGetStatusZipResponse(xmlResp *GetStatusZipResponseXML) *types.GetStatusZipResponse {
	return &types.GetStatusZipResponse{
		ZipKey:        xmlResp.ZipKey,
		ContentFile:   xmlResp.ContentFile,
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
	}
}

// ToGetXmlByDocumentKeyResponse convierte GetXmlByDocumentKeyResponseXML a GetXmlByDocumentKeyResponse
func ToGetXmlByDocumentKeyResponse(xmlResp *GetXmlByDocumentKeyResponseXML) *types.GetXmlByDocumentKeyResponse {
	return &types.GetXmlByDocumentKeyResponse{
		XmlBase64Bytes: xmlResp.XmlBase64Bytes,
		StatusCode:     xmlResp.StatusCode,
		StatusMessage:  xmlResp.StatusMessage,
	}
}

// ToSendBillAttachmentAsyncResponse convierte ResponseXML a SendBillAttachmentAsyncResponse
func ToSendBillAttachmentAsyncResponse(xmlResp *ResponseXML) *types.SendBillAttachmentAsyncResponse {
	resp := &types.SendBillAttachmentAsyncResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToSendEventResponse convierte ResponseXML a SendEventResponse
func ToSendEventResponse(xmlResp *ResponseXML) *types.SendEventResponse {
	resp := &types.SendEventResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToSendNominaSyncResponse convierte ResponseXML a SendNominaSyncResponse
func ToSendNominaSyncResponse(xmlResp *ResponseXML) *types.SendNominaSyncResponse {
	resp := &types.SendNominaSyncResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToGetStatusEventResponse convierte ResponseXML a GetStatusEventResponse
func ToGetStatusEventResponse(xmlResp *ResponseXML) *types.GetStatusEventResponse {
	resp := &types.GetStatusEventResponse{
		Response: types.Response{
			IsValid:           xmlResp.IsValid == "true",
			StatusCode:        xmlResp.StatusCode,
			StatusDescription: xmlResp.StatusDescription,
			StatusMessage:     xmlResp.StatusMessage,
			XmlDocumentKey:    xmlResp.XmlDocumentKey,
			XmlBase64Bytes:    xmlResp.XmlBase64Bytes,
		},
	}

	for _, errMsg := range xmlResp.ErrorMessage {
		resp.ErrorMessage = append(resp.ErrorMessage, types.ErrorMessage{
			Code:        errMsg.Code,
			Description: errMsg.Description,
		})
	}

	return resp
}

// ToGetNumberingRangeResponse convierte ResponseXML a GetNumberingRangeResponse
func ToGetNumberingRangeResponse(xmlResp *ResponseXML) *types.GetNumberingRangeResponse {
	resp := &types.GetNumberingRangeResponse{
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
		Ranges:        []types.NumberingRange{},
	}
	return resp
}

// ToGetReferenceNotesResponse convierte ResponseXML a GetReferenceNotesResponse
func ToGetReferenceNotesResponse(xmlResp *ResponseXML) *types.GetReferenceNotesResponse {
	resp := &types.GetReferenceNotesResponse{
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
		Notes:         []types.ReferenceNote{},
	}
	return resp
}

// ToGetDocumentInfoResponse convierte ResponseXML a GetDocumentInfoResponse
func ToGetDocumentInfoResponse(xmlResp *ResponseXML) *types.GetDocumentInfoResponse {
	resp := &types.GetDocumentInfoResponse{
		DocumentKey:   xmlResp.XmlDocumentKey,
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
	}
	return resp
}

// ToGetAcquirerResponse convierte ResponseXML a GetAcquirerResponse
func ToGetAcquirerResponse(xmlResp *ResponseXML) *types.GetAcquirerResponse {
	resp := &types.GetAcquirerResponse{
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
	}
	return resp
}

// ToGetExchangeEmailsResponse convierte ResponseXML a GetExchangeEmailsResponse
func ToGetExchangeEmailsResponse(xmlResp *ResponseXML) *types.GetExchangeEmailsResponse {
	resp := &types.GetExchangeEmailsResponse{
		StatusCode:    xmlResp.StatusCode,
		StatusMessage: xmlResp.StatusMessage,
		Emails:        []string{},
	}
	return resp
}
