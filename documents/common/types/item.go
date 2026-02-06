package types

// ItemXML artículo/producto
type ItemXML struct {
	Description                  []CBCElement                      `xml:"cbc:Description,omitempty"`
	BrandName                    *CBCElement                       `xml:"cbc:BrandName,omitempty"`
	ModelName                    *CBCElement                       `xml:"cbc:ModelName,omitempty"`
	StandardItemIdentification   *StandardItemIdentificationXML    `xml:"cac:StandardItemIdentification,omitempty"`
	AdditionalItemIdentification *AdditionalItemIdentificationXML  `xml:"cac:AdditionalItemIdentification,omitempty"`
	SellersItemIdentification    *SellersItemIdentificationXML     `xml:"cac:SellersItemIdentification,omitempty"`
}

// StandardItemIdentificationXML identificación estándar del artículo
type StandardItemIdentificationXML struct {
	ID IDElement `xml:"cbc:ID"`
}

// AdditionalItemIdentificationXML identificación adicional del artículo
type AdditionalItemIdentificationXML struct {
	ID IDElement `xml:"cbc:ID"`
}

// SellersItemIdentificationXML identificación del vendedor
type SellersItemIdentificationXML struct {
	ID CBCElement `xml:"cbc:ID"`
}

// PriceXML precio
type PriceXML struct {
	PriceAmount  AmountElement    `xml:"cbc:PriceAmount"`
	BaseQuantity *QuantityElement `xml:"cbc:BaseQuantity,omitempty"`
}
