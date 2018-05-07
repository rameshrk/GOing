package main

import "fmt"

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("passed ",x)
			x, y = y, x+y
		case o := <-quit:
			
			fmt.Println("quit", o)
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print("\treceived ",<-c, "\n")

		}
		fmt.Println("\tdone")
		quit <- 1
	}()
	fibonacci(c, quit)
}
