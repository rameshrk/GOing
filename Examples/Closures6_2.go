package main

import "fmt"

func main() {
	n := 0
	counter := func() int {
		n += 1
		return n
	}

	fmt.Println(counter())
	n = 0
	fmt.Println(counter())
}
