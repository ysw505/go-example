package main

import (
	"github.com/gopherjs/gopherjs/compiler/natives/src/sync"
	"fmt"
)



// fan in fan out
func main() {
	jobCount := 100
	jobQueue := make(chan string, jobCount)

	// make 100 job
	for i := 1; i <= jobCount; i++ {
		jobQueue <- fmt.Sprintf("job-%d", i)
	}

	close(jobQueue)

	wg := sync.WaitGroup{}
	wg.Add(5)

	workerCount := 5

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

}

func worker(idx int, wg *sync.WaitGroup, job  <-chan string, result chan <- string) {
	for j := range job {
		workerLog := fmt.Sprintf("Worker %d", idx)
		result <- workerLog + " " + j
	}
	wg.Add(-1)
}
