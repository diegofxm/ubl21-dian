package signature

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// CalculateCUFE calcula el CUFE (Código Único de Factura Electrónica)
// Algoritmo: SHA-384 de la concatenación de campos específicos
func CalculateCUFE(
	invoiceNumber string,
	issueDate time.Time,
	issueTime string,
	taxExclusiveAmount float64,
	taxAmount1 float64, // IVA
	taxAmount2 float64, // INC
	taxAmount3 float64, // ICA
	payableAmount float64,
	supplierNIT string,
	customerNIT string,
	technicalKey string,
	environment string, // "1" = Producción, "2" = Habilitación
) string {
	// Formatear fecha y hora
	dateStr := issueDate.Format("2006-01-02")
	timeStr := issueTime

	// Formatear montos con 2 decimales
	taxExclusiveStr := fmt.Sprintf("%.2f", taxExclusiveAmount)
	taxAmount1Str := fmt.Sprintf("%.2f", taxAmount1)
	taxAmount2Str := fmt.Sprintf("%.2f", taxAmount2)
	taxAmount3Str := fmt.Sprintf("%.2f", taxAmount3)
	payableAmountStr := fmt.Sprintf("%.2f", payableAmount)

	// Concatenar campos
	cufeString := strings.Join([]string{
		invoiceNumber,
		dateStr,
		timeStr,
		taxExclusiveStr,
		"01", // Código impuesto IVA
		taxAmount1Str,
		"04", // Código impuesto INC
		taxAmount2Str,
		"03", // Código impuesto ICA
		taxAmount3Str,
		payableAmountStr,
		supplierNIT,
		customerNIT,
		technicalKey,
		environment,
	}, "")

	// Calcular SHA-384
	hash := sha512.New384()
	hash.Write([]byte(cufeString))
	cufe := hex.EncodeToString(hash.Sum(nil))

	return cufe
}

// CalculateCUDE calcula el CUDE (Código Único de Documento Electrónico)
// Para Notas Crédito y Débito
func CalculateCUDE(
	documentNumber string,
	issueDate time.Time,
	issueTime string,
	taxExclusiveAmount float64,
	taxAmount1 float64,
	taxAmount2 float64,
	taxAmount3 float64,
	payableAmount float64,
	supplierNIT string,
	customerNIT string,
	technicalKey string,
	environment string,
) string {
	// El CUDE usa el mismo algoritmo que el CUFE
	return CalculateCUFE(
		documentNumber,
		issueDate,
		issueTime,
		taxExclusiveAmount,
		taxAmount1,
		taxAmount2,
		taxAmount3,
		payableAmount,
		supplierNIT,
		customerNIT,
		technicalKey,
		environment,
	)
}

// CalculateSoftwareSecurityCode calcula el código de seguridad del software
// Algoritmo: SHA-384 de (SoftwareID + SoftwarePIN + InvoiceNumber)
func CalculateSoftwareSecurityCode(softwareID, softwarePIN, invoiceNumber string) string {
	input := softwareID + softwarePIN + invoiceNumber

	hash := sha512.New384()
	hash.Write([]byte(input))
	securityCode := hex.EncodeToString(hash.Sum(nil))

	return securityCode
}

// GenerateQRCode genera el contenido del código QR para DIAN
// Retorna solo la URL según el formato de DIAN
func GenerateQRCode(
	invoiceNumber string,
	issueDate time.Time,
	supplierNIT string,
	customerNIT string,
	taxExclusiveAmount float64,
	taxAmount float64,
	payableAmount float64,
	cufe string,
	environment string, // "1" = Producción, "2" = Habilitación
) string {
	baseURL := "https://catalogo-vpfe.dian.gov.co"
	if environment == "2" {
		baseURL = "https://catalogo-vpfe-hab.dian.gov.co"
	}

	// Solo retornar la URL como en factura_real.xml
	return fmt.Sprintf("%s/document/searchqr?documentkey=%s", baseURL, cufe)
}

// CalculateLineExtensionAmount calcula el monto de extensión de línea
func CalculateLineExtensionAmount(quantity, unitPrice float64) float64 {
	return quantity * unitPrice
}

// CalculateTaxAmount calcula el monto de impuesto
func CalculateTaxAmount(taxableAmount, taxPercent float64) float64 {
	return (taxableAmount * taxPercent) / 100
}
