package signature

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ConvertP12ToPEMWithOpenSSL convierte un P12 a PEM usando OpenSSL
// Esto maneja certificados en formato BER que Go no puede leer directamente
// Incluye TODOS los certificados (usuario + intermedios + raíz) para firma de XML
func ConvertP12ToPEMWithOpenSSL(p12Path, password string) (string, error) {
	// Crear archivo temporal para el PEM
	dir := filepath.Dir(p12Path)
	pemPath := filepath.Join(dir, filepath.Base(p12Path)+".pem")

	// Verificar si ya existe el PEM y tiene contenido válido
	if pemData, err := os.ReadFile(pemPath); err == nil && len(pemData) > 0 {
		// PEM ya existe y tiene contenido, usarlo directamente
		return pemPath, nil
	}

	// Ejecutar OpenSSL para convertir P12 a PEM
	// Combina certificado y clave privada en un solo archivo
	// -legacy: Soporta algoritmos antiguos como RC2-40-CBC (certificados DIAN)
	cmd := exec.Command("openssl", "pkcs12",
		"-in", p12Path,
		"-out", pemPath,
		"-nodes", // No encriptar la clave privada
		"-legacy", // Soportar algoritmos legacy
		"-passin", "pass:"+password,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to convert P12 to PEM with OpenSSL: %w\nOutput: %s", err, string(output))
	}

	// Verificar que el archivo PEM se creó y tiene contenido
	pemData, err := os.ReadFile(pemPath)
	if err != nil {
		return "", fmt.Errorf("failed to read generated PEM file: %w", err)
	}

	if len(pemData) == 0 {
		return "", fmt.Errorf("generated PEM file is empty")
	}

	return pemPath, nil
}

// ConvertP12ToClientPEM convierte un P12 a PEM con SOLO el certificado de cliente
// (sin certificados CA intermedios ni raíz) para usar en SOAP security header
func ConvertP12ToClientPEM(p12Path, password string) (string, error) {
	dir := filepath.Dir(p12Path)
	clientPemPath := filepath.Join(dir, filepath.Base(p12Path)+".client.pem")

	// Verificar si ya existe
	if pemData, err := os.ReadFile(clientPemPath); err == nil && len(pemData) > 0 {
		return clientPemPath, nil
	}

	// Ejecutar OpenSSL con -clcerts para extraer SOLO certificado de cliente
	// -clcerts: Solo certificados de cliente (no CA certs)
	cmd := exec.Command("openssl", "pkcs12",
		"-in", p12Path,
		"-out", clientPemPath,
		"-nodes",
		"-legacy",
		"-clcerts", // IMPORTANTE: Solo certificado de cliente
		"-passin", "pass:"+password,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to convert P12 to client PEM: %w\nOutput: %s", err, string(output))
	}

	// Verificar contenido
	pemData, err := os.ReadFile(clientPemPath)
	if err != nil {
		return "", fmt.Errorf("failed to read client PEM: %w", err)
	}

	if len(pemData) == 0 {
		return "", fmt.Errorf("client PEM is empty")
	}

	return clientPemPath, nil
}

// NewSignerFromP12WithFallback intenta cargar P12 nativamente, si falla usa OpenSSL
func NewSignerFromP12WithFallback(p12Path, password string) (*Signer, error) {
	// Intento 1: Cargar P12 directamente con Go
	signer, err := NewSignerFromP12(p12Path, password)
	if err == nil {
		return signer, nil
	}

	// Intento 2: Convertir a PEM con OpenSSL y cargar
	pemPath, err := ConvertP12ToPEMWithOpenSSL(p12Path, password)
	if err != nil {
		return nil, fmt.Errorf("failed to load P12: native Go failed and OpenSSL conversion failed: %w", err)
	}

	// Cargar desde PEM
	signer, err = NewSignerFromSinglePEM(pemPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load converted PEM: %w", err)
	}

	return signer, nil
}
