package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	ID      int
	Payload int
}

type Result struct {
	JobID int
	Value int
}

func WorkerPool(jobs []Job, workers int) []Result {
	results := make([]Result, 0, len(jobs))

	var writeWg sync.WaitGroup
	var readWg sync.WaitGroup

	jobCh := make(chan Job)
	resCh := make(chan Result)

	if workers < 1 {
		workers = 1
	}

	for range workers {
		writeWg.Go(func() {
			for job := range jobCh {
				resCh <- Result{JobID: job.ID, Value: job.Payload * 2}
			}
		})
	}

	readWg.Go(func() {
		for res := range resCh {
			results = append(results, res)
		}
	})

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	writeWg.Wait()
	close(resCh)

	readWg.Wait()

	return results
}

func main() {
	jobs := make([]Job, 20)
	for i := range jobs {
		jobs[i] = Job{i + 1, rand.Intn(20)}
	}

	rs := WorkerPool(jobs, 3)
	for _, r := range rs {
		fmt.Println(r.JobID, r.Value)
	}
}
