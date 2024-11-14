package scanner

import (
	"fmt"
	"web-vulnerability-scanner/internal/headers"
	"web-vulnerability-scanner/internal/ssl"
)

// ScanWithProgress ejecuta el escaneo en una URL y muestra el progreso
func ScanWithProgress(url string) []string {
	var results []string

	// Progreso: Escaneo de encabezados HTTP
	fmt.Println("Escaneando encabezados de seguridad...")
	headerResults := headers.AnalyzeHeaders(url)
	results = append(results, headerResults...)

	// Progreso: Escaneo de configuración SSL/TLS
	fmt.Println("Escaneando configuración SSL/TLS...")
	sslResults := ssl.AnalyzeSSL(url)
	results = append(results, sslResults...)

	fmt.Println("Escaneo completado.")
	return results
}
