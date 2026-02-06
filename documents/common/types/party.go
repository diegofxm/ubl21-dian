package types

// PartyXML representa una parte (emisor o receptor)
type PartyXML struct {
	IndustryClassificationCode *CBCElement           `xml:"cbc:IndustryClassificationCode,omitempty"`
	PartyName                  []PartyNameXML        `xml:"cac:PartyName,omitempty"`
	PhysicalLocation           *LocationXML          `xml:"cac:PhysicalLocation,omitempty"`
	PartyTaxScheme             []PartyTaxSchemeXML   `xml:"cac:PartyTaxScheme"`
	PartyLegalEntity           []PartyLegalEntityXML `xml:"cac:PartyLegalEntity"`
	Contact                    *ContactXML           `xml:"cac:Contact,omitempty"`
	Person                     *PersonXML            `xml:"cac:Person,omitempty"`
}

// PartyNameXML nombre de la parte
type PartyNameXML struct {
	Name CBCElement `xml:"cbc:Name"`
}

// LocationXML ubicación física
type LocationXML struct {
	Address AddressXML `xml:"cac:Address"`
}

// AddressXML dirección
type AddressXML struct {
	ID                   *CBCElement     `xml:"cbc:ID,omitempty"`
	CityName             CBCElement      `xml:"cbc:CityName"`
	PostalZone           *CBCElement     `xml:"cbc:PostalZone,omitempty"`
	CountrySubentity     CBCElement      `xml:"cbc:CountrySubentity"`
	CountrySubentityCode CBCElement      `xml:"cbc:CountrySubentityCode"`
	AddressLine          *AddressLineXML `xml:"cac:AddressLine,omitempty"`
	Country              CountryXML      `xml:"cac:Country"`
}

// AddressLineXML línea de dirección
type AddressLineXML struct {
	Line CBCElement `xml:"cbc:Line"`
}

// CountryXML país
type CountryXML struct {
	IdentificationCode CBCElement         `xml:"cbc:IdentificationCode"`
	Name               *CountryNameElement `xml:"cbc:Name,omitempty"`
}

// PartyTaxSchemeXML esquema tributario de la parte
type PartyTaxSchemeXML struct {
	RegistrationName    CBCElement          `xml:"cbc:RegistrationName"`
	CompanyID           IDElement           `xml:"cbc:CompanyID"`
	TaxLevelCode        TaxLevelCodeElement `xml:"cbc:TaxLevelCode,omitempty"`
	RegistrationAddress *AddressXML         `xml:"cac:RegistrationAddress,omitempty"`
	TaxScheme           TaxSchemeXML        `xml:"cac:TaxScheme"`
}

// PartyLegalEntityXML entidad legal
type PartyLegalEntityXML struct {
	RegistrationName            CBCElement                      `xml:"cbc:RegistrationName"`
	CompanyID                   IDElement                       `xml:"cbc:CompanyID"`
	CorporateRegistrationScheme *CorporateRegistrationSchemeXML `xml:"cac:CorporateRegistrationScheme,omitempty"`
}

// CorporateRegistrationSchemeXML esquema de registro corporativo
type CorporateRegistrationSchemeXML struct {
	ID   *CBCElement `xml:"cbc:ID,omitempty"`
	Name *CBCElement `xml:"cbc:Name,omitempty"`
}

// ContactXML información de contacto
type ContactXML struct {
	Name           *CBCElement `xml:"cbc:Name,omitempty"`
	Telephone      *CBCElement `xml:"cbc:Telephone,omitempty"`
	Telefax        *CBCElement `xml:"cbc:Telefax,omitempty"`
	ElectronicMail *CBCElement `xml:"cbc:ElectronicMail,omitempty"`
}

// PersonXML información de persona
type PersonXML struct {
	FirstName  CBCElement  `xml:"cbc:FirstName"`
	FamilyName CBCElement  `xml:"cbc:FamilyName"`
	MiddleName *CBCElement `xml:"cbc:MiddleName,omitempty"`
}
