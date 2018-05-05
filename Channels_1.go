package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println("")
	fmt.Println("abc", msg)

	go func() {
		messages <- "ping1"
	}()
	msg = <-messages
	fmt.Println("", msg)

	go func() {
		messages <- "ping2"
	}()
	msg = <-messages
	fmt.Println("", msg)

	go func() {
		messages <- "ping3"
	}()
	msg = <-messages
	fmt.Println("", msg)

	fmt.Println("done", )
}
