
# Flujo Completo de Facturación Electrónica DIAN (Software Propio)

Este documento describe **paso a paso**, de forma **real y práctica**, el flujo completo que sigue un software propio para facturación electrónica ante la **DIAN (Colombia)**, incluyendo:

- Generación de XML UBL
- Firma XAdES-BES
- Envío SOAP
- Recepción y validación de respuesta
- Construcción del AttachedDocument

---

## Visión General

> **La DIAN solo recibe un ZIP con el Invoice UBL firmado.**
>
> Todo lo demás es **proceso interno del software** o **respuesta DIAN**.

El `AttachedDocument` **NO se envía a la DIAN**.  
Se genera **solo para el cliente**.

---

## Flujo General

```
FE-123.xml                 (Invoice UBL sin firma)
   ↓
FES-123.xml                (Invoice UBL firmado)
   ↓
FES-123.zip                (ZIP con invoice firmado)
   ↓
ReqFE-123.xml              (SOAP Request)
   ↓
DIAN
   ↓
RptaFE-123.xml             (SOAP Response)
   ↓
ApplicationResponse.xml
   ↓
AttachedDocument.xml
   ↓
ad123.xml                  (AttachedDocument firmado - cliente)
```

---

## 1. Generar Invoice UBL sin firmar

**Archivo:** `FE-{numero}.xml`

Documento **UBL 2.1 puro**, sin firma.

### Contiene:
- Datos del emisor
- Datos del adquirente
- Totales
- Impuestos
- CUFE (calculado, no firmado)

### No contiene:
- `<ds:Signature>`
- `<ext:UBLExtensions>` con firma

📌 **Nunca se envía a la DIAN**

---

## 2. Firmar Invoice (XAdES-BES)

**Archivo:** `FES-{numero}.xml`

Se agrega la firma digital dentro de:

```xml
<ext:UBLExtensions>
  <ext:UBLExtension>
    <ext:ExtensionContent>
      <ds:Signature>...</ds:Signature>
    </ext:ExtensionContent>
  </ext:UBLExtension>
</ext:UBLExtensions>
```

### Firma:
- Tipo: XAdES-BES
- Certificado: `.p12`
- Incluye:
  - SignedProperties
  - Referencia al documento
  - Certificado

📌 Este es el **Invoice válido ante DIAN**

---

## 3. Comprimir Invoice firmado

**Archivo:** `FES-{numero}.zip`

Contenido del ZIP:

```
FES-{numero}.xml
```

❌ No incluir:
- AttachedDocument
- ApplicationResponse
- Otros archivos

---

## 4. Construir SOAP Request

**Archivo:** `ReqFE-{numero}.xml`

Operación:
- `SendBillSync`

Estructura básica:

```xml
<soapenv:Envelope>
  <soapenv:Body>
    <SendBillSync>
      <fileName>FES-123.zip</fileName>
      <contentFile>BASE64_DEL_ZIP</contentFile>
    </SendBillSync>
  </soapenv:Body>
</soapenv:Envelope>
```

📌 El SOAP **NO se firma**

---

## 5. Enviar a la DIAN

- Endpoint habilitación o producción
- Autenticación TLS con certificado

La DIAN:
- Descomprime ZIP
- Valida UBL
- Valida firma XAdES
- Valida CUFE
- Valida numeración

---

## 6. Respuesta DIAN (SOAP)

**Archivo:** `RptaFE-{numero}.xml`

Contiene:

```xml
<IsValid>true</IsValid>
<StatusCode>00</StatusCode>
<XmlBase64Bytes>...</XmlBase64Bytes>
```

📌 `XmlBase64Bytes` es clave

---

## 7. Decodificar ApplicationResponse

**Archivo:** `ApplicationResponse-{numero}.xml`

Resultado de decodificar Base64.

Contiene:
- CUFE validado
- Estado DIAN
- Fecha
- Firma DIAN

📌 Este XML es la **aceptación legal**

---

## 8. Generar AttachedDocument

**Archivo:** `AttachedDocument-{numero}.xml`

Incluye en Base64:

- Invoice firmado
- ApplicationResponse

Estructura conceptual:

```xml
<AttachedDocument>
  <cac:ParentDocumentLineReference>
    <cac:DocumentReference>
      <cbc:DocumentType>Invoice</cbc:DocumentType>
      <cac:Attachment>
        <cbc:EmbeddedDocumentBinaryObject>
          BASE64(INVOICE)
        </cbc:EmbeddedDocumentBinaryObject>
      </cac:Attachment>
    </cac:DocumentReference>

    <cac:DocumentReference>
      <cbc:DocumentType>ApplicationResponse</cbc:DocumentType>
      <cac:Attachment>
        <cbc:EmbeddedDocumentBinaryObject>
          BASE64(APPLICATION_RESPONSE)
        </cbc:EmbeddedDocumentBinaryObject>
      </cac:Attachment>
    </cac:DocumentReference>
  </cac:ParentDocumentLineReference>
</AttachedDocument>
```

📌 **No se envía a DIAN**

---

## 9. Firmar AttachedDocument

**Archivo final:** `ad{numero}.xml`

- Firma XAdES-BES
- Mismo certificado
- Documento final para el cliente

---

## Errores comunes

❌ Enviar AttachedDocument a DIAN  
❌ Firmar SOAP  
❌ Incluir ApplicationResponse en ZIP  
❌ Enviar Invoice sin firma  
❌ Calcular CUFE después de firmar  

---

## Conclusión

- DIAN solo recibe: **ZIP con Invoice firmado**
- AttachedDocument es solo para el cliente
- El flujo correcto evita el 90% de errores de habilitación

---

**Autor:** Guía técnica – Software Propio DIAN  
