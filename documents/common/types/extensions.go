package types

// UBLExtensions contiene las extensiones UBL
type UBLExtensions struct {
	UBLExtension []UBLExtension `xml:"ext:UBLExtension"`
}

// UBLExtension representa una extensión UBL individual
type UBLExtension struct {
	ExtensionContent ExtensionContent `xml:"ext:ExtensionContent"`
}

// ExtensionContent contiene el contenido de la extensión
type ExtensionContent struct {
	InnerXML string `xml:",innerxml"`
}
