package invoice

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/model"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// InvoiceXML representa una factura electrónica UBL 2.1 con tags XML
type InvoiceXML struct {
	XMLName xml.Name `xml:"Invoice"`

	// Namespaces (orden exacto para coincidir con C14N)
	Xmlns             string `xml:"xmlns,attr"`
	XmlnsCAC          string `xml:"xmlns:cac,attr"`
	XmlnsCBC          string `xml:"xmlns:cbc,attr"`
	XmlnsCCTS         string `xml:"xmlns:ccts,attr"`
	XmlnsDS           string `xml:"xmlns:ds,attr"`
	XmlnsExt          string `xml:"xmlns:ext,attr"`
	XmlnsSts          string `xml:"xmlns:sts,attr"`
	XmlnsXades        string `xml:"xmlns:xades,attr"`
	XmlnsXades141     string `xml:"xmlns:xades141,attr"`
	XmlnsXsi          string `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation string `xml:"xsi:schemaLocation,attr"`

	// UBLExtensions
	UBLExtensions types.UBLExtensions `xml:"ext:UBLExtensions"`

	// Identificación básica
	UBLVersionID       types.CBCElement  `xml:"cbc:UBLVersionID"`
	CustomizationID    types.CBCElement  `xml:"cbc:CustomizationID"`
	ProfileID          types.CBCElement  `xml:"cbc:ProfileID"`
	ProfileExecutionID types.CBCElement  `xml:"cbc:ProfileExecutionID"`
	ID                 types.CBCElement  `xml:"cbc:ID"`
	UUID               types.UUIDElement `xml:"cbc:UUID"`
	IssueDate          types.CBCElement  `xml:"cbc:IssueDate"`
	IssueTime          types.CBCElement  `xml:"cbc:IssueTime"`
	DueDate            *types.CBCElement `xml:"cbc:DueDate,omitempty"`
	InvoiceTypeCode    types.CBCElement  `xml:"cbc:InvoiceTypeCode"`

	// Notas
	Note []types.CBCElement `xml:"cbc:Note,omitempty"`

	// Moneda
	DocumentCurrencyCode types.CurrencyCodeElement `xml:"cbc:DocumentCurrencyCode"`
	LineCountNumeric     types.CBCElement          `xml:"cbc:LineCountNumeric"`

	// Período de facturación
	InvoicePeriod *InvoicePeriodXML `xml:"cac:InvoicePeriod,omitempty"`

	// Referencias de orden
	OrderReference       *OrderReferenceXML        `xml:"cac:OrderReference,omitempty"`
	BillingReference     []model.BillingReferenceXML `xml:"cac:BillingReference,omitempty"`

	// Partes
	AccountingSupplierParty model.AccountingPartyXML `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty model.AccountingPartyXML `xml:"cac:AccountingCustomerParty"`

	// Entrega
	Delivery *types.DeliveryXML `xml:"cac:Delivery,omitempty"`

	// Medios de pago
	PaymentMeans []types.PaymentMeansXML `xml:"cac:PaymentMeans,omitempty"`
	PaymentTerms []types.PaymentTermsXML `xml:"cac:PaymentTerms,omitempty"`

	// Pagos anticipados
	PrepaidPayment *types.PrepaidPaymentXML `xml:"cac:PrepaidPayment,omitempty"`

	// Descuentos/Cargos globales
	AllowanceCharge []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`

	// Totales de impuestos
	TaxTotal []types.TaxTotalXML `xml:"cac:TaxTotal,omitempty"`

	// Retenciones (withholding tax)
	WithholdingTaxTotal []types.TaxTotalXML `xml:"cac:WithholdingTaxTotal,omitempty"`

	// Totales monetarios
	LegalMonetaryTotal LegalMonetaryTotalXML `xml:"cac:LegalMonetaryTotal"`

	// Líneas de factura
	InvoiceLine []InvoiceLineXML `xml:"cac:InvoiceLine"`
}

// InvoicePeriodXML período de facturación
type InvoicePeriodXML struct {
	StartDate types.CBCElement `xml:"cbc:StartDate"`
	EndDate   types.CBCElement `xml:"cbc:EndDate"`
}

// OrderReferenceXML referencia de orden
type OrderReferenceXML struct {
	ID types.CBCElement `xml:"cbc:ID"`
}

// LegalMonetaryTotalXML totales monetarios legales
type LegalMonetaryTotalXML struct {
	LineExtensionAmount  types.AmountElement  `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount   types.AmountElement  `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount   types.AmountElement  `xml:"cbc:TaxInclusiveAmount"`
	AllowanceTotalAmount *types.AmountElement `xml:"cbc:AllowanceTotalAmount,omitempty"`
	ChargeTotalAmount    *types.AmountElement `xml:"cbc:ChargeTotalAmount,omitempty"`
	PrepaidAmount        *types.AmountElement `xml:"cbc:PrepaidAmount,omitempty"`
	PayableAmount        types.AmountElement  `xml:"cbc:PayableAmount"`
}

// InvoiceLineXML línea de factura
type InvoiceLineXML struct {
	ID                    types.CBCElement             `xml:"cbc:ID"`
	InvoicedQuantity      types.QuantityElement        `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount   types.AmountElement          `xml:"cbc:LineExtensionAmount"`
	FreeOfChargeIndicator *types.CBCElement            `xml:"cbc:FreeOfChargeIndicator,omitempty"`
	AllowanceCharge       []types.AllowanceChargeXML   `xml:"cac:AllowanceCharge,omitempty"`
	TaxTotal              []types.TaxTotalXML          `xml:"cac:TaxTotal,omitempty"`
	WithholdingTaxTotal   []types.TaxTotalXML          `xml:"cac:WithholdingTaxTotal,omitempty"`
	Item                  types.ItemXML                `xml:"cac:Item"`
	Price                 types.PriceXML               `xml:"cac:Price"`
}
