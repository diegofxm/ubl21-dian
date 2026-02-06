package types

// AllowanceChargeXML descuento o cargo
type AllowanceChargeXML struct {
	ID                      *CBCElement    `xml:"cbc:ID,omitempty"`
	ChargeIndicator         CBCElement     `xml:"cbc:ChargeIndicator"`
	AllowanceChargeReason   *CBCElement    `xml:"cbc:AllowanceChargeReason,omitempty"`
	MultiplierFactorNumeric *CBCElement    `xml:"cbc:MultiplierFactorNumeric,omitempty"`
	Amount                  AmountElement  `xml:"cbc:Amount"`
	BaseAmount              *AmountElement `xml:"cbc:BaseAmount,omitempty"`
}
