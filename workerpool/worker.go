package workerpool

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func newWorker(timeout time.Duration) *worker {
	return &worker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (w *worker) process(job Job) Result {
	result := Result{URL: job.URL}
	now := time.Now()

	resp, err := w.client.Get(job.URL)
	if err != nil {
		result.Error = err

		return result
	}

	result.StatusCode = resp.StatusCode
	result.ResponseTime = time.Since(now)

	return result
}