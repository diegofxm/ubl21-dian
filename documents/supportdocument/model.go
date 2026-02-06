package supportdocument

import (
	"encoding/xml"

	"github.com/diegofxm/ubl21-dian/documents/common/model"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

// SupportDocumentXML representa un Documento Soporte UBL 2.1 con tags XML
type SupportDocumentXML struct {
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
	InvoiceTypeCode    types.CBCElement  `xml:"cbc:InvoiceTypeCode"` // "05" para Documento Soporte

	// Notas
	Note []types.CBCElement `xml:"cbc:Note,omitempty"`

	// Moneda
	DocumentCurrencyCode types.CurrencyCodeElement `xml:"cbc:DocumentCurrencyCode"`
	LineCountNumeric     types.CBCElement          `xml:"cbc:LineCountNumeric"`

	// Referencias a documentos externos (facturas de proveedor)
	BillingReference []model.BillingReferenceXML `xml:"cac:BillingReference,omitempty"`

	// Partes (Comprador como emisor, Proveedor como receptor)
	AccountingSupplierParty model.AccountingPartyXML `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty model.AccountingPartyXML `xml:"cac:AccountingCustomerParty"`

	// Entrega
	Delivery *types.DeliveryXML `xml:"cac:Delivery,omitempty"`

	// Medios de pago
	PaymentMeans []types.PaymentMeansXML `xml:"cac:PaymentMeans,omitempty"`
	PaymentTerms []types.PaymentTermsXML `xml:"cac:PaymentTerms,omitempty"`

	// Descuentos/Cargos globales
	AllowanceCharge []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`

	// Totales de impuestos
	TaxTotal []types.TaxTotalXML `xml:"cac:TaxTotal,omitempty"`

	// Retenciones (withholding tax) - IMPORTANTE para documento soporte
	WithholdingTaxTotal []types.TaxTotalXML `xml:"cac:WithholdingTaxTotal,omitempty"`

	// Totales monetarios
	LegalMonetaryTotal LegalMonetaryTotalXML `xml:"cac:LegalMonetaryTotal"`

	// Líneas del documento
	InvoiceLine []InvoiceLineXML `xml:"cac:InvoiceLine"`
}

// LegalMonetaryTotalXML totales monetarios del documento soporte
type LegalMonetaryTotalXML struct {
	LineExtensionAmount types.AmountElement `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  types.AmountElement `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  types.AmountElement `xml:"cbc:TaxInclusiveAmount"`
	AllowanceTotalAmount *types.AmountElement `xml:"cbc:AllowanceTotalAmount,omitempty"`
	ChargeTotalAmount    *types.AmountElement `xml:"cbc:ChargeTotalAmount,omitempty"`
	PayableAmount       types.AmountElement `xml:"cbc:PayableAmount"`
}

// InvoiceLineXML línea del documento soporte
type InvoiceLineXML struct {
	ID                    types.CBCElement      `xml:"cbc:ID"`
	InvoicedQuantity      types.QuantityElement `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount   types.AmountElement   `xml:"cbc:LineExtensionAmount"`
	FreeOfChargeIndicator *types.CBCElement     `xml:"cbc:FreeOfChargeIndicator,omitempty"`
	
	// Descuentos/Cargos a nivel de línea
	AllowanceCharge []types.AllowanceChargeXML `xml:"cac:AllowanceCharge,omitempty"`
	
	// Impuestos de la línea
	TaxTotal []types.TaxTotalXML `xml:"cac:TaxTotal,omitempty"`
	
	// Retenciones de la línea
	WithholdingTaxTotal []types.TaxTotalXML `xml:"cac:WithholdingTaxTotal,omitempty"`
	
	// Item
	Item ItemXML `xml:"cac:Item"`
	
	// Precio
	Price PriceXML `xml:"cac:Price"`
}

// ItemXML item/producto del documento soporte
type ItemXML struct {
	Description                  []types.CBCElement                          `xml:"cbc:Description"`
	StandardItemIdentification   *types.StandardItemIdentificationXML        `xml:"cac:StandardItemIdentification,omitempty"`
	AdditionalItemIdentification *types.AdditionalItemIdentificationXML      `xml:"cac:AdditionalItemIdentification,omitempty"`
}

// PriceXML precio del item
type PriceXML struct {
	PriceAmount  types.AmountElement    `xml:"cbc:PriceAmount"`
	BaseQuantity *types.QuantityElement `xml:"cbc:BaseQuantity,omitempty"`
}
