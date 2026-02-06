package xml

import (
	"fmt"
	"text/template"
	"time"
)

// TemplateFunctions retorna las funciones disponibles en templates
func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"formatDate":   formatDate,
		"formatTime":   formatTime,
		"formatAmount": formatAmount,
		"add":          add,
		"sub":          sub,
		"mul":          mul,
		"div":          div,
		"percent":      percent,
	}
}

// formatDate formatea una fecha a YYYY-MM-DD
func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// formatTime formatea una hora a HH:MM:SS-05:00
func formatTime(t time.Time) string {
	return t.Format("15:04:05-07:00")
}

// formatAmount formatea un monto con 2 decimales
func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}

// add suma dos números
func add(a, b float64) float64 {
	return a + b
}

// sub resta dos números
func sub(a, b float64) float64 {
	return a - b
}

// mul multiplica dos números
func mul(a, b float64) float64 {
	return a * b
}

// div divide dos números
func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// percent calcula el porcentaje
func percent(amount, percentage float64) float64 {
	return (amount * percentage) / 100
}
