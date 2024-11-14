package ssl

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"
)

// AnalyzeSSL revisa las configuraciones de seguridad SSL/TLS
func AnalyzeSSL(rawURL string) []string {
	var results []string

	// Parsear la URL para extraer solo el host
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		results = append(results, fmt.Sprintf("Error al analizar la URL %s: %v", rawURL, err))
		return results
	}

	// Extraer el host (dominio) sin el prefijo `https://`
	host := parsedURL.Host
	if !strings.Contains(host, ":") {
		host = host + ":443" // Asegurar el puerto 443 para HTTPS
	}

	// Configura una conexión TLS
	conn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		results = append(results, fmt.Sprintf("Error al conectar con %s: %v", host, err))
		return results
	}
	defer conn.Close()

	// Verifica la versión de TLS
	switch conn.ConnectionState().Version {
	case tls.VersionSSL30:
		results = append(results, "SSL 3.0 está desactualizado e inseguro")
	case tls.VersionTLS10:
		results = append(results, "TLS 1.0 está desactualizado e inseguro")
	case tls.VersionTLS11:
		results = append(results, "TLS 1.1 está desactualizado e inseguro")
	default:
		results = append(results, "La versión TLS es segura")
	}

	return results
}
