package envelope

import (
	"bytes"
	_ "embed"
	"text/template"
)

// NOTA: Los templates se cargan desde disco en runtime porque go:embed no soporta paths relativos con ../
const (
	envelopeTemplatePath = "ubl21-dian/soap/templates/envelope.tmpl"
)

// Builder construye el SOAP Envelope completo
type Builder struct {
	securityHeader string
	body           string
}

// New crea un nuevo SOAP envelope builder
func New(securityHeader, body string) *Builder {
	return &Builder{
		securityHeader: securityHeader,
		body:           body,
	}
}

// Build construye el XML completo del SOAP envelope usando template
func (e *Builder) Build() string {
	tmpl, err := template.ParseFiles(envelopeTemplatePath)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, map[string]string{
		"SecurityHeader": e.securityHeader,
		"Body":           e.body,
	}); err != nil {
		return ""
	}

	return buffer.String()
}
