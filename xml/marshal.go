package xml

import (
	"encoding/xml"
)

// Marshal serializa una estructura a XML sin indentación
// CRÍTICO: No usar MarshalIndent para documentos que serán firmados
// La firma digital requiere que el XML sea idéntico byte a byte
func Marshal(v interface{}) ([]byte, error) {
	// Usar Marshal sin indent para XML determinístico
	data, err := xml.Marshal(v)
	if err != nil {
		return nil, err
	}

	// Agregar declaración XML al inicio con standalone="no" como requiere DIAN
	xmlDeclaration := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>`)
	result := append(xmlDeclaration, data...)

	return result, nil
}

// MarshalNoHeader serializa una estructura a XML sin la declaración XML
// Útil para elementos que se insertarán dentro de otro documento
func MarshalNoHeader(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
