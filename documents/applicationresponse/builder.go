package applicationresponse

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/diegofxm/ubl21-dian/documents/common"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// Builder constructor para ApplicationResponse
type Builder struct {
	doc ApplicationResponseXML
}

// NewBuilder crea un nuevo builder de ApplicationResponse
func NewBuilder() *Builder {
	return &Builder{
		doc: ApplicationResponseXML{
			// Namespaces UBL 2.1
			Xmlns:             "urn:oasis:names:specification:ubl:schema:xsd:ApplicationResponse-2",
			XmlnsCAC:          "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
			XmlnsCBC:          "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
			XmlnsCCTS:         "urn:un:unece:uncefact:documentation:2",
			XmlnsDS:           "http://www.w3.org/2000/09/xmldsig#",
			XmlnsExt:          "urn:oasis:names:specification:ubl:schema:xsd:CommonExtensionComponents-2",
			XmlnsSts:          "dian:gov:co:facturaelectronica:Structures-2-1",
			XmlnsXades:        "http://uri.etsi.org/01903/v1.3.2#",
			XmlnsXades141:     "http://uri.etsi.org/01903/v1.4.1#",
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "urn:oasis:names:specification:ubl:schema:xsd:ApplicationResponse-2 http://docs.oasis-open.org/ubl/os-UBL-2.1/xsd/maindoc/UBL-ApplicationResponse-2.1.xsd",

			// Valores por defecto
			UBLVersionID:       types.CBCElement{Value: "UBL 2.1"},
			CustomizationID:    types.CBCElement{Value: "1"},
			ProfileID:          types.CBCElement{Value: "DIAN 2.1"},
			ProfileExecutionID: types.CBCElement{Value: "1"}, // 1=Producción, 2=Habilitación
		},
	}
}

// SetID establece el ID del ApplicationResponse
func (b *Builder) SetID(id string) *Builder {
	b.doc.ID = types.CBCElement{Value: id}
	return b
}

// SetCUDE establece el CUDE (Código Único de Documento Electrónico)
func (b *Builder) SetCUDE(cude string) *Builder {
	b.doc.UUID = types.UUIDElement{
		SchemeName: "CUDE-SHA384",
		Value:      cude,
	}
	return b
}

// SetIssueDate establece la fecha de emisión (YYYY-MM-DD)
func (b *Builder) SetIssueDate(date string) *Builder {
	b.doc.IssueDate = types.CBCElement{Value: date}
	return b
}

// SetIssueTime establece la hora de emisión (HH:MM:SS-05:00)
func (b *Builder) SetIssueTime(time string) *Builder {
	b.doc.IssueTime = types.CBCElement{Value: time}
	return b
}

// SetProfileExecutionID establece el ambiente (1=Producción, 2=Habilitación)
func (b *Builder) SetProfileExecutionID(id string) *Builder {
	b.doc.ProfileExecutionID = types.CBCElement{Value: id}
	return b
}

// AddNote agrega una nota al ApplicationResponse
func (b *Builder) AddNote(note string) *Builder {
	b.doc.Note = append(b.doc.Note, types.CBCElement{Value: note})
	return b
}

// SetSenderParty establece el emisor (DIAN)
func (b *Builder) SetSenderParty(party PartyData) *Builder {
	b.doc.SenderParty = SenderPartyXML{
		PartyTaxScheme: PartyTaxSchemeXML{
			RegistrationName: types.CBCElement{Value: party.RegistrationName},
			CompanyID: types.IDElement{
				SchemeID:     party.SchemeID,
				SchemeName:   party.SchemeName,
				SchemeAgencyID: "195",
				SchemeAgencyName: "CO, DIAN (Dirección de Impuestos y Aduanas Nacionales)",
				Value:        party.CompanyID,
			},
			TaxLevelCode: types.TaxLevelCodeElement{
				ListName: "Fiscal Responsibility",
				Value:    party.TaxLevelCode,
			},
			TaxScheme: TaxSchemeXML{
				ID:   types.CBCElement{Value: party.TaxSchemeID},
				Name: types.CBCElement{Value: party.TaxSchemeName},
			},
		},
	}
	return b
}

// SetReceiverParty establece el receptor (empresa)
func (b *Builder) SetReceiverParty(party PartyData) *Builder {
	b.doc.ReceiverParty = ReceiverPartyXML{
		PartyTaxScheme: PartyTaxSchemeXML{
			RegistrationName: types.CBCElement{Value: party.RegistrationName},
			CompanyID: types.IDElement{
				SchemeID:     party.SchemeID,
				SchemeName:   party.SchemeName,
				SchemeAgencyID: "195",
				SchemeAgencyName: "CO, DIAN (Dirección de Impuestos y Aduanas Nacionales)",
				Value:        party.CompanyID,
			},
			TaxLevelCode: types.TaxLevelCodeElement{
				ListName: "Fiscal Responsibility",
				Value:    party.TaxLevelCode,
			},
			TaxScheme: TaxSchemeXML{
				ID:   types.CBCElement{Value: party.TaxSchemeID},
				Name: types.CBCElement{Value: party.TaxSchemeName},
			},
		},
	}
	return b
}

// SetResponse establece la respuesta de validación
func (b *Builder) SetResponse(code string, descriptions ...string) *Builder {
	response := ResponseXML{
		ResponseCode: types.CBCElement{Value: code},
	}
	
	for _, desc := range descriptions {
		response.Description = append(response.Description, types.CBCElement{Value: desc})
	}
	
	b.doc.DocumentResponse.Response = response
	return b
}

// SetDocumentReference establece la referencia al documento validado
func (b *Builder) SetDocumentReference(ref DocumentReferenceData) *Builder {
	docRef := DocumentReferenceXML{
		ID: types.CBCElement{Value: ref.ID},
		UUID: types.UUIDElement{
			SchemeName: "CUFE-SHA384",
			Value:      ref.UUID,
		},
		IssueDate: types.CBCElement{Value: ref.IssueDate.Format("2006-01-02")},
	}
	
	if ref.DocumentTypeCode != "" {
		docRef.DocumentTypeCode = &types.CBCElement{Value: ref.DocumentTypeCode}
	}
	
	if ref.DocumentType != "" {
		docRef.DocumentType = &types.CBCElement{Value: ref.DocumentType}
	}
	
	if ref.ValidationResult != nil {
		validation := &ResultOfVerificationXML{
			ValidatorID:          types.CBCElement{Value: ref.ValidationResult.ValidatorID},
			ValidationResultCode: types.CBCElement{Value: ref.ValidationResult.ValidationResultCode},
			ValidationDate:       types.CBCElement{Value: ref.ValidationResult.ValidationDate.Format("2006-01-02")},
			ValidationTime:       types.CBCElement{Value: ref.ValidationResult.ValidationTime},
		}
		
		if ref.ValidationResult.ValidateProcess != "" {
			validation.ValidateProcess = &types.CBCElement{Value: ref.ValidationResult.ValidateProcess}
		}
		if ref.ValidationResult.ValidateTool != "" {
			validation.ValidateTool = &types.CBCElement{Value: ref.ValidationResult.ValidateTool}
		}
		if ref.ValidationResult.ValidateToolVersion != "" {
			validation.ValidateToolVersion = &types.CBCElement{Value: ref.ValidationResult.ValidateToolVersion}
		}
		
		docRef.ResultOfVerification = validation
	}
	
	b.doc.DocumentResponse.DocumentReference = docRef
	return b
}

// Build genera el XML del ApplicationResponse usando templates
func (b *Builder) Build() (string, error) {
	// Cargar templates comunes + específicos
	tmpl, err := common.LoadCommonAndSpecificTemplates(templatesFS, "templates/*.tmpl")
	if err != nil {
		return "", fmt.Errorf("error loading templates: %w", err)
	}
	
	// Preparar datos para template
	data := b.prepareTemplateData()
	
	// Ejecutar template
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "applicationresponse.tmpl", data); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}
	
	return buf.String(), nil
}

// prepareTemplateData prepara los datos para el template
func (b *Builder) prepareTemplateData() ApplicationResponseTemplateData {
	data := ApplicationResponseTemplateData{
		ProfileExecutionID: b.doc.ProfileExecutionID.Value,
		ID:                 b.doc.ID.Value,
		CUDE:               b.doc.UUID.Value,
		Environment:        b.doc.ProfileExecutionID.Value,
		IssueDate:          b.doc.IssueDate.Value,
		IssueTime:          b.doc.IssueTime.Value,
		ResponseCode:       b.doc.DocumentResponse.Response.ResponseCode.Value,
	}
	
	// Notes
	for _, note := range b.doc.Note {
		data.Notes = append(data.Notes, note.Value)
	}
	
	// Descriptions
	for _, desc := range b.doc.DocumentResponse.Response.Description {
		data.Descriptions = append(data.Descriptions, desc.Value)
	}
	
	// Sender Party
	data.SenderParty = PartyTemplateData{
		RegistrationName: b.doc.SenderParty.PartyTaxScheme.RegistrationName.Value,
		CompanyID:        b.doc.SenderParty.PartyTaxScheme.CompanyID.Value,
		SchemeID:         b.doc.SenderParty.PartyTaxScheme.CompanyID.SchemeID,
		SchemeName:       b.doc.SenderParty.PartyTaxScheme.CompanyID.SchemeName,
		TaxLevelCode:     b.doc.SenderParty.PartyTaxScheme.TaxLevelCode.Value,
		TaxSchemeID:      b.doc.SenderParty.PartyTaxScheme.TaxScheme.ID.Value,
		TaxSchemeName:    b.doc.SenderParty.PartyTaxScheme.TaxScheme.Name.Value,
	}
	
	// Receiver Party
	data.ReceiverParty = PartyTemplateData{
		RegistrationName: b.doc.ReceiverParty.PartyTaxScheme.RegistrationName.Value,
		CompanyID:        b.doc.ReceiverParty.PartyTaxScheme.CompanyID.Value,
		SchemeID:         b.doc.ReceiverParty.PartyTaxScheme.CompanyID.SchemeID,
		SchemeName:       b.doc.ReceiverParty.PartyTaxScheme.CompanyID.SchemeName,
		TaxLevelCode:     b.doc.ReceiverParty.PartyTaxScheme.TaxLevelCode.Value,
		TaxSchemeID:      b.doc.ReceiverParty.PartyTaxScheme.TaxScheme.ID.Value,
		TaxSchemeName:    b.doc.ReceiverParty.PartyTaxScheme.TaxScheme.Name.Value,
	}
	
	// Document Reference
	data.DocumentReference = DocumentReferenceTemplateData{
		ID:               b.doc.DocumentResponse.DocumentReference.ID.Value,
		UUID:             b.doc.DocumentResponse.DocumentReference.UUID.Value,
		IssueDate:        b.doc.DocumentResponse.DocumentReference.IssueDate.Value,
	}
	
	if b.doc.DocumentResponse.DocumentReference.DocumentTypeCode != nil {
		data.DocumentReference.DocumentTypeCode = b.doc.DocumentResponse.DocumentReference.DocumentTypeCode.Value
	}
	
	if b.doc.DocumentResponse.DocumentReference.DocumentType != nil {
		data.DocumentReference.DocumentType = b.doc.DocumentResponse.DocumentReference.DocumentType.Value
	}
	
	if b.doc.DocumentResponse.DocumentReference.ResultOfVerification != nil {
		validation := b.doc.DocumentResponse.DocumentReference.ResultOfVerification
		data.DocumentReference.ValidationResult = &ValidationResultTemplateData{
			ValidatorID:          validation.ValidatorID.Value,
			ValidationResultCode: validation.ValidationResultCode.Value,
			ValidationDate:       validation.ValidationDate.Value,
			ValidationTime:       validation.ValidationTime.Value,
		}
		
		if validation.ValidateProcess != nil {
			data.DocumentReference.ValidationResult.ValidateProcess = validation.ValidateProcess.Value
		}
		if validation.ValidateTool != nil {
			data.DocumentReference.ValidationResult.ValidateTool = validation.ValidateTool.Value
		}
		if validation.ValidateToolVersion != nil {
			data.DocumentReference.ValidationResult.ValidateToolVersion = validation.ValidateToolVersion.Value
		}
	}
	
	return data
}

// GetModel retorna el modelo XML completo
func (b *Builder) GetModel() ApplicationResponseXML {
	return b.doc
}
