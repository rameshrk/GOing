package main

import (
	"fmt"
	"sync"
	"time"
)

func messagechange() func(messages chan string, count int) chan string {

	return func(messages chan string, count int) chan string {
		time.Sleep(time.Millisecond * time.Duration(1000))
		//for i := 0; i < count; i++ {
		messages <- fmt.Sprintf("hello: %d", count)
		//}
		//msg := <- *messages
		//fmt.Println("", msg)
		return messages
	}
}

func sayhello(messages chan<- string, count int) {
	time.Sleep(time.Millisecond * time.Duration(1000))
	messages <- fmt.Sprintf("hello: %d", count)
}


func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)

	//go func() { messages <- "ping" }()

	forfunc := messagechange()
	//forfunc1 := messagechange()(&messages,10)
	//fmt.Println("", forfunc)
	//fmt.Println("", forfunc1)
	for x := 1; x <= 5; x++ {
	forfunc2 := sayhello()

	//go func() { forfunc(&messages, 10) }()
	forfunc(messages, 10)
	fmt.Println(messages)
	msg := <-messages
	fmt.Println(msg)

	go func() {
		wg.Wait()
		close(messages)
	}()
	fmt.Scanln()
	fmt.Println("done")
}
