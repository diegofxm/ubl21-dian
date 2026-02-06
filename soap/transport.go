package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Transport maneja el transporte HTTP/HTTPS con mTLS
type Transport struct {
	httpClient *http.Client
	url        string
	debugDir   string
}

// NewTransport crea un nuevo transport SOAP
func NewTransport(url string, tlsConfig *tls.Config, timeout time.Duration) *Transport {
	return &Transport{
		httpClient: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
		url:      url,
		debugDir: "storage/logs/debug/soap",
	}
}

// Send envía un request SOAP y retorna la respuesta
func (t *Transport) Send(soapXML string) ([]byte, error) {
	// Guardar request para debugging
	timestamp := time.Now().Format("20060102_150405")
	os.MkdirAll(t.debugDir, 0755)
	
	debugRequestPath := fmt.Sprintf("%s/soap_request_%s.xml", t.debugDir, timestamp)
	if err := os.WriteFile(debugRequestPath, []byte(soapXML), 0644); err != nil {
		fmt.Printf("Warning: Failed to save debug request: %v\n", err)
	}

	// Crear request HTTP
	req, err := http.NewRequest("POST", t.url, bytes.NewBufferString(soapXML))
	if err != nil {
		return nil, NewSOAPError("Transport", ErrHTTPTransport, "failed to create HTTP request", err)
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(soapXML)))

	// Enviar request
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, NewSOAPError("Transport", ErrHTTPTransport, "failed to send HTTP request", err)
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewSOAPError("Transport", ErrHTTPTransport, "failed to read response body", err)
	}

	// Verificar status code
	if resp.StatusCode != http.StatusOK {
		return nil, NewSOAPError("Transport", ErrHTTPTransport, 
			fmt.Sprintf("HTTP error %d: %s", resp.StatusCode, string(body)), nil)
	}

	// Guardar response para debugging
	debugResponsePath := fmt.Sprintf("%s/soap_response_%s.xml", t.debugDir, timestamp)
	if err := os.WriteFile(debugResponsePath, body, 0644); err != nil {
		fmt.Printf("Warning: Failed to save debug response: %v\n", err)
	}

	return body, nil
}

// LoadClientTLSConfig carga el certificado y clave privada para mTLS
func LoadClientTLSConfig(certPath, keyPath string) (*tls.Config, error) {
	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, NewSOAPError("TLS", ErrCertificateLoad, "failed to read certificate", err)
	}

	var keyPEM []byte
	if keyPath != "" && keyPath != certPath {
		keyPEM, err = os.ReadFile(keyPath)
		if err != nil {
			return nil, NewSOAPError("TLS", ErrCertificateLoad, "failed to read private key", err)
		}
	} else {
		keyPEM = certPEM
	}

	// Extraer certificado y clave privada en un solo recorrido
	var certDER []byte
	var keyDER []byte
	
	// Primero buscar en certPEM (que puede contener ambos)
	pemData := certPEM
	for {
		block, rest := pem.Decode(pemData)
		if block == nil {
			break
		}
		
		if block.Type == "CERTIFICATE" && certDER == nil {
			certDER = block.Bytes
		} else if (block.Type == "PRIVATE KEY" || block.Type == "RSA PRIVATE KEY") && keyDER == nil {
			keyDER = block.Bytes
		}
		
		pemData = rest
		
		// Si ya encontramos ambos, salir
		if certDER != nil && keyDER != nil {
			break
		}
	}
	
	// Si no encontramos la clave en certPEM, buscar en keyPEM
	if keyDER == nil && keyPEM != nil && len(keyPEM) > 0 {
		pemData = keyPEM
		for {
			block, rest := pem.Decode(pemData)
			if block == nil {
				break
			}
			if block.Type == "PRIVATE KEY" || block.Type == "RSA PRIVATE KEY" {
				keyDER = block.Bytes
				break
			}
			pemData = rest
		}
	}

	if certDER == nil || keyDER == nil {
		return nil, NewSOAPError("TLS", ErrCertificateLoad, "failed to decode certificate or key", nil)
	}

	cert, err := tls.X509KeyPair(
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}),
		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}),
	)
	if err != nil {
		return nil, NewSOAPError("TLS", ErrCertificateLoad, "failed to create X509 key pair", err)
	}

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}, nil
}
