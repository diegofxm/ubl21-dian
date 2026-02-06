package xml

import (
	"fmt"
	"strings"
	"unicode"
)

// Sanitize limpia una cadena de texto para uso en XML
// Elimina saltos de línea, retornos de carro y espacios múltiples
// DEBE usarse ANTES de construir el XML para evitar problemas de firma
func Sanitize(s string) string {
	// Reemplazar saltos de línea y retornos de carro con espacios
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\t", " ")

	// Eliminar espacios múltiples
	s = collapseSpaces(s)

	// Trim espacios al inicio y final
	s = strings.TrimSpace(s)

	return s
}

// SanitizeStrict aplica sanitización estricta
// Además de Sanitize(), elimina caracteres de control y no imprimibles
func SanitizeStrict(s string) string {
	// Primero aplicar sanitización básica
	s = Sanitize(s)

	// Eliminar caracteres de control y no imprimibles
	s = removeControlChars(s)

	return s
}

// collapseSpaces reemplaza múltiples espacios consecutivos con uno solo
func collapseSpaces(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	lastWasSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !lastWasSpace {
				result.WriteRune(' ')
				lastWasSpace = true
			}
		} else {
			result.WriteRune(r)
			lastWasSpace = false
		}
	}

	return result.String()
}

// removeControlChars elimina caracteres de control excepto espacios
func removeControlChars(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	for _, r := range s {
		// Mantener solo caracteres imprimibles y espacios
		if unicode.IsPrint(r) || r == ' ' {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// SanitizeAmount formatea un monto monetario de manera consistente
// Siempre usa 2 decimales, formato: 1234.56
func SanitizeAmount(amount float64) string {
	return formatFloat(amount, 2)
}

// SanitizePercent formatea un porcentaje de manera consistente
// Usa 2 decimales, formato: 19.00
func SanitizePercent(percent float64) string {
	return formatFloat(percent, 2)
}

// SanitizeQuantity formatea una cantidad de manera consistente
// Usa 6 decimales para mayor precisión
func SanitizeQuantity(quantity float64) string {
	return formatFloat(quantity, 6)
}

// formatFloat formatea un float64 con precisión específica
// Siempre mantiene exactamente 'decimals' decimales para consistencia
func formatFloat(value float64, decimals int) string {
	// Usar fmt.Sprintf para formato consistente
	format := fmt.Sprintf("%%.%df", decimals)
	return fmt.Sprintf(format, value)
}
