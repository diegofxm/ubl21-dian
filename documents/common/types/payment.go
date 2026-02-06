package types

// PaymentMeansXML medio de pago
type PaymentMeansXML struct {
	ID               CBCElement  `xml:"cbc:ID"`
	PaymentMeansCode CBCElement  `xml:"cbc:PaymentMeansCode"`
	PaymentDueDate   *CBCElement `xml:"cbc:PaymentDueDate,omitempty"`
	PaymentID        *CBCElement `xml:"cbc:PaymentID,omitempty"`
}

// PaymentTermsXML términos de pago
type PaymentTermsXML struct {
	ReferenceEventCode *CBCElement          `xml:"cbc:ReferenceEventCode,omitempty"`
	SettlementPeriod   *SettlementPeriodXML `xml:"cac:SettlementPeriod,omitempty"`
}

// SettlementPeriodXML período de liquidación
type SettlementPeriodXML struct {
	DurationMeasure DurationMeasureElement `xml:"cbc:DurationMeasure"`
}

// PrepaidPaymentXML pago anticipado
type PrepaidPaymentXML struct {
	ID         *CBCElement   `xml:"cbc:ID,omitempty"`
	PaidAmount AmountElement `xml:"cbc:PaidAmount"`
}
