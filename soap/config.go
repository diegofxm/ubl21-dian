package soap

import (
	"github.com/diegofxm/ubl21-dian/soap/types"
)

// URLs de servicios DIAN
const (
	URLProduccion   = "https://vpfe.dian.gov.co/WcfDianCustomerServices.svc"
	URLHabilitacion = "https://vpfe-hab.dian.gov.co/WcfDianCustomerServices.svc"
)

// GetURL retorna la URL según el environment
func GetURL(env types.Environment) string {
	if env == types.Produccion {
		return URLProduccion
	}
	return URLHabilitacion
}
