package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	workerpool "github.com/troxanna/monitoring-service/workerpool"
)

var urls = []string{
	"https://www.google.com/",
	"https://golangify.com/",
	"https://habr.com/ru/feed/",
}

const (
	INTERVAL = 	time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT = 3 
)


func main() {
	results := make(chan workerpool.Result)
	pool := workerpool.New(WORKERS_COUNT, INTERVAL, results)

	pool.Init()

	go generateJobs(pool)
	go processResults(results)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<- quit

	pool.Stop()
}

func generateJobs(pool *workerpool.Pool) {
	for {
		for _, url := range urls {
			pool.Push(workerpool.Job{URL: url})
		}

		time.Sleep(INTERVAL)
	}
}

func processResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result.Info())
		}
	}()
}