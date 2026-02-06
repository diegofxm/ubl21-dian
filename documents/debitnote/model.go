package debitnote

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/model"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// DebitNoteXML representa una nota débito electrónica UBL 2.1
type DebitNoteXML struct {
	XMLName xml.Name `xml:"DebitNote"`

	// Namespaces
	Xmlns             string `xml:"xmlns,attr"`
	XmlnsCAC          string `xml:"xmlns:cac,attr"`
	XmlnsCBC          string `xml:"xmlns:cbc,attr"`
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
	DebitNoteTypeCode  types.CBCElement  `xml:"cbc:DebitNoteTypeCode"`

	// Notas
	Note []types.CBCElement `xml:"cbc:Note,omitempty"`

	// Moneda
	DocumentCurrencyCode types.CurrencyCodeElement `xml:"cbc:DocumentCurrencyCode"`
	LineCountNumeric     types.CBCElement          `xml:"cbc:LineCountNumeric"`

	// Referencia a factura
	BillingReference []model.BillingReferenceXML `xml:"cac:BillingReference,omitempty"`

	// Partes
	AccountingSupplierParty model.AccountingPartyXML `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty model.AccountingPartyXML `xml:"cac:AccountingCustomerParty"`

	// Entrega
	Delivery *types.DeliveryXML `xml:"cac:Delivery,omitempty"`

	// Medios de pago
	PaymentMeans []types.PaymentMeansXML `xml:"cac:PaymentMeans,omitempty"`

	// Descuentos/Cargos
	AllowanceCharge []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`

	// Totales de impuestos
	TaxTotal []types.TaxTotalXML `xml:"cac:TaxTotal,omitempty"`

	// Totales monetarios (RequestedMonetaryTotal es específico de DebitNote)
	RequestedMonetaryTotal RequestedMonetaryTotalXML `xml:"cac:RequestedMonetaryTotal"`

	// Líneas
	DebitNoteLine []DebitNoteLineXML `xml:"cac:DebitNoteLine"`
}

// RequestedMonetaryTotalXML totales monetarios solicitados (específico de DebitNote)
type RequestedMonetaryTotalXML struct {
	LineExtensionAmount  types.AmountElement  `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount   types.AmountElement  `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount   types.AmountElement  `xml:"cbc:TaxInclusiveAmount"`
	AllowanceTotalAmount *types.AmountElement `xml:"cbc:AllowanceTotalAmount,omitempty"`
	ChargeTotalAmount    *types.AmountElement `xml:"cbc:ChargeTotalAmount,omitempty"`
	PayableAmount        types.AmountElement  `xml:"cbc:PayableAmount"`
}

// DebitNoteLineXML línea de nota débito
type DebitNoteLineXML struct {
	ID                    types.CBCElement           `xml:"cbc:ID"`
	DebitedQuantity       types.QuantityElement      `xml:"cbc:DebitedQuantity"`
	LineExtensionAmount   types.AmountElement        `xml:"cbc:LineExtensionAmount"`
	FreeOfChargeIndicator *types.CBCElement          `xml:"cbc:FreeOfChargeIndicator,omitempty"`
	AllowanceCharge       []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`
	TaxTotal              []types.TaxTotalXML        `xml:"cac:TaxTotal,omitempty"`
	Item                  types.ItemXML              `xml:"cac:Item"`
	Price                 types.PriceXML             `xml:"cac:Price"`
}
