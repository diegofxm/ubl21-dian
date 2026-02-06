package types

// TaxSchemeXML esquema tributario
type TaxSchemeXML struct {
	ID   CBCElement `xml:"cbc:ID"`
	Name CBCElement `xml:"cbc:Name"`
}

// TaxTotalXML total de impuestos
type TaxTotalXML struct {
	TaxAmount   AmountElement    `xml:"cbc:TaxAmount"`
	TaxSubtotal []TaxSubtotalXML `xml:"cac:TaxSubtotal,omitempty"`
}

// TaxSubtotalXML subtotal de impuesto
type TaxSubtotalXML struct {
	TaxableAmount AmountElement  `xml:"cbc:TaxableAmount"`
	TaxAmount     AmountElement  `xml:"cbc:TaxAmount"`
	TaxCategory   TaxCategoryXML `xml:"cac:TaxCategory"`
}

// TaxCategoryXML categoría de impuesto
type TaxCategoryXML struct {
	Percent   *CBCElement  `xml:"cbc:Percent,omitempty"`
	TaxScheme TaxSchemeXML `xml:"cac:TaxScheme"`
}
