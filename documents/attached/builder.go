package attached

import (
	"time"

	"github.com/diegofxm/ubl21-dian/documents/common/types"
	xmlpkg "github.com/diegofxm/ubl21-dian/xml"
)

// Builder construye un AttachedDocument UBL 2.1 directamente con structs XML
type Builder struct {
	doc *AttachedDocumentXML
}

// NewBuilder crea un nuevo builder de AttachedDocument
func NewBuilder() *Builder {
	return &Builder{
		doc: &AttachedDocumentXML{
			// Namespaces UBL 2.1 estándar para AttachedDocument
			Xmlns:         "urn:oasis:names:specification:ubl:schema:xsd:AttachedDocument-2",
			XmlnsCAC:      "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
			XmlnsCBC:      "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
			XmlnsCCTS:     "urn:un:unece:uncefact:data:specification:CoreComponentTypeSchemaModule:2",
			XmlnsExt:      "urn:oasis:names:specification:ubl:schema:xsd:CommonExtensionComponents-2",
			XmlnsXades:    "http://uri.etsi.org/01903/v1.3.2#",
			XmlnsXades141: "http://uri.etsi.org/01903/v1.4.1#",
			XmlnsDS:       "http://www.w3.org/2000/09/xmldsig#",

			// Valores por defecto DIAN
			UBLVersionID: types.CBCElement{Value: "UBL 2.1"},
			CustomizationID: types.IDElement{
				SchemeName: "31",
				Value:      "Documentos adjuntos",
			},
			ProfileID:    types.CBCElement{Value: "Factura Electrónica de Venta"},
			DocumentType: types.CBCElement{Value: "Contenedor de Factura Electrónica"},

			// Inicializar UBLExtensions vacío (se llenará con la firma)
			UBLExtensions: types.UBLExtensions{
				UBLExtension: []types.UBLExtension{
					{
						ExtensionContent: types.ExtensionContent{
							InnerXML: "",
						},
					},
				},
			},
		},
	}
}

// SetProfileExecutionID establece el ambiente
// "1" = Producción, "2" = Habilitación
func (b *Builder) SetProfileExecutionID(id string) *Builder {
	b.doc.ProfileExecutionID = types.CBCElement{Value: id}
	return b
}

// SetID establece el ID único del AttachedDocument (puede ser el CUFE)
func (b *Builder) SetID(id string) *Builder {
	b.doc.ID = types.CBCElement{Value: xmlpkg.Sanitize(id)}
	return b
}

// SetIssueDate establece la fecha y hora de emisión
func (b *Builder) SetIssueDate(date time.Time) *Builder {
	b.doc.IssueDate = types.CBCElement{Value: date.Format("2006-01-02")}
	b.doc.IssueTime = types.CBCElement{Value: date.Format("15:04:05-07:00")}
	return b
}

// SetParentDocumentID establece el número de la factura padre
func (b *Builder) SetParentDocumentID(id string) *Builder {
	b.doc.ParentDocumentID = types.CBCElement{Value: xmlpkg.Sanitize(id)}
	return b
}

// SetSender establece el emisor del AttachedDocument
func (b *Builder) SetSender(sender PartyData) *Builder {
	b.doc.SenderParty = SenderPartyXML{
		PartyTaxScheme: types.PartyTaxSchemeXML{
			RegistrationName: types.CBCElement{Value: xmlpkg.Sanitize(sender.RegistrationName)},
			CompanyID: types.IDElement{
				SchemeAgencyID:   "195",
				SchemeAgencyName: "CO, DIAN (Dirección de Impuestos y Aduanas Nacionales)",
				SchemeID:         sender.SchemeID,
				SchemeName:       sender.SchemeName,
				Value:            sender.CompanyID,
			},
			TaxLevelCode: types.TaxLevelCodeElement{ListName: "Régimen del contribuyente", Value: sender.TaxLevelCode},
			TaxScheme: types.TaxSchemeXML{
				ID:   types.CBCElement{Value: sender.TaxSchemeID},
				Name: types.CBCElement{Value: sender.TaxSchemeName},
			},
		},
	}
	return b
}

// SetReceiver establece el receptor del AttachedDocument (opcional)
func (b *Builder) SetReceiver(receiver PartyData) *Builder {
	b.doc.ReceiverParty = &ReceiverPartyXML{
		PartyTaxScheme: types.PartyTaxSchemeXML{
			RegistrationName: types.CBCElement{Value: xmlpkg.Sanitize(receiver.RegistrationName)},
			CompanyID: types.IDElement{
				SchemeAgencyID:   "195",
				SchemeAgencyName: "CO, DIAN (Dirección de Impuestos y Aduanas Nacionales)",
				SchemeID:         receiver.SchemeID,
				SchemeName:       receiver.SchemeName,
				Value:            receiver.CompanyID,
			},
			TaxLevelCode: types.TaxLevelCodeElement{ListName: "Régimen del contribuyente", Value: receiver.TaxLevelCode},
			TaxScheme: types.TaxSchemeXML{
				ID:   types.CBCElement{Value: receiver.TaxSchemeID},
				Name: types.CBCElement{Value: receiver.TaxSchemeName},
			},
		},
	}
	return b
}

// SetSignedInvoiceXML establece el XML de la factura firmada completa en CDATA
func (b *Builder) SetSignedInvoiceXML(signedXML string) *Builder {
	b.doc.Attachment = AttachmentXML{
		ExternalReference: ExternalReferenceXML{
			MimeCode:     types.CBCElement{Value: "text/xml"},
			EncodingCode: types.CBCElement{Value: "UTF-8"},
			Description: types.CDATAElement{
				Value: signedXML,
			},
		},
	}
	return b
}

// SetApplicationResponse establece el ApplicationResponse de DIAN en CDATA
func (b *Builder) SetApplicationResponse(appResponse ApplicationResponseData) *Builder {
	b.doc.ParentDocumentLineReference = ParentDocumentLineReferenceXML{
		LineID: types.CBCElement{Value: "1"},
		DocumentReference: DocumentReferenceXML{
			ID: types.CBCElement{Value: xmlpkg.Sanitize(appResponse.InvoiceID)},
			UUID: types.UUIDElement{
				SchemeName: "CUFE-SHA384",
				Value:      appResponse.CUFE,
			},
			IssueDate:    types.CBCElement{Value: appResponse.IssueDate.Format("2006-01-02")},
			DocumentType: types.CBCElement{Value: "ApplicationResponse"},
			Attachment: AttachmentXML{
				ExternalReference: ExternalReferenceXML{
					MimeCode:     types.CBCElement{Value: "text/xml"},
					EncodingCode: types.CBCElement{Value: "UTF-8"},
					Description: types.CDATAElement{
						Value: appResponse.ResponseXML,
					},
				},
			},
			ResultOfVerification: ResultOfVerificationXML{
				ValidatorID:          types.CBCElement{Value: "Unidad Especial Dirección de Impuestos y Aduanas Nacionales"},
				ValidationResultCode: types.CBCElement{Value: appResponse.ValidationResultCode},
				ValidationDate:       types.CBCElement{Value: appResponse.ValidationDate},
				ValidationTime:       types.CBCElement{Value: appResponse.ValidationTime},
			},
		},
	}
	return b
}

// Build construye y retorna el AttachedDocument XML
func (b *Builder) Build() *AttachedDocumentXML {
	return b.doc
}

// ToXML genera el XML del AttachedDocument
func (b *Builder) ToXML() ([]byte, error) {
	return xmlpkg.Marshal(b.doc)
}
