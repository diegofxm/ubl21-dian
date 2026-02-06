# ubl21-dian

Librería Go para generación de documentos electrónicos UBL 2.1 y envío a DIAN (Colombia).

## 🚀 Características

- ✅ Generación de XML UBL 2.1 con templates Go
- ✅ Firma digital XAdES-BES con canonicalización C14N
- ✅ Cálculo automático de CUFE/CUDE
- ✅ Cliente SOAP para envío a DIAN
- ✅ Soporte para múltiples tipos de documentos:
  - Factura Electrónica (01)
  - Nota Crédito (91)
  - Nota Débito (92)
  - Documento Soporte (05)
  - Nómina Electrónica (102, 103)
- ✅ Type-safe con validaciones
- ✅ Templates embebidos (no requiere archivos externos)

## 📦 Instalación

```bash
go get github.com/diegofxm/ubl21-dian
```

## 🔧 Uso Básico

### Crear una Factura

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/diegofxm/ubl21-dian/invoice"
    "github.com/diegofxm/ubl21-dian/core"
)

func main() {
    // 1. Crear factura con Builder
    inv, err := invoice.NewBuilder().
        SetID("FACT-001").
        SetIssueDate(time.Now()).
        SetSupplier(core.Party{
            Name: "Mi Empresa SAS",
            PartyIdentification: core.PartyIdentification{
                ID: "900123456",
                SchemeID: "31", // NIT
            },
        }).
        SetCustomer(core.Party{
            Name: "Cliente XYZ",
            PartyIdentification: core.PartyIdentification{
                ID: "800654321",
                SchemeID: "31",
            },
        }).
        AddLine(invoice.InvoiceLine{
            ID: "1",
            InvoicedQuantity: core.Quantity{
                Value: 10,
                UnitCode: "EA",
            },
            LineExtensionAmount: core.MonetaryAmount{
                Value: 100000,
                CurrencyID: "COP",
            },
        }).
        Build()
    
    if err != nil {
        panic(err)
    }
    
    // 2. Renderizar a XML
    renderer, _ := invoice.NewRenderer()
    xmlString, err := renderer.RenderString(inv)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(xmlString)
}
```

### Firmar y Enviar a DIAN

```go
import (
    "github.com/diegofxm/ubl21-dian/signature"
    "github.com/diegofxm/ubl21-dian/dian"
)

// 3. Firmar XML
signer, _ := signature.NewSigner("certificado.p12", "password")
signedXML, err := signer.SignXML(xmlString)

// 4. Enviar a DIAN
client := dian.NewClient(dian.Habilitacion)
response, err := client.SendDocument(signedXML, "TestSetId")
```

## 📁 Estructura del Proyecto

```
ubl21-dian/
├── core/           # Tipos compartidos (Party, Address, Tax, etc.)
├── xml/            # Motor de templates y canonicalización C14N
├── invoice/        # Módulo de facturas
├── creditnote/     # Módulo de notas crédito
├── debitnote/      # Módulo de notas débito
├── signature/      # Firma digital XAdES-BES
├── dian/           # Cliente SOAP para DIAN
└── examples/       # Ejemplos de uso
```

## 🔐 Requisitos

- Go 1.21 o superior
- Certificado digital (.p12) emitido por DIAN
- SoftwareID y PIN asignados por DIAN

## 📖 Documentación

Ver [ejemplos completos](./examples/) para casos de uso detallados.

## 📄 Licencia

MIT License

## 🤝 Contribuciones

Contribuciones son bienvenidas. Por favor abre un issue primero para discutir cambios mayores.
