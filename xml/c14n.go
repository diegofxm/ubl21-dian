package xml

/*
#cgo pkg-config: libxml-2.0
#include <libxml/c14n.h>
#include <libxml/parser.h>
#include <libxml/tree.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Canonicalize canonicaliza XML según C14N 1.0 Inclusive (sin comentarios)
// Usa libxml2 (la misma librería que usa PHP DOMDocument::C14N)
// Como lo requiere el Anexo Técnico de DIAN sección 10.7
func Canonicalize(xmlData []byte) ([]byte, error) {
	// Inicializar libxml2 parser
	C.xmlInitParser()
	defer C.xmlCleanupParser()

	// Parse XML con libxml2
	cXML := C.CString(string(xmlData))
	defer C.free(unsafe.Pointer(cXML))

	doc := C.xmlReadMemory(
		cXML,
		C.int(len(xmlData)),
		nil, // URL
		nil, // encoding
		C.XML_PARSE_NONET, // options: no network access
	)
	if doc == nil {
		return nil, fmt.Errorf("failed to parse XML with libxml2")
	}
	defer C.xmlFreeDoc(doc)

	// Canonicalizar con C14N 1.0 (sin comentarios, modo inclusivo)
	var buf *C.xmlChar
	size := C.xmlC14NDocDumpMemory(
		doc,
		nil, // nodes (nil = todo el documento)
		C.XML_C14N_1_0, // mode: C14N 1.0 sin comentarios
		nil, // inclusive_ns_prefixes
		0,   // with_comments: 0 = sin comentarios
		&buf,
	)

	if size < 0 {
		return nil, fmt.Errorf("C14N canonicalization failed")
	}
	defer C.free(unsafe.Pointer(buf))

	// Convertir a []byte de Go
	canonical := C.GoBytes(unsafe.Pointer(buf), C.int(size))

	return canonical, nil
}

// CanonicalizeExclusive canonicaliza XML según Exclusive C14N (xml-exc-c14n#)
// Usado para SOAP Security Headers según WS-Security
func CanonicalizeExclusive(xmlData []byte, inclusiveNamespaces []string) ([]byte, error) {
	// Inicializar libxml2 parser
	C.xmlInitParser()
	defer C.xmlCleanupParser()

	// Parse XML con libxml2
	cXML := C.CString(string(xmlData))
	defer C.free(unsafe.Pointer(cXML))

	doc := C.xmlReadMemory(
		cXML,
		C.int(len(xmlData)),
		nil,
		nil,
		C.XML_PARSE_NONET,
	)
	if doc == nil {
		return nil, fmt.Errorf("failed to parse XML with libxml2")
	}
	defer C.xmlFreeDoc(doc)

	// Preparar lista de namespaces inclusivos
	var inclusiveNS **C.xmlChar
	if len(inclusiveNamespaces) > 0 {
		// Crear array de strings C terminado en NULL
		cStrings := make([]*C.char, len(inclusiveNamespaces)+1)
		for i, ns := range inclusiveNamespaces {
			cStrings[i] = C.CString(ns)
			defer C.free(unsafe.Pointer(cStrings[i]))
		}
		cStrings[len(inclusiveNamespaces)] = nil
		inclusiveNS = (**C.xmlChar)(unsafe.Pointer(&cStrings[0]))
	}

	// Canonicalizar con Exclusive C14N
	var buf *C.xmlChar
	size := C.xmlC14NDocDumpMemory(
		doc,
		nil,                 // nodes (nil = todo el documento)
		C.XML_C14N_EXCLUSIVE_1_0, // mode: Exclusive C14N 1.0
		inclusiveNS,         // inclusive_ns_prefixes
		0,                   // with_comments: 0 = sin comentarios
		&buf,
	)

	if size < 0 {
		return nil, fmt.Errorf("Exclusive C14N canonicalization failed")
	}
	defer C.free(unsafe.Pointer(buf))

	canonical := C.GoBytes(unsafe.Pointer(buf), C.int(size))
	return canonical, nil
}
