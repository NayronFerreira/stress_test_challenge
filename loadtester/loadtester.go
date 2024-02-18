package loadtester

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/NayronFerreira/stress_test_challenge/constants"
	"github.com/NayronFerreira/stress_test_challenge/models"
)

func RunLoadTest(url string, totalRequests, concurrency int) models.TotalResult {
	results := make([]models.Result, 0, totalRequests)
	resultChan := make(chan models.Result, totalRequests)

	client := &http.Client{
		Timeout: constants.DefaultRequestTimeout,
	}

	var wg sync.WaitGroup

	totalStartTime := time.Now()

	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{} // Adquire

			result := performRequest(client, url)
			resultChan <- result

			<-semaphore // Libera
		}()
	}

	wg.Wait()
	close(resultChan)

	totalDuration := time.Since(totalStartTime).Milliseconds()

	for result := range resultChan {
		results = append(results, result)
	}

	return models.TotalResult{URL: url, Results: results, TotalDuration: float64(totalDuration)}
}

func performRequest(client *http.Client, url string) models.Result {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
	defer cancel()

	startTime := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return models.Result{Error: true, ErrorMessage: err.Error()}
	}

	resp, err := client.Do(req)

	duration := time.Since(startTime).Milliseconds()

	result := models.Result{
		Duration: float64(duration),
	}
	if err != nil {
		result.Error = true
		result.StatusCode = http.StatusInternalServerError
		result.ErrorMessage = err.Error()
	} else {
		defer resp.Body.Close()
		result.StatusCode = resp.StatusCode
		result.Duration = float64(duration)

		if resp.StatusCode != http.StatusOK {
			result.Error = true
			resBody, err := io.ReadAll(resp.Body)
			if err == nil {
				result.ErrorMessage = string(resBody)
			} else {
				result.ErrorMessage = "Erro ao ler o corpo da resposta"
			}
		}
	}

	return result
}
