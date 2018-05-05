package main

import "fmt"

func messagechange1() func(messages *chan string, count int)  chan string {

	return func(messages *chan string, count int)  chan string {
		for i := 0; i < count; i++ {
			*messages <- "pong"

		}
		//msg := <- *messages
		//fmt.Println("", msg)
		return *messages
	}
}

func main() {
	messages := make(chan string)

	//go func() { messages <- "ping" }()

	forfunc := messagechange1()
	forfunc1 := messagechange1()(&messages,10)
	fmt.Println("", forfunc)
	fmt.Println("", forfunc1)

	//go func() { forfunc(&messages, 10) }()

	fmt.Println(messages)
	msg := <-messages
	fmt.Println(msg)
}
