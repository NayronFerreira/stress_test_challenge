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

	requestTimeout := 10 * time.Second

	client := &http.Client{
		Timeout: requestTimeout,
	}

	var wg sync.WaitGroup

	totalStartTime := time.Now()

	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{} // Adquire

			ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
			defer cancel()

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
			if err != nil {
				result.Error = true
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
