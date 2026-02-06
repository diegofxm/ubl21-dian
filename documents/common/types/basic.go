package types

// CBCElement elemento básico con texto
type CBCElement struct {
	Value string `xml:",chardata"`
}

// IDElement elemento ID con atributos (orden alfabético para C14N)
type IDElement struct {
	SchemeAgencyID   string `xml:"schemeAgencyID,attr,omitempty"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr,omitempty"`
	SchemeID         string `xml:"schemeID,attr,omitempty"`
	SchemeName       string `xml:"schemeName,attr,omitempty"`
	SchemeVersionID  string `xml:"schemeVersionID,attr,omitempty"`
	Value            string `xml:",chardata"`
}

// UUIDElement elemento UUID con atributos
type UUIDElement struct {
	SchemeID   string `xml:"schemeID,attr,omitempty"`
	SchemeName string `xml:"schemeName,attr,omitempty"`
	Value      string `xml:",chardata"`
}

// CurrencyCodeElement elemento de código de moneda con atributos
type CurrencyCodeElement struct {
	ListID         string `xml:"listID,attr,omitempty"`
	ListAgencyID   string `xml:"listAgencyID,attr,omitempty"`
	ListAgencyName string `xml:"listAgencyName,attr,omitempty"`
	Value          string `xml:",chardata"`
}

// AmountElement elemento de monto con atributo de moneda
type AmountElement struct {
	CurrencyID string `xml:"currencyID,attr"`
	Value      string `xml:",chardata"`
}

// QuantityElement elemento de cantidad con atributo de unidad
type QuantityElement struct {
	UnitCode string `xml:"unitCode,attr"`
	Value    string `xml:",chardata"`
}

// IdentificationCodeElement elemento con atributos para código de identificación
type IdentificationCodeElement struct {
	ListAgencyID   string `xml:"listAgencyID,attr,omitempty"`
	ListAgencyName string `xml:"listAgencyName,attr,omitempty"`
	ListSchemeURI  string `xml:"listSchemeURI,attr,omitempty"`
	Value          string `xml:",chardata"`
}

// CountryNameElement elemento nombre de país con languageID
type CountryNameElement struct {
	LanguageID string `xml:"languageID,attr,omitempty"`
	Value      string `xml:",chardata"`
}

// TaxLevelCodeElement código de nivel tributario con atributos
type TaxLevelCodeElement struct {
	ListName string `xml:"listName,attr"`
	Value    string `xml:",chardata"`
}

// DurationMeasureElement medida de duración con atributo
type DurationMeasureElement struct {
	UnitCode string `xml:"unitCode,attr"`
	Value    string `xml:",chardata"`
}

// CDATAElement elemento con contenido CDATA
type CDATAElement struct {
	Value string `xml:",cdata"`
}
