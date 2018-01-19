package main

import (
	"sync"
	"fmt"
	"time"
	"runtime"
)

// fan in fan out
func main() {

	starTime := time.Now()

	jobCount := 10
	jobQueue := make(chan string, jobCount)

	// make 100 job
	for i := 1; i <= jobCount; i++ {
		jobQueue <- fmt.Sprintf("job-%d", i)
	}

	close(jobQueue)

	runtime.GOMAXPROCS(4)

	workerCount := 10
	wg := sync.WaitGroup{}

	wg.Add(workerCount)

	result := make(chan string, jobCount)


	for i := 0; i < workerCount; i++ {
		go worker(i, &wg, jobQueue, result)
	}

	fmt.Println("Wait group")
	wg.Wait()
	close(result)

	for r := range result {

		fmt.Println(r)
	}

	duration := time.Since(starTime)

	fmt.Println(duration)

}

func worker(idx int, wg *sync.WaitGroup, job <-chan string, result chan<- string) {
	for j := range job {
		workerLog := fmt.Sprintf("Worker %d", idx)
		time.Sleep(time.Second * 1)
		result <- workerLog + " " + j
	}
	wg.Add(-1)
}
