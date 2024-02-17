package loadtester

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type TotalResult struct {
	URL           string
	Results       []Result
	TotalDuration float64
}

type Result struct {
	StatusCode   int
	Duration     float64
	Error        bool
	ErrorMessage string
}

func RunLoadTest(url string, totalRequests, concurrency int) TotalResult {
	results := make([]Result, 0, totalRequests)
	resultChan := make(chan Result, totalRequests)

	var wg sync.WaitGroup

	// Definir um timeout para cada requisição
	requestTimeout := 10 * time.Second

	// Criar um client HTTP com timeout
	client := &http.Client{
		Timeout: requestTimeout,
	}

	totalStartTime := time.Now()

	// Usar um semáforo para limitar a concorrência
	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{} // Adquire

			// Criar um contexto com timeout para esta requisição
			ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
			defer cancel() // Importante para evitar vazamentos de recursos

			startTime := time.Now()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				resultChan <- Result{Error: true}
				<-semaphore // Libera
				return
			}

			resp, err := client.Do(req)

			duration := time.Since(startTime).Milliseconds()

			result := Result{
				Duration: float64(duration),
			}

			if err != nil || resp.StatusCode != http.StatusOK {
				result.StatusCode = resp.StatusCode
				result.Error = true
				if err != nil {
					result.ErrorMessage = err.Error()
				} else {
					resBody, _ := io.ReadAll(resp.Body)
					bodyMessage := string(resBody)
					result.ErrorMessage = bodyMessage
					resp.Body.Close()
				}

			} else {
				result.StatusCode = resp.StatusCode
				resp.Body.Close()
			}

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

	return TotalResult{URL: url, Results: results, TotalDuration: float64(totalDuration)}
}
