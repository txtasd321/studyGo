package main

import (
	"fmt"
	"time"
)

// TODOworker pool(goroutine池)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}
	// 5个任务
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 20; a++ {
		<-result
	}
}
