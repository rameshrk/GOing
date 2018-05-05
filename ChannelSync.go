package main

import (
	"fmt"
	"time"
)

func worker (done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool,1 )
	go func() {worker(done)}()
	go func() {worker(done)}()


	m := <-done
	fmt.Println("", m)

}