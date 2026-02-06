package debitnote

// DeliveryTemplateData información de entrega
type DeliveryTemplateData struct {
	ActualDeliveryDate string
	Address            AddressTemplateData
}

// AllowanceChargeTemplateData descuentos o cargos
type AllowanceChargeTemplateData struct {
	ChargeIndicator bool
	Reason          string
	MultiplierFactor string
	Amount          string
	BaseAmount      string
	CurrencyID      string
}
