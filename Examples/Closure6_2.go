package main

import "fmt"

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		fmt.Print("called me ")
		return n
	}
}



func main() {
	counter := newCounter
	fmt.Println("newCounter",newCounter)			//newCounter 0x4968c0
	fmt.Println("newCounter()",newCounter())		//newCounter() 0x496e80
	fmt.Println("counter",counter)				//counter 0x4968c0
	fmt.Println("counter()",counter())			//counter() 0x496e80
	fmt.Println()
	counter1 := newCounter()
	fmt.Println("newCounter",newCounter)			//newCounter 0x4968c0
	fmt.Println("newCounter()",newCounter())		//newCounter() 0x496e80
	fmt.Println("counter1",counter1)				//counter1 0x496e80
	fmt.Println("counter1()",counter1())			//called me counter1() 1
}
