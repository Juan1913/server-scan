package main

import (
	"bufio"
	"fmt"
	"os"
	"web-vulnerability-scanner/internal/reporting"
	"web-vulnerability-scanner/internal/scanner"
)

func main() {

	showBanner()

	url := getUserInput("Ingrese la URL a escanear: ")

	if url == "" {
		fmt.Println("URL no válida. Por favor, intente de nuevo.")
		return
	}

	fmt.Println("Iniciando escaneo...")

	results := scanner.ScanWithProgress(url)

	reporting.PrintResults(results)
}

func showBanner() {
	fmt.Println(`
    ========================================
    ||  WEB VULNERABILITY SCANNER V1.0    ||
    ||   Escanea encabezados y SSL/TLS    ||
    ========================================
    `)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
