package soap

// Namespaces XML
const (
	NSAddressing        = "http://www.w3.org/2005/08/addressing"
	NSSOAPEnvelope      = "http://www.w3.org/2003/05/soap-envelope"
	NSDIANColombia      = "http://wcf.dian.colombia"
	NSXMLDSig           = "http://www.w3.org/2000/09/xmldsig#"
	NSWSSecurity        = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	NSWSSecurityUtility = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	NSExcC14N           = "http://www.w3.org/2001/10/xml-exc-c14n#"
	NSRSASHA256         = "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"
	NSSHA256            = "http://www.w3.org/2001/04/xmlenc#sha256"
	NSX509V3            = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-x509-token-profile-1.0#X509v3"
	NSBase64Binary      = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
)

// SOAP Actions - 15 operaciones DIAN
const (
	// Grupo 1: Envío de Documentos
	ActionSendBillSync            = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendBillSync"
	ActionSendBillAsync           = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendBillAsync"
	ActionSendTestSetAsync        = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendTestSetAsync"
	ActionSendBillAttachmentAsync = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendBillAttachmentAsync"
	ActionSendNominaSync          = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendNominaSync"

	// Grupo 2: Consulta de Estado
	ActionGetStatus      = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetStatus"
	ActionGetStatusZip   = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetStatusZip"
	ActionGetStatusEvent = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetStatusEvent"

	// Grupo 3: Eventos de Documentos
	ActionSendEventUpdateStatus = "http://wcf.dian.colombia/IWcfDianCustomerServices/SendEventUpdateStatus"

	// Grupo 4: Consultas de Información
	ActionGetNumberingRange   = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetNumberingRange"
	ActionGetXmlByDocumentKey = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetXmlByDocumentKey"
	ActionGetReferenceNotes   = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetReferenceNotes"
	ActionGetDocumentInfo     = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetDocumentInfo"
	ActionGetAcquirer         = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetAcquirer"
	ActionGetExchangeEmails   = "http://wcf.dian.colombia/IWcfDianCustomerServices/GetExchangeEmails"
)

// IDs para elementos de seguridad
const (
	IDBinarySecurityToken    = "TORRESOFTWARE"
	IDSecurityTokenReference = "STR"
	IDSignature              = "SIG"
	IDTimestamp              = "TS"
	IDKeyInfo                = "KI"
	IDTo                     = "ID"
)
