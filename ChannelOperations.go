package main

import "fmt"

func main() {
	//messages := make(chan string)
	messages := make(chan string,1)
	signals := make(chan bool)

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("nothing to send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received", )
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received", )
	}




	msg = "hi_2"
	select {
	case messages <- msg:
		fmt.Println("sent message to messages ", msg)
	case sig := <- signals:
		fmt.Println("sent message to sig", sig)
	default:
		fmt.Println("nothing to send")
	}

}
