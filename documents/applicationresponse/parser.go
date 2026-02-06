package applicationresponse

import (
	"encoding/xml"
	"fmt"
	"time"
)

// ParseFromXML parsea un XML de ApplicationResponse de DIAN
func ParseFromXML(xmlData []byte) (*ApplicationResponseData, error) {
	var appResp ApplicationResponseXML
	
	if err := xml.Unmarshal(xmlData, &appResp); err != nil {
		return nil, fmt.Errorf("error parsing ApplicationResponse XML: %w", err)
	}
	
	data := &ApplicationResponseData{
		ID:                 appResp.ID.Value,
		CUDE:               appResp.UUID.Value,
		ProfileExecutionID: appResp.ProfileExecutionID.Value,
	}
	
	// Parse dates
	if issueDate, err := time.Parse("2006-01-02", appResp.IssueDate.Value); err == nil {
		data.IssueDate = issueDate
	}
	data.IssueTime = appResp.IssueTime.Value
	
	// Parse notes
	for _, note := range appResp.Note {
		data.Notes = append(data.Notes, note.Value)
	}
	
	// Parse sender party (DIAN)
	data.SenderParty = PartyData{
		RegistrationName: appResp.SenderParty.PartyTaxScheme.RegistrationName.Value,
		CompanyID:        appResp.SenderParty.PartyTaxScheme.CompanyID.Value,
		SchemeID:         appResp.SenderParty.PartyTaxScheme.CompanyID.SchemeID,
		SchemeName:       appResp.SenderParty.PartyTaxScheme.CompanyID.SchemeName,
		TaxLevelCode:     appResp.SenderParty.PartyTaxScheme.TaxLevelCode.Value,
		TaxSchemeID:      appResp.SenderParty.PartyTaxScheme.TaxScheme.ID.Value,
		TaxSchemeName:    appResp.SenderParty.PartyTaxScheme.TaxScheme.Name.Value,
	}
	
	// Parse receiver party (empresa)
	data.ReceiverParty = PartyData{
		RegistrationName: appResp.ReceiverParty.PartyTaxScheme.RegistrationName.Value,
		CompanyID:        appResp.ReceiverParty.PartyTaxScheme.CompanyID.Value,
		SchemeID:         appResp.ReceiverParty.PartyTaxScheme.CompanyID.SchemeID,
		SchemeName:       appResp.ReceiverParty.PartyTaxScheme.CompanyID.SchemeName,
		TaxLevelCode:     appResp.ReceiverParty.PartyTaxScheme.TaxLevelCode.Value,
		TaxSchemeID:      appResp.ReceiverParty.PartyTaxScheme.TaxScheme.ID.Value,
		TaxSchemeName:    appResp.ReceiverParty.PartyTaxScheme.TaxScheme.Name.Value,
	}
	
	// Parse response
	data.ResponseCode = appResp.DocumentResponse.Response.ResponseCode.Value
	for _, desc := range appResp.DocumentResponse.Response.Description {
		data.Descriptions = append(data.Descriptions, desc.Value)
	}
	
	// Parse document reference
	docRef := appResp.DocumentResponse.DocumentReference
	data.DocumentReference = DocumentReferenceData{
		ID:   docRef.ID.Value,
		UUID: docRef.UUID.Value,
	}
	
	if issueDate, err := time.Parse("2006-01-02", docRef.IssueDate.Value); err == nil {
		data.DocumentReference.IssueDate = issueDate
	}
	
	if docRef.DocumentTypeCode != nil {
		data.DocumentReference.DocumentTypeCode = docRef.DocumentTypeCode.Value
	}
	
	if docRef.DocumentType != nil {
		data.DocumentReference.DocumentType = docRef.DocumentType.Value
	}
	
	// Parse validation result
	if docRef.ResultOfVerification != nil {
		validation := docRef.ResultOfVerification
		validationData := &ValidationResultData{
			ValidatorID:          validation.ValidatorID.Value,
			ValidationResultCode: validation.ValidationResultCode.Value,
			ValidationTime:       validation.ValidationTime.Value,
		}
		
		if validationDate, err := time.Parse("2006-01-02", validation.ValidationDate.Value); err == nil {
			validationData.ValidationDate = validationDate
		}
		
		if validation.ValidateProcess != nil {
			validationData.ValidateProcess = validation.ValidateProcess.Value
		}
		if validation.ValidateTool != nil {
			validationData.ValidateTool = validation.ValidateTool.Value
		}
		if validation.ValidateToolVersion != nil {
			validationData.ValidateToolVersion = validation.ValidateToolVersion.Value
		}
		
		data.DocumentReference.ValidationResult = validationData
	}
	
	return data, nil
}

// ParseFromString parsea un string XML de ApplicationResponse
func ParseFromString(xmlString string) (*ApplicationResponseData, error) {
	return ParseFromXML([]byte(xmlString))
}

// IsValidated verifica si el documento fue validado exitosamente por DIAN
func (a *ApplicationResponseData) IsValidated() bool {
	return a.ResponseCode == "02"
}

// GetValidationDateTime retorna la fecha y hora de validación combinadas
func (a *ApplicationResponseData) GetValidationDateTime() (time.Time, error) {
	if a.DocumentReference.ValidationResult == nil {
		return time.Time{}, fmt.Errorf("no validation result available")
	}
	
	dateStr := a.DocumentReference.ValidationResult.ValidationDate.Format("2006-01-02")
	timeStr := a.DocumentReference.ValidationResult.ValidationTime
	
	// Combinar fecha y hora
	dateTimeStr := dateStr + " " + timeStr
	
	// Parse con timezone
	layout := "2006-01-02 15:04:05-07:00"
	return time.Parse(layout, dateTimeStr)
}
