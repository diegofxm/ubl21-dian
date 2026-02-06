package invoice

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/diegofxm/ubl21-dian/documents/common"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// Builder construye facturas usando templates
type Builder struct {
	data InvoiceTemplateData
}

// NewBuilder crea un nuevo builder basado en templates
func NewBuilder() *Builder {
	return &Builder{
		data: InvoiceTemplateData{
			CurrencyCode:    "COP",
			InvoiceTypeCode: "01",
			Environment:     "2", // Habilitación por defecto
		},
	}
}

// SetInvoiceData establece los datos básicos de la factura
func (b *Builder) SetInvoiceData(number, cufe, issueDate, issueTime, dueDate string) *Builder {
	b.data.InvoiceNumber = number
	b.data.CUFE = cufe
	b.data.IssueDate = issueDate
	b.data.IssueTime = issueTime
	b.data.DueDate = dueDate
	return b
}

// SetDianExtensions establece los datos de extensiones DIAN
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

// SetSupplier establece los datos del proveedor
func (b *Builder) SetSupplier(supplier PartyTemplateData) *Builder {
	b.data.Supplier = supplier
	return b
}

// SetCustomer establece los datos del cliente
func (b *Builder) SetCustomer(customer PartyTemplateData) *Builder {
	b.data.Customer = customer
	return b
}

// SetDelivery establece los datos de entrega
func (b *Builder) SetDelivery(delivery *DeliveryTemplateData) *Builder {
	b.data.Delivery = delivery
	return b
}

// SetPaymentMeans agrega un medio de pago
func (b *Builder) SetPaymentMeans(id, code, dueDate string) *Builder {
	b.data.PaymentMeans = append(b.data.PaymentMeans, PaymentMeansTemplateData{
		ID:      id,
		Code:    code,
		DueDate: dueDate,
	})
	return b
}

// SetMonetaryTotals establece los totales monetarios
func (b *Builder) SetMonetaryTotals(lineExt, taxExcl, taxIncl, prepaid, payable string) *Builder {
	b.data.LineExtensionAmount = lineExt
	b.data.TaxExclusiveAmount = taxExcl
	b.data.TaxInclusiveAmount = taxIncl
	b.data.PrepaidAmount = prepaid
	b.data.PayableAmount = payable
	return b
}

// AddInvoiceLine agrega una línea de factura
func (b *Builder) AddInvoiceLine(line InvoiceLineTemplateData) *Builder {
	b.data.InvoiceLines = append(b.data.InvoiceLines, line)
	b.data.LineCount = len(b.data.InvoiceLines)
	return b
}

// SetProfileExecutionID establece el ambiente (1=Producción, 2=Habilitación)
func (b *Builder) SetProfileExecutionID(env string) *Builder {
	b.data.ProfileExecutionID = env
	b.data.Environment = env
	return b
}

// SetNote agrega una nota a la factura
func (b *Builder) SetNote(note string) *Builder {
	b.data.Notes = append(b.data.Notes, note)
	return b
}

// AddTaxTotal agrega un total de impuestos
func (b *Builder) AddTaxTotal(taxTotal TaxTotalTemplateData) *Builder {
	b.data.TaxTotals = append(b.data.TaxTotals, taxTotal)
	return b
}

// Build genera el XML de la factura usando templates
func (b *Builder) Build() ([]byte, error) {
	// Cargar templates comunes + específicos
	tmpl, err := common.LoadCommonAndSpecificTemplates(templatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("error loading templates: %w", err)
	}

	// Ejecutar template principal
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "invoice.tmpl", b.data); err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buf.Bytes(), nil
}

// GetData retorna los datos actuales del builder (útil para debugging)
func (b *Builder) GetData() InvoiceTemplateData {
	return b.data
}
