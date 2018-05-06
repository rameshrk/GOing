package main

import "fmt"

func sum(s []int, c chan int, tag string) {
	sum := 0
	for _, v := range s {
		fmt.Println("", v, tag)
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5,6}
	c := make(chan int)
	go sum(s[:len(s)/2], c, "a")
	fmt.Println("", )
	go sum(s[len(s)/2:], c, "b")
	x, y := <-c, <-c
	//c <- 1
	fmt.Println("", x, y)
}
