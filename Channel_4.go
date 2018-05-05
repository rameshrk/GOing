package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {messages <- "a"}()

	m := <- messages
	fmt.Println("", m)
}