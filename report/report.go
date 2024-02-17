package report

import (
	"fmt"

	"github.com/NayronFerreira/stress_test_challenge/loadtester"
)

func GenerateReport(totalResult loadtester.TotalResult) {
	totalRequests := len(totalResult.Results)
	successCount := 0
	errorCount := 0
	statusDistribution := make(map[int]int)
	minDuration := 500.0 // Isso pode ser problemático se todos os tempos forem maiores que 500
	var maxDuration float64

	for _, result := range totalResult.Results {
		if result.Error {
			errorCount++
		} else {
			if result.StatusCode == 200 {
				successCount++
			}
		}

		if result.Duration < minDuration || minDuration == 500.0 {
			minDuration = result.Duration
		}

		if result.Duration > maxDuration {
			maxDuration = result.Duration
		}

		statusDistribution[result.StatusCode]++
	}

	fmt.Println("=========================================")
	fmt.Println("           RELATÓRIO DE TESTE DE CARGA")
	fmt.Println("=========================================")
	fmt.Printf("URL Testada: %s\n", totalResult.URL) // Supondo que URL seja parte de TotalResult
	fmt.Printf("Total de Requisições: %d\n", totalRequests)
	fmt.Printf("Requisições Bem-Sucedidas (200 OK): %d\n", successCount)
	fmt.Printf("Requisições com Erros: %d\n", errorCount)
	fmt.Printf("Duração Total do Teste: %.2f ms\n", totalResult.TotalDuration)
	fmt.Printf("Duração Média por Requisição: %.2f ms\n", totalResult.TotalDuration/float64(totalRequests))
	fmt.Printf("Duração Mínima da Requisição: %.2f ms\n", minDuration)
	fmt.Printf("Duração Máxima da Requisição: %.2f ms\n", maxDuration)
	fmt.Println("\nDistribuição dos Códigos de Status HTTP:")
	for status, count := range statusDistribution {
		fmt.Printf("  %d: %d\n", status, count)
	}
	fmt.Println("=========================================")
}
