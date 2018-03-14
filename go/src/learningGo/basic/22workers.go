package main

import (
	"fmt"
	"time"
)

func workersDoingWork(id int, jobs <-chan int, results <-chan int) {
	for j:=range jobs {
		fmt.Println("worker", id, "started", j)
		time.Sleep(time.Second * 1)
		fmt.Println("worker", id, "finished", j)
		results <- j
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w:=1 ;w<=3;w++ {
		go workersDoingWork(w, jobs, results)
	}

	for j:=1;j<=5;j++ {
		jobs <- j
	}
	close(jobs)

	for a:=1 ;a<=5;a++ {
		<-results
	}
}