package supportdocument

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/diegofxm/ubl21-dian/documents/common"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// Builder constructor para SupportDocument
type Builder struct {
	data SupportDocumentTemplateData
}

// NewBuilder crea un nuevo builder de SupportDocument
func NewBuilder() *Builder {
	return &Builder{
		data: SupportDocumentTemplateData{
			CurrencyCode:     "COP",
			DocumentTypeCode: "05",
			ProfileExecutionID: "2", // Habilitación por defecto
		},
	}
}

// SetSupportDocumentData establece datos básicos
func (b *Builder) SetSupportDocumentData(number, cuds, issueDate, issueTime string) *Builder {
	b.data.SupportDocNumber = number
	b.data.CUDS = cuds
	b.data.IssueDate = issueDate
	b.data.IssueTime = issueTime
	return b
}

// SetID establece el número del documento soporte
func (b *Builder) SetID(id string) *Builder {
	b.data.SupportDocNumber = id
	return b
}

// SetCUDS establece el CUDS
func (b *Builder) SetCUDS(cuds string) *Builder {
	b.data.CUDS = cuds
	return b
}

// SetIssueDate establece la fecha de emisión
func (b *Builder) SetIssueDate(date string) *Builder {
	b.data.IssueDate = date
	return b
}

// SetIssueTime establece la hora de emisión
func (b *Builder) SetIssueTime(time string) *Builder {
	b.data.IssueTime = time
	return b
}

// SetProfileExecutionID establece el ambiente
func (b *Builder) SetProfileExecutionID(id string) *Builder {
	b.data.ProfileExecutionID = id
	b.data.Environment = id
	return b
}

// AddNote agrega una nota
func (b *Builder) AddNote(note string) *Builder {
	b.data.Notes = append(b.data.Notes, note)
	return b
}

// SetDianExtensions establece extensiones DIAN
func (b *Builder) SetDianExtensions(auth, startDate, endDate, prefix, from, to, providerID, providerSchemeID, providerSchemeName, softwareID, securityCode, qrCode string) *Builder {
	b.data.InvoiceAuthorization = auth
	b.data.AuthPeriodStartDate = startDate
	b.data.AuthPeriodEndDate = endDate
	b.data.Prefix = prefix
	b.data.From = from
	b.data.To = to
	b.data.ProviderID = providerID
	b.data.ProviderSchemeID = providerSchemeID
	b.data.ProviderSchemeName = providerSchemeName
	b.data.SoftwareID = softwareID
	b.data.SecurityCode = securityCode
	b.data.QRCode = qrCode
	return b
}

// AddBillingReference agrega referencia a factura de proveedor
func (b *Builder) AddBillingReference(invoiceID, uuid, issueDate string) *Builder {
	b.data.BillingReferences = append(b.data.BillingReferences, BillingReferenceTemplateData{
		InvoiceID: invoiceID,
		UUID:      uuid,
		IssueDate: issueDate,
	})
	return b
}

// SetBuyer establece el comprador
func (b *Builder) SetBuyer(buyer PartyTemplateData) *Builder {
	b.data.Buyer = buyer
	return b
}

// SetSupplier establece el proveedor
func (b *Builder) SetSupplier(supplier PartyTemplateData) *Builder {
	b.data.Supplier = supplier
	return b
}

// AddLine agrega una línea
func (b *Builder) AddLine(line SupportDocumentLineTemplateData) *Builder {
	b.data.SupportDocumentLines = append(b.data.SupportDocumentLines, line)
	return b
}

// SetTotals establece los totales
func (b *Builder) SetTotals(lineExt, taxExc, taxInc, payable string) *Builder {
	b.data.LineExtensionAmount = lineExt
	b.data.TaxExclusiveAmount = taxExc
	b.data.TaxInclusiveAmount = taxInc
	b.data.PayableAmount = payable
	return b
}

// AddTaxTotal agrega total de impuestos
func (b *Builder) AddTaxTotal(taxTotal TaxTotalTemplateData) *Builder {
	b.data.TaxTotals = append(b.data.TaxTotals, taxTotal)
	return b
}

// AddWithholdingTaxTotal agrega total de retenciones
func (b *Builder) AddWithholdingTaxTotal(taxTotal TaxTotalTemplateData) *Builder {
	b.data.WithholdingTaxTotals = append(b.data.WithholdingTaxTotals, taxTotal)
	return b
}

// Build genera el XML del SupportDocument
func (b *Builder) Build() (string, error) {
	// Actualizar line count
	b.data.LineCount = len(b.data.SupportDocumentLines)
	
	// Cargar templates
	tmpl, err := common.LoadCommonAndSpecificTemplates(templatesFS, "templates/*.tmpl")
	if err != nil {
		return "", fmt.Errorf("error loading templates: %w", err)
	}
	
	// Ejecutar template
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "supportdocument.tmpl", b.data); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}
	
	return buf.String(), nil
}
