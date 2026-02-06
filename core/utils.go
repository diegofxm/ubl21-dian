package core

import (
	"fmt"
	"time"
)

// FormatDate formatea una fecha a YYYY-MM-DD
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTime formatea una hora a HH:MM:SS-05:00
func FormatTime(t time.Time) string {
	return t.Format("15:04:05-07:00")
}

// FormatAmount formatea un monto con 2 decimales
func FormatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}

// CalculateDV calcula el dígito de verificación de un NIT
func CalculateDV(nit string) int {
	// Implementación del algoritmo de DV colombiano
	weights := []int{71, 67, 59, 53, 47, 43, 41, 37, 29, 23, 19, 17, 13, 7, 3}
	sum := 0

	for i, digit := range nit {
		if digit >= '0' && digit <= '9' {
			sum += int(digit-'0') * weights[len(weights)-len(nit)+i]
		}
	}

	remainder := sum % 11
	if remainder == 0 || remainder == 1 {
		return remainder
	}

	return 11 - remainder
}
