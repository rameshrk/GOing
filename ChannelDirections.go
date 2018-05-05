package main

import "fmt"

func ping(in chan<- string, msg string){
	in <- msg
}

func pong (in chan<- string, out <-chan string ){
	msg := <- out
	in <- msg
}

func main() {
	ins := make(chan string,1)
	outs := make(chan string,1)
	ping (ins,"message 1 ")
	pong (outs, ins)
	fmt.Println(<-outs)
}