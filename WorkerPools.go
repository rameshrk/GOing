package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("\t -- worker1", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("\t -- worker1", id, "finished job", j)
		results <- j * 2
	}
}

func workernext(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println(">> worker2", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println(">> worker2", id, "finiahed job", j)
		results <- j * 3
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
		go workernext(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

}