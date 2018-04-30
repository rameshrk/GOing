package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

var DoStuff func() = func() {
	// Do stuff
}



func getPrintMessage() func (string) {
	return func(i string) {
		fmt.Println(i)
	}
}

func main() {
	getPrintMessage()("abc")
}