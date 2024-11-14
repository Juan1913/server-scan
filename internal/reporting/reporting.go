package reporting

import "fmt"

// PrintResults muestra los resultados de una manera estructurada
func PrintResults(results []string) {
	fmt.Println("\nResultados del Escaneo:")
	for _, result := range results {
		fmt.Println("- ", result)
	}
}
