package creditnote

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/model"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// CreditNoteXML representa una nota crédito electrónica UBL 2.1
type CreditNoteXML struct {
	XMLName xml.Name `xml:"CreditNote"`

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
	CreditNoteTypeCode types.CBCElement  `xml:"cbc:CreditNoteTypeCode"`

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

	// Totales monetarios
	LegalMonetaryTotal LegalMonetaryTotalXML `xml:"cac:LegalMonetaryTotal"`

	// Líneas
	CreditNoteLine []CreditNoteLineXML `xml:"cac:CreditNoteLine"`
}

// LegalMonetaryTotalXML totales monetarios
type LegalMonetaryTotalXML struct {
	LineExtensionAmount  types.AmountElement  `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount   types.AmountElement  `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount   types.AmountElement  `xml:"cbc:TaxInclusiveAmount"`
	AllowanceTotalAmount *types.AmountElement `xml:"cbc:AllowanceTotalAmount,omitempty"`
	ChargeTotalAmount    *types.AmountElement `xml:"cbc:ChargeTotalAmount,omitempty"`
	PayableAmount        types.AmountElement  `xml:"cbc:PayableAmount"`
}

// CreditNoteLineXML línea de nota crédito
type CreditNoteLineXML struct {
	ID                    types.CBCElement           `xml:"cbc:ID"`
	CreditedQuantity      types.QuantityElement      `xml:"cbc:CreditedQuantity"`
	LineExtensionAmount   types.AmountElement        `xml:"cbc:LineExtensionAmount"`
	FreeOfChargeIndicator *types.CBCElement          `xml:"cbc:FreeOfChargeIndicator,omitempty"`
	AllowanceCharge       []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`
	TaxTotal              []types.TaxTotalXML        `xml:"cac:TaxTotal,omitempty"`
	Item                  types.ItemXML              `xml:"cac:Item"`
	Price                 types.PriceXML             `xml:"cac:Price"`
}
