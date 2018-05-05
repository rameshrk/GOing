package main

import (
	"fmt"
	"time"
)

func worker(done chan int, i int ) {
	fmt.Println("Working", i)
	time.Sleep(time.Second)
	fmt.Println("Done", i)
	done <- i
}

func main() {
	num := 10
	messages := make(chan int, num)

	for i := 1; i <= num; i++ {
		fmt.Println("", i)
		go func() { worker(messages, i) }()
	}
	m := 0
	for i := 0; i < num; i++ {
		m = <-messages
		fmt.Println(m)
	}

}
