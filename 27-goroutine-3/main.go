package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job %d\n", id, job)
		results <- job * 2
	}
}

// workPool  goroutine
func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//开始三个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
}
