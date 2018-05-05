package main

import "fmt"

func main() {

	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "buffered1"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages)

	//message1 := make(chan string)
	//message1 <- "a"
	//fmt.Println(<-message1)


}
