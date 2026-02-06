package model

import "github.com/diegofxm/ubl21-dian/documents/common/types"

// AccountingPartyXML parte contable (supplier o customer) - común para todos los documentos
type AccountingPartyXML struct {
	AdditionalAccountID types.CBCElement `xml:"cbc:AdditionalAccountID"`
	Party               types.PartyXML   `xml:"cac:Party"`
}
