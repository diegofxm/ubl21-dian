# 📡 SOAP Client para DIAN

Cliente SOAP profesional para envío de facturas electrónicas a DIAN (Colombia).

## 🚀 Características

- ✅ **WS-Security 1.0** con firma digital X.509
- ✅ **WS-Addressing** para routing
- ✅ **Múltiples métodos de envío**: Sync, Async, TestSet
- ✅ **Consulta de estado** y descarga de respuestas
- ✅ **Manejo robusto de errores** SOAP Fault
- ✅ **Timeout configurable** y retry logic
- ✅ **Ambientes**: Producción y Habilitación

## 📦 Instalación

```bash
go get github.com/diegofxm/ubl21-dian/soap
```

## 🔧 Uso Básico

### 1. Crear Cliente

```go
import "github.com/diegofxm/ubl21-dian/soap"

client, err := soap.NewClient(&soap.Config{
    Environment: soap.Habilitacion, // o soap.Produccion
    Certificate: "path/to/certificate.pem",
    PrivateKey:  "path/to/private_key.pem", // Opcional si está en Certificate
    Timeout:     180 * time.Second,
})
```

### 2. Enviar Factura (TestSet)

```go
// Crear ZIP con XML firmado
zipData := createZIP("invoice.xml", signedXML)
zipBase64 := base64.StdEncoding.EncodeToString(zipData)

// Enviar a DIAN
response, err := client.SendTestSetAsync(&soap.SendTestSetAsyncRequest{
    FileName:    "invoice.zip",
    ContentFile: zipBase64,
    TestSetId:   "292537a5-3771-4d32-93ea-24d235565231",
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("IsValid: %v\n", response.IsValid)
fmt.Printf("StatusCode: %s\n", response.StatusCode)
fmt.Printf("StatusMessage: %s\n", response.StatusMessage)
```

### 3. Consultar Estado

```go
status, err := client.GetStatus(&soap.GetStatusRequest{
    TrackId: "tracking-id-from-dian",
})

fmt.Printf("Estado: %s\n", status.StatusMessage)
```

### 4. Descargar ZIP de Respuesta

```go
zipResp, err := client.GetStatusZip(&soap.GetStatusZipRequest{
    TrackId: "tracking-id",
})

// Decodificar y guardar ZIP
zipData, _ := base64.StdEncoding.DecodeString(zipResp.ContentFile)
os.WriteFile("response.zip", zipData, 0644)
```

## 📋 Métodos Disponibles

### Envío de Documentos

| Método | Descripción | Uso |
|--------|-------------|-----|
| `SendBillSync` | Envío síncrono de factura | Producción |
| `SendBillAsync` | Envío asíncrono de factura | Producción |
| `SendTestSetAsync` | Envío a set de pruebas | Habilitación |
| `SendBillAttachmentAsync` | Envío de documento adjunto | Producción |
| `SendEvent` | Envío de evento (acuse, rechazo) | Producción |

### Consultas

| Método | Descripción |
|--------|-------------|
| `GetStatus` | Consultar estado de documento |
| `GetStatusZip` | Descargar ZIP de respuesta DIAN |
| `GetXmlByDocumentKey` | Obtener XML por CUFE |

## 🔐 Seguridad (WS-Security)

El cliente implementa automáticamente:

- **BinarySecurityToken**: Certificado X.509 en base64
- **Timestamp**: Validez del mensaje (60 minutos)
- **Signature**: Firma digital RSA-SHA256 del elemento `wsa:To`
- **Canonicalización**: Exclusive C14N
- **SecurityTokenReference**: Referencia al certificado

## 📊 Estructura de Respuesta

```go
type Response struct {
    IsValid           bool
    StatusCode        string
    StatusDescription string
    StatusMessage     string
    ErrorMessage      []ErrorMessage
    XmlDocumentKey    string
    XmlBase64Bytes    string
}

type ErrorMessage struct {
    Code        string
    Description string
}
```

## ⚠️ Manejo de Errores

```go
response, err := client.SendTestSetAsync(req)
if err != nil {
    // Error de red, timeout, o SOAP Fault
    log.Fatal(err)
}

if !response.IsValid {
    // Documento rechazado por DIAN
    for _, errMsg := range response.ErrorMessage {
        fmt.Printf("[%s] %s\n", errMsg.Code, errMsg.Description)
    }
}
```

## 🌐 Ambientes

### Habilitación (Pruebas)
```go
Environment: soap.Habilitacion
URL: https://vpfe-hab.dian.gov.co/WcfDianCustomerServices.svc
```

### Producción
```go
Environment: soap.Produccion
URL: https://vpfe.dian.gov.co/WcfDianCustomerServices.svc
```

## 📝 Ejemplo Completo

Ver `examples/soap_send/main.go` para un ejemplo completo de:
1. Cargar XML firmado
2. Crear ZIP
3. Enviar a DIAN
4. Consultar estado
5. Manejar respuestas

## 🔗 Integración con otros módulos

```go
// 1. Generar factura
invoice := invoice.NewBuilder()...Build()

// 2. Renderizar XML
renderer := invoice.NewRenderer()
xmlString := renderer.RenderString(invoice)

// 3. Firmar XML
signer := signature.NewSignerFromSinglePEM("cert.pem")
signedXML := signer.SignXML([]byte(xmlString))

// 4. Enviar a DIAN
client := soap.NewClient(config)
response := client.SendTestSetAsync(...)
```

## 📚 Referencias

- [DIAN - Facturación Electrónica](https://www.dian.gov.co/fizcalizacioncontrol/herramienconsulta/FacturaElectronica/Paginas/default.aspx)
- [WS-Security 1.0](http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0.pdf)
- [WS-Addressing](https://www.w3.org/Submission/ws-addressing/)

## 📄 Licencia

MIT
