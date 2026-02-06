# Flujo Completo de Facturación Electrónica DIAN

## 📋 Hallazgos Clave del Análisis de co-apidian2021

### 🔍 Problema Identificado
Nuestra librería Go (`ubl21-dian`) **NO está generando el AttachedDocument**, que es **CRÍTICO** para el proceso completo de facturación electrónica DIAN.

---

## 📦 Archivos Generados por Factura (7 archivos)

### 1. **FE-SETP990000001.xml** (Factura sin firmar)
- XML UBL 2.1 Invoice sin firma digital
- Contiene todos los datos de la factura
- **Estado actual Go:** ✅ Se genera correctamente

### 2. **FES-SETP990000001.xml** (Factura firmada)
- Mismo XML pero con firma digital XAdES
- Incluye certificado digital y firma
- **Estado actual Go:** ✅ Se genera correctamente

### 3. **FES-SETP990000001.zip** (ZIP de factura firmada)
- Contiene el XML firmado comprimido
- Se envía a DIAN para validación
- **Estado actual Go:** ✅ Se genera correctamente

### 4. **ReqFE-SETP990000001.xml** (Request SOAP)
- Envelope SOAP con seguridad WS-Security
- Contiene el ZIP en base64
- **Estado actual Go:** ✅ Se genera correctamente

### 5. **RptaFE-SETP990000001.xml** (Response de DIAN)
- ApplicationResponse que devuelve DIAN
- Contiene:
  - `ResponseCode`: 02 (validado) o error
  - `Description`: "Documento validado por la DIAN"
  - CUDE (Código Único de Documento Electrónico)
  - Fecha y hora de validación
- **Estado actual Go:** ✅ Se recibe correctamente

### 6. **FES-SETP990000001.pdf** (PDF de la factura)
- Representación gráfica con QR
- **Estado actual Go:** ✅ Ejemplo creado (main_simple.go)

### 7. **AttachmentDocument-SETP990000001.xml** (Documento Adjunto) ⚠️
- **ESTE ES EL QUE FALTA EN GO**
- Se crea **DESPUÉS** de recibir respuesta exitosa de DIAN
- Contiene DOS documentos XML completos dentro de CDATA:
  1. La factura firmada completa
  2. El ApplicationResponse de DIAN
- Se firma nuevamente con XAdES
- **Estado actual Go:** ❌ **NO SE GENERA**

---

## 🔄 Flujo Completo Correcto

```
1. Crear Invoice XML (sin firmar)
   └─> FE-SETP990000001.xml

2. Firmar con XAdES
   └─> FES-SETP990000001.xml

3. Comprimir a ZIP
   └─> FES-SETP990000001.zip

4. Crear SOAP Request con WS-Security
   └─> ReqFE-SETP990000001.xml

5. Enviar a DIAN (SendBillSync o SendBillAsync)
   └─> Recibe respuesta

6. Parsear ApplicationResponse de DIAN
   └─> RptaFE-SETP990000001.xml
   └─> Extraer: ResponseCode, CUDE, fecha/hora validación

7. ⚠️ PASO CRÍTICO FALTANTE EN GO:
   Crear AttachedDocument con:
   - Factura firmada completa en CDATA
   - ApplicationResponse en CDATA
   - Metadatos (CUFE, fechas, partes)
   └─> AttachmentDocument-SETP990000001.xml (sin firmar)

8. Firmar AttachedDocument con XAdES
   └─> AttachmentDocument-SETP990000001.xml (firmado)

9. Generar PDF con QR
   └─> FES-SETP990000001.pdf
```

---

## 📄 Estructura del AttachedDocument

```xml
<AttachedDocument xmlns="urn:oasis:names:specification:ubl:schema:xsd:AttachedDocument-2">
  <ext:UBLExtensions>
    <!-- Firma XAdES del AttachedDocument -->
  </ext:UBLExtensions>
  
  <cbc:UBLVersionID>UBL 2.1</cbc:UBLVersionID>
  <cbc:CustomizationID schemeName="31">Documentos adjuntos</cbc:CustomizationID>
  <cbc:ProfileID>Factura Electrónica de Venta</cbc:ProfileID>
  <cbc:ProfileExecutionID>1</cbc:ProfileExecutionID>
  <cbc:ID>229944253</cbc:ID> <!-- ID único del AttachedDocument -->
  <cbc:IssueDate>2025-12-14</cbc:IssueDate>
  <cbc:IssueTime>08:17:34-05:00</cbc:IssueTime>
  <cbc:DocumentType>Contenedor de Factura Electrónica</cbc:DocumentType>
  <cbc:ParentDocumentID>BEC496329154</cbc:ParentDocumentID> <!-- ID de la factura -->
  
  <!-- Emisor -->
  <cac:SenderParty>
    <cac:PartyTaxScheme>
      <cbc:RegistrationName>COLOMBIA TELECOMUNICACIONES S.A. E.S.P. BIC</cbc:RegistrationName>
      <cbc:CompanyID>830122566</cbc:CompanyID>
      <cbc:TaxLevelCode>O-13;O-15;O-23</cbc:TaxLevelCode>
      <cac:TaxScheme>
        <cbc:ID>01</cbc:ID>
        <cbc:Name>IVA</cbc:Name>
      </cac:TaxScheme>
    </cac:PartyTaxScheme>
  </cac:SenderParty>
  
  <!-- Receptor -->
  <cac:ReceiverParty>
    <cac:PartyTaxScheme>
      <cbc:RegistrationName>DIEGO FERNANDO MONTOYA</cbc:RegistrationName>
      <cbc:CompanyID>6382356</cbc:CompanyID>
      <cbc:TaxLevelCode>49</cbc:TaxLevelCode>
      <cac:TaxScheme>
        <cbc:ID>01</cbc:ID>
        <cbc:Name>IVA</cbc:Name>
      </cac:TaxScheme>
    </cac:PartyTaxScheme>
  </cac:ReceiverParty>
  
  <!-- PRIMER ATTACHMENT: FACTURA FIRMADA COMPLETA -->
  <cac:Attachment>
    <cac:ExternalReference>
      <cbc:MimeCode>text/xml</cbc:MimeCode>
      <cbc:EncodingCode>UTF-8</cbc:EncodingCode>
      <cbc:Description><![CDATA[
        <?xml version="1.0" encoding="UTF-8"?>
        <Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2">
          <!-- FACTURA FIRMADA COMPLETA AQUÍ -->
        </Invoice>
      ]]></cbc:Description>
    </cac:ExternalReference>
  </cac:Attachment>
  
  <!-- SEGUNDO ATTACHMENT: APPLICATION RESPONSE DE DIAN -->
  <cac:ParentDocumentLineReference>
    <cbc:LineID>1</cbc:LineID>
    <cac:DocumentReference>
      <cbc:ID>BEC496329154</cbc:ID>
      <cbc:UUID schemeName="CUFE-SHA384">ffff0d032c292b88b3f839f75a51e8459ab645eda8049b3c221649fd18aaea09d5b31c8787e071c6a7d4db6983faaead</cbc:UUID>
      <cbc:IssueDate>2025-12-14</cbc:IssueDate>
      <cbc:DocumentType>ApplicationResponse</cbc:DocumentType>
      
      <cac:Attachment>
        <cac:ExternalReference>
          <cbc:MimeCode>text/xml</cbc:MimeCode>
          <cbc:EncodingCode>UTF-8</cbc:EncodingCode>
          <cbc:Description><![CDATA[
            <?xml version="1.0" encoding="utf-8"?>
            <ApplicationResponse xmlns="urn:oasis:names:specification:ubl:schema:xsd:ApplicationResponse-2">
              <!-- APPLICATION RESPONSE DE DIAN AQUÍ -->
              <cac:DocumentResponse>
                <cac:Response>
                  <cbc:ResponseCode>02</cbc:ResponseCode>
                  <cbc:Description>Documento validado por la DIAN</cbc:Description>
                </cac:Response>
              </cac:DocumentResponse>
            </ApplicationResponse>
          ]]></cbc:Description>
        </cac:ExternalReference>
      </cac:Attachment>
      
      <cac:ResultOfVerification>
        <cbc:ValidatorID>Unidad Especial Dirección de Impuestos y Aduanas Nacionales</cbc:ValidatorID>
        <cbc:ValidationResultCode>02</cbc:ValidationResultCode>
        <cbc:ValidationDate>2025-12-14</cbc:ValidationDate>
        <cbc:ValidationTime>08:17:34-05:00</cbc:ValidationTime>
      </cac:ResultOfVerification>
    </cac:DocumentReference>
  </cac:ParentDocumentLineReference>
</AttachedDocument>
```

---

## 🔑 Código PHP que lo Genera

### Ubicación en co-apidian2021:
- **Template:** `resources/templates/xml/89.blade.php`
- **Controller:** `app/Http/Controllers/Api/InvoiceController.php` (líneas 298-306, 372-380)
- **Trait:** `app/Traits/DocumentTrait.php` (método `createXML`)
- **Firma:** `ubl21dian/src/XAdES/SignAttachedDocument.php`

### Flujo en PHP:
```php
// 1. Recibir respuesta de DIAN
$appresponsexml = $SendBillSync->signToSend($request->GuardarEn."\\ReqFE-{$resolution->next_consecutive}.xml")->getResponseToDocument(storage_path("app/public/{$company->identification_number}/ReqFE-{$resolution->next_consecutive}.xml"), storage_path("app/public/{$company->identification_number}/RptaFE-{$resolution->next_consecutive}.xml"));

// 2. Parsear respuesta
$ar = new DOMDocument();
$ar->loadXML($appresponsexml);
$fechavalidacion = $ar->getElementsByTagName('IssueDate')->item(0)->nodeValue;
$horavalidacion = $ar->getElementsByTagName('IssueTime')->item(0)->nodeValue;

// 3. Crear AttachedDocument (template 89.blade.php)
$attacheddocument = $this->createXML(compact(
    'user', 'company', 'customer', 'resolution', 'typeDocument', 
    'cufecude', 'signedxml', 'appresponsexml', 
    'fechavalidacion', 'horavalidacion', 'document_number'
));

// 4. Firmar AttachedDocument
$signAttachedDocument = new SignAttachedDocument($company->certificate->path, $company->certificate->password);
$signAttachedDocument->GuardarEn = storage_path("app/public/{$company->identification_number}/{$filename}.xml");
$at = $signAttachedDocument->sign($attacheddocument)->xml;

// 5. Guardar
$file = fopen(storage_path("app/public/{$company->identification_number}/{$filename}.xml"), "w");
fwrite($file, $at);
fclose($file);
```

---

## ❌ Lo que FALTA en ubl21-dian (Go)

### Archivos/Funciones que NO existen:

1. **Template AttachedDocument**
   - No existe generador de XML AttachedDocument
   - Necesita crear estructura UBL 2.1 AttachedDocument

2. **Función para incrustar CDATA**
   - Debe incrustar factura firmada completa en CDATA
   - Debe incrustar ApplicationResponse en CDATA
   - Go necesita escapar correctamente el XML dentro de CDATA

3. **Firma de AttachedDocument**
   - Necesita firmar el AttachedDocument con XAdES
   - Reutilizar lógica de firma existente

4. **Integración en el flujo**
   - Después de recibir respuesta exitosa de DIAN
   - Antes de generar el PDF

---

## 📝 Plan de Implementación para Go

### Paso 1: Crear estructura AttachedDocument
```go
// invoice/attached_document.go
type AttachedDocument struct {
    UBLVersionID        string
    CustomizationID     string
    ProfileID           string
    ProfileExecutionID  string
    ID                  string
    IssueDate           string
    IssueTime           string
    DocumentType        string
    ParentDocumentID    string
    SenderParty         Party
    ReceiverParty       Party
    SignedInvoiceXML    string // XML firmado completo
    ApplicationResponse string // Response de DIAN
    ValidationDate      string
    ValidationTime      string
    CUFE                string
}
```

### Paso 2: Crear generador XML
```go
func (ad *AttachedDocument) ToXML() (string, error) {
    // Usar template o construcción manual
    // Incrustar SignedInvoiceXML y ApplicationResponse en CDATA
}
```

### Paso 3: Integrar en flujo SOAP
```go
// Después de enviar y recibir respuesta exitosa
if responseCode == "02" { // Validado
    // Crear AttachedDocument
    attached := &AttachedDocument{
        SignedInvoiceXML:    signedXML,
        ApplicationResponse: appResponse,
        ValidationDate:      validationDate,
        ValidationTime:      validationTime,
        CUFE:                cufe,
        // ... otros campos
    }
    
    // Generar XML
    attachedXML, err := attached.ToXML()
    
    // Firmar
    signedAttached, err := SignXML(attachedXML, cert)
    
    // Guardar
    os.WriteFile("AttachmentDocument-"+invoiceID+".xml", []byte(signedAttached), 0644)
}
```

---

## 🎯 Conclusión

**El AttachedDocument es el "contenedor final" que:**
1. Agrupa la factura firmada + respuesta de DIAN
2. Sirve como evidencia de validación
3. Se firma nuevamente para garantizar integridad
4. Es el documento que se archiva y puede ser consultado

**Sin este documento, el proceso está incompleto y puede causar:**
- ❌ Rechazo en auditorías DIAN
- ❌ Problemas de trazabilidad
- ❌ Falta de evidencia de validación
- ❌ Incumplimiento normativo

**Prioridad:** 🔴 **CRÍTICA** - Debe implementarse antes de producción
