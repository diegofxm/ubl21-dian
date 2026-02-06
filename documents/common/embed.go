package common

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*.tmpl
var commonTemplatesFS embed.FS

// LoadCommonTemplates carga los templates comunes embebidos
// Retorna un template con todos los templates comunes parseados
func LoadCommonTemplates() (*template.Template, error) {
	tmpl := template.New("common")
	
	// Parsear todos los templates comunes embebidos
	tmpl, err := tmpl.ParseFS(commonTemplatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("error parsing common templates: %w", err)
	}
	
	return tmpl, nil
}

// LoadCommonAndSpecificTemplates carga templates comunes + específicos
// specificFS: filesystem embebido del paquete específico (invoice, creditnote, etc.)
// specificPattern: patrón para los templates específicos (ej: "templates/*.tmpl")
func LoadCommonAndSpecificTemplates(specificFS embed.FS, specificPattern string) (*template.Template, error) {
	// Cargar templates comunes
	tmpl, err := LoadCommonTemplates()
	if err != nil {
		return nil, err
	}
	
	// Cargar templates específicos
	tmpl, err = tmpl.ParseFS(specificFS, specificPattern)
	if err != nil {
		return nil, fmt.Errorf("error parsing specific templates: %w", err)
	}
	
	return tmpl, nil
}
