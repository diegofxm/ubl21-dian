package types

// DeliveryXML entrega
type DeliveryXML struct {
	ActualDeliveryDate *CBCElement `xml:"cbc:ActualDeliveryDate,omitempty"`
	DeliveryAddress    *AddressXML `xml:"cac:DeliveryAddress,omitempty"`
}
