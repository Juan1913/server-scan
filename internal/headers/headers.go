package headers

import (
	"fmt"
	"net/http"
)

// AnalyzeHeaders realiza un análisis básico de encabezados de seguridad
func AnalyzeHeaders(url string) []string {
	var results []string

	resp, err := http.Head(url)
	if err != nil {
		results = append(results, fmt.Sprintf("Error al conectar con %s: %v", url, err))
		return results
	}
	defer resp.Body.Close()

	// Encabezados de seguridad recomendados
	requiredHeaders := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
		"X-Frame-Options",
	}

	// Verifica la presencia de cada encabezado
	for _, header := range requiredHeaders {
		if _, ok := resp.Header[header]; !ok {
			results = append(results, fmt.Sprintf("Falta el encabezado de seguridad: %s", header))
		}
	}

	if len(results) == 0 {
		results = append(results, "Todos los encabezados de seguridad están presentes.")
	}

	return results
}
