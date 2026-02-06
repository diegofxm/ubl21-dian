package applicationresponse

import (
	"strings"
	"testing"
	"time"
)

func TestApplicationResponseBuilder(t *testing.T) {
	t.Run("Build ApplicationResponse from DIAN", func(t *testing.T) {
		// Datos de prueba
		dianParty := PartyData{
			RegistrationName: "Unidad Especial Dirección de Impuestos y Aduanas Nacionales",
			CompanyID:        "800197268",
			SchemeID:         "4",
			SchemeName:       "31",
			TaxLevelCode:     "O-13",
			TaxSchemeID:      "01",
			TaxSchemeName:    "IVA",
		}
		
		companyParty := PartyData{
			RegistrationName: "MI EMPRESA SAS",
			CompanyID:        "900123456",
			SchemeID:         "1",
			SchemeName:       "31",
			TaxLevelCode:     "O-13",
			TaxSchemeID:      "01",
			TaxSchemeName:    "IVA",
		}
		
		docRef := DocumentReferenceData{
			ID:               "SETP990000001",
			UUID:             "ffff0d032c292b88b3f839f75a51e8459ab645eda8049b3c221649fd18aaea09d5b31c8787e071c6a7d4db6983faaead",
			IssueDate:        time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
			DocumentTypeCode: "01",
			DocumentType:     "Invoice",
			ValidationResult: &ValidationResultData{
				ValidatorID:          "Unidad Especial Dirección de Impuestos y Aduanas Nacionales",
				ValidationResultCode: "02",
				ValidationDate:       time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
				ValidationTime:       "14:30:00-05:00",
			},
		}
		
		// Construir ApplicationResponse
		builder := NewBuilder()
		xml, err := builder.
			SetID("18760000001").
			SetCUDE("abc123def456").
			SetIssueDate("2025-01-31").
			SetIssueTime("14:30:00-05:00").
			SetProfileExecutionID("1").
			AddNote("Documento validado correctamente").
			SetSenderParty(dianParty).
			SetReceiverParty(companyParty).
			SetResponse("02", "Documento validado por la DIAN").
			SetDocumentReference(docRef).
			Build()
		
		if err != nil {
			t.Fatalf("Error building ApplicationResponse: %v", err)
		}
		
		// Verificaciones
		if !strings.Contains(xml, "<ApplicationResponse") {
			t.Error("XML should contain ApplicationResponse root element")
		}
		
		if !strings.Contains(xml, "18760000001") {
			t.Error("XML should contain ID")
		}
		
		if !strings.Contains(xml, "abc123def456") {
			t.Error("XML should contain CUDE")
		}
		
		if !strings.Contains(xml, "<cbc:ResponseCode>02</cbc:ResponseCode>") {
			t.Error("XML should contain ResponseCode 02")
		}
		
		if !strings.Contains(xml, "Documento validado por la DIAN") {
			t.Error("XML should contain validation description")
		}
		
		if !strings.Contains(xml, "SETP990000001") {
			t.Error("XML should contain document reference ID")
		}
		
		if !strings.Contains(xml, "ffff0d032c292b88") {
			t.Error("XML should contain document CUFE")
		}
		
		t.Log("✓ ApplicationResponse XML generated successfully")
	})
	
	t.Run("Parse ApplicationResponse XML", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<ApplicationResponse xmlns="urn:oasis:names:specification:ubl:schema:xsd:ApplicationResponse-2" 
  xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2" 
  xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2">
  <cbc:UBLVersionID>UBL 2.1</cbc:UBLVersionID>
  <cbc:ID>18760000001</cbc:ID>
  <cbc:UUID schemeName="CUDE-SHA384">abc123def456</cbc:UUID>
  <cbc:IssueDate>2025-01-31</cbc:IssueDate>
  <cbc:IssueTime>14:30:00-05:00</cbc:IssueTime>
  <cbc:ProfileExecutionID>1</cbc:ProfileExecutionID>
  <cac:SenderParty>
    <cac:PartyTaxScheme>
      <cbc:RegistrationName>DIAN</cbc:RegistrationName>
      <cbc:CompanyID schemeID="4" schemeName="31">800197268</cbc:CompanyID>
      <cbc:TaxLevelCode>O-13</cbc:TaxLevelCode>
      <cac:TaxScheme>
        <cbc:ID>01</cbc:ID>
        <cbc:Name>IVA</cbc:Name>
      </cac:TaxScheme>
    </cac:PartyTaxScheme>
  </cac:SenderParty>
  <cac:ReceiverParty>
    <cac:PartyTaxScheme>
      <cbc:RegistrationName>MI EMPRESA SAS</cbc:RegistrationName>
      <cbc:CompanyID schemeID="1" schemeName="31">900123456</cbc:CompanyID>
      <cbc:TaxLevelCode>O-13</cbc:TaxLevelCode>
      <cac:TaxScheme>
        <cbc:ID>01</cbc:ID>
        <cbc:Name>IVA</cbc:Name>
      </cac:TaxScheme>
    </cac:PartyTaxScheme>
  </cac:ReceiverParty>
  <cac:DocumentResponse>
    <cac:Response>
      <cbc:ResponseCode>02</cbc:ResponseCode>
      <cbc:Description>Documento validado por la DIAN</cbc:Description>
    </cac:Response>
    <cac:DocumentReference>
      <cbc:ID>SETP990000001</cbc:ID>
      <cbc:UUID schemeName="CUFE-SHA384">ffff0d032c292b88</cbc:UUID>
      <cbc:IssueDate>2025-01-31</cbc:IssueDate>
      <cac:ResultOfVerification>
        <cbc:ValidatorID>DIAN</cbc:ValidatorID>
        <cbc:ValidationResultCode>02</cbc:ValidationResultCode>
        <cbc:ValidationDate>2025-01-31</cbc:ValidationDate>
        <cbc:ValidationTime>14:30:00-05:00</cbc:ValidationTime>
      </cac:ResultOfVerification>
    </cac:DocumentReference>
  </cac:DocumentResponse>
</ApplicationResponse>`
		
		data, err := ParseFromString(xmlData)
		if err != nil {
			t.Fatalf("Error parsing ApplicationResponse: %v", err)
		}
		
		// Verificaciones
		if data.ID != "18760000001" {
			t.Errorf("Expected ID 18760000001, got %s", data.ID)
		}
		
		if data.CUDE != "abc123def456" {
			t.Errorf("Expected CUDE abc123def456, got %s", data.CUDE)
		}
		
		if data.ResponseCode != "02" {
			t.Errorf("Expected ResponseCode 02, got %s", data.ResponseCode)
		}
		
		if !data.IsValidated() {
			t.Error("Document should be validated")
		}
		
		if data.SenderParty.CompanyID != "800197268" {
			t.Errorf("Expected sender NIT 800197268, got %s", data.SenderParty.CompanyID)
		}
		
		if data.ReceiverParty.CompanyID != "900123456" {
			t.Errorf("Expected receiver NIT 900123456, got %s", data.ReceiverParty.CompanyID)
		}
		
		if data.DocumentReference.ID != "SETP990000001" {
			t.Errorf("Expected document ID SETP990000001, got %s", data.DocumentReference.ID)
		}
		
		if data.DocumentReference.ValidationResult == nil {
			t.Fatal("Expected validation result")
		}
		
		if data.DocumentReference.ValidationResult.ValidationResultCode != "02" {
			t.Errorf("Expected validation code 02, got %s", data.DocumentReference.ValidationResult.ValidationResultCode)
		}
		
		t.Log("✓ ApplicationResponse parsed successfully")
		t.Logf("  - ID: %s", data.ID)
		t.Logf("  - CUDE: %s", data.CUDE)
		t.Logf("  - ResponseCode: %s", data.ResponseCode)
		t.Logf("  - Validated: %v", data.IsValidated())
	})
	
	t.Run("Validation helpers", func(t *testing.T) {
		data := &ApplicationResponseData{
			ResponseCode: "02",
			DocumentReference: DocumentReferenceData{
				ValidationResult: &ValidationResultData{
					ValidationDate: time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC),
					ValidationTime: "14:30:00-05:00",
				},
			},
		}
		
		if !data.IsValidated() {
			t.Error("Should be validated with code 02")
		}
		
		dateTime, err := data.GetValidationDateTime()
		if err != nil {
			t.Errorf("Error getting validation datetime: %v", err)
		}
		
		expectedYear := 2025
		if dateTime.Year() != expectedYear {
			t.Errorf("Expected year %d, got %d", expectedYear, dateTime.Year())
		}
		
		t.Log("✓ Validation helpers work correctly")
	})
}
