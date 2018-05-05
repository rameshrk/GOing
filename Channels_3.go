package main

import "fmt"

func main() {

	message1 := make(chan string,1 )
	message1 <- "a"
	fmt.Println(<-message1)


}
