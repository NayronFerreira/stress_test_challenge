package models

type ReportStats struct {
	TotalRequests      int
	SuccessCount       int
	ErrorCount         int
	MinDuration        float64
	MaxDuration        float64
	StatusDistribution map[int]int
}

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
