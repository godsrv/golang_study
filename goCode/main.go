package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numJob = 3
	const numWorker = 3

	jobs := make(chan int, numJob)
	results := make(chan int, numJob)

	// 启动协程池
	var wg sync.WaitGroup
	for w := 1; w <= numWorker; w++ {
		wg.Add(1)
		go func(workerId int) {
			worker(workerId, jobs, results, &wg)
		}(w)
	}

	for j := 1; j <= numJob; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()

	close(results)

	for v := range results {
		fmt.Println("result:", v)
	}

	// // 从 results 通道读取结果
	// for a := 1; a <= numJob; a++ {
	// 	fmt.Println("Result:", <-results)
	// }

}

func worker(id int, jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)

		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}
