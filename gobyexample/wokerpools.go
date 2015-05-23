package main

import (
	"fmt"
	"time"
)

//this worker will receive work on the job channel
//and send the corresponding results on 'results'
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		//simulate an expensive work
		time.Sleep(time.Second)
		results <- j * 10
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 1)

	for w := 1; w <= 3; w++ {
		//the goroutines are initially blocked
		//because the results have not received the value
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	//when we sent the message to results at this step
	//the worker goroutines could be activited
	for a := 1; a <= 9; a++ {
		<-results
	}

}
