# Arquitectura profesional UBL 2.1 DIAN en Go (sin templates)

## Contexto del problema

Al enviar documentos electrónicos a la **DIAN**, se presentan rechazos como:

- **ZE02 – Valor de la firma inválido**
- **FAJ43b – Nombre no corresponde al RUT**
- **FAR03 – SourceCurrencyBaseRate = 1.00**

Inicialmente se pensó que el problema era el **certificado digital**, pero el proveedor confirmó que:

> El certificado **sí es válido** y **funciona correctamente**.  
> El rechazo ocurre porque **el XML firmado no es idéntico al XML enviado**.

La causa más común:
- Saltos de línea invisibles
- Espacios adicionales
- Caracteres especiales no controlados
- Construcción del XML con templates (`.tmpl`)

---

## Qué quiso decir el proveedor (explicado simple)

Cuando DIAN valida una factura:

1. Toma el XML recibido
2. Recalcula la firma digital
3. Compara el resultado con la firma incluida

Si **un solo byte cambia**, la firma se invalida.

Esto pasa cuando:
- Se usan templates
- Se concatenan strings
- Se modifica el XML después de firmarlo
- Existen saltos de línea (`\n`, `\r`) o espacios ocultos

👉 **El certificado NO falla, falla el XML**

---

## Error clave: usar templates para XML

### ❌ Ejemplo con template (incorrecto)

```xml
<cbc:Description>
    {{ .Description }}
</cbc:Description>
```

Problemas:
- El template introduce saltos de línea
- El formato depende del archivo `.tmpl`
- No hay control byte a byte

Resultado: **firma inválida (ZE02)**

---

## Solución correcta: encoding/xml

Usar `encoding/xml` significa:

- NO escribir XML manualmente
- NO usar templates
- Construir el XML desde **structs Go**
- Dejar que Go genere un XML determinístico

---

## Ejemplo claro con encoding/xml

### Structs

```go
type Description struct {
    Text string `xml:",chardata"`
}

type InvoiceLine struct {
    Description Description `xml:"cbc:Description"`
}
```

### Sanitización (ANTES de firmar)

```go
func sanitize(s string) string {
    s = strings.ReplaceAll(s, "\n", " ")
    s = strings.ReplaceAll(s, "\r", " ")
    s = strings.TrimSpace(s)
    return s
}
```

### Construcción

```go
line := InvoiceLine{
    Description: Description{
        Text: sanitize("Servicio técnico\nincluye instalación"),
    },
}

xmlBytes, _ := xml.Marshal(line)
```

### Resultado seguro

```xml
<cbc:Description>Servicio técnico incluye instalación</cbc:Description>
```

✔ Sin saltos  
✔ Sin espacios ocultos  
✔ Firma válida  

---

## Propuesta de arquitectura modular profesional (Go)

```text
ubl21-dian/
├── core/                # Tipos UBL base (SOLO structs)
│   ├── party.go
│   ├── address.go
│   ├── tax.go
│   └── common.go
│
├── documents/
│   ├── invoice/
│   │   ├── model.go     # struct Invoice UBL
│   │   └── builder.go   # lógica de negocio
│   ├── creditnote/
│   └── debitnote/
│
├── xml/
│   ├── marshal.go       # xml.Marshal sin indent
│   └── sanitize.go      # limpieza de texto
│
├── signature/
│   ├── pkcs12.go        # carga .p12 desde base64
│   └── xades.go         # firma XAdES-BES
│
├── dian/
│   ├── soap.go
│   └── send.go
│
└── examples/
```

---

## Principios clave del sistema

### 1. El XML se genera UNA sola vez
- `xml.Marshal`
- Sin `Indent`
- Sin modificar después

### 2. La firma usa EXACTAMENTE esos bytes

```go
signed := signature.Sign(xmlBytes)
```

Nunca:
- `string(xmlBytes)`
- `fmt.Println(xmlBytes)`
- `xml.Unmarshal` nuevamente

---

## Regla de oro DIAN

> **Si tú puedes ver o editar el XML como string, ya es tarde para firmarlo.**

---

## Conclusión

- El certificado está bien
- OpenSSL no es el problema
- El problema es **cómo se construye el XML**
- `encoding/xml` es obligatorio para XML limpios y firmables
- La arquitectura modular propuesta es escalable y profesional

---

## Recomendación final

Migrar progresivamente:
1. Eliminar templates
2. Reemplazar por structs
3. Centralizar sanitización
4. Firmar solo `[]byte` finales

Con esto:
✅ ZE02 desaparece  
✅ XML válidos DIAN  
✅ Sistema robusto y profesional  
