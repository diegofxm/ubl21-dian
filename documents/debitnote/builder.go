package debitnote

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/diegofxm/ubl21-dian/documents/common"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// Builder constructor de notas débito
type Builder struct {
	data DebitNoteTemplateData
}

// NewBuilder crea un nuevo builder
func NewBuilder() *Builder {
	return &Builder{
		data: DebitNoteTemplateData{
			CurrencyCode:      "COP",
			DebitNoteTypeCode: "92",
			Environment:       "2",
		},
	}
}

// SetDebitNoteData establece datos básicos
func (b *Builder) SetDebitNoteData(number, cude, issueDate, issueTime string) *Builder {
	b.data.DebitNoteNumber = number
	b.data.CUDE = cude
	b.data.IssueDate = issueDate
	b.data.IssueTime = issueTime
	return b
}

// SetProfileExecutionID establece el ambiente
func (b *Builder) SetProfileExecutionID(id string) *Builder {
	b.data.ProfileExecutionID = id
	b.data.Environment = id
	return b
}

// SetNote establece nota descriptiva
func (b *Builder) SetNote(note string) *Builder {
	b.data.Note = note
	return b
}

// SetDianExtensions establece extensiones DIAN
func (b *Builder) SetDianExtensions(invoiceAuth, authStartDate, authEndDate, prefix, from, to,
	providerID, providerSchemeID, providerSchemeName, softwareID, securityCode, qrCode string) *Builder {
	b.data.InvoiceAuthorization = invoiceAuth
	b.data.AuthPeriodStartDate = authStartDate
	b.data.AuthPeriodEndDate = authEndDate
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

// SetBillingReference establece referencia a factura
func (b *Builder) SetBillingReference(invoiceID, uuid, issueDate string) *Builder {
	b.data.BillingReference = &BillingReferenceTemplateData{
		ID:        invoiceID,
		UUID:      uuid,
		IssueDate: issueDate,
	}
	return b
}

// SetSupplier establece emisor
func (b *Builder) SetSupplier(supplier PartyTemplateData) *Builder {
	b.data.Supplier = supplier
	return b
}

// SetCustomer establece receptor
func (b *Builder) SetCustomer(customer PartyTemplateData) *Builder {
	b.data.Customer = customer
	return b
}

// SetTotals establece totales
func (b *Builder) SetTotals(lineExtension, taxExclusive, taxInclusive, payable string) *Builder {
	b.data.LineExtensionAmount = lineExtension
	b.data.TaxExclusiveAmount = taxExclusive
	b.data.TaxInclusiveAmount = taxInclusive
	b.data.PayableAmount = payable
	return b
}

// AddLine agrega línea
func (b *Builder) AddLine(line DebitNoteLineTemplateData) *Builder {
	b.data.DebitNoteLines = append(b.data.DebitNoteLines, line)
	b.data.LineCount = len(b.data.DebitNoteLines)
	return b
}

// AddTaxTotal agrega total de impuestos
func (b *Builder) AddTaxTotal(taxTotal TaxTotalTemplateData) *Builder {
	b.data.TaxTotals = append(b.data.TaxTotals, taxTotal)
	return b
}

// Build genera el XML
func (b *Builder) Build() ([]byte, error) {
	// Cargar templates comunes + específicos usando helper
	tmpl, err := common.LoadCommonAndSpecificTemplates(templatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "debitnote.tmpl", b.data); err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buf.Bytes(), nil
}
