package main

import "fmt"

func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func fact(n int ) int {
	if n ==0 {
		return 1
	}
	return n* fact(n-1)
}

func main() {
	fmt.Println(fact(10))
}
