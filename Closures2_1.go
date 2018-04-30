package main

import "fmt"

func main() {

	a := fib()
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())
	fmt.Println("---", a())

}

func fib() func() int {
	sum := 1
	n := 0

	return func() int {
		//sum =  n + sum
		//fmt.Println("s ", sum)
		//n = sum
		//fmt.Println("n ", n)
		sum = (sum+n)
		n = sum
		return n
	}
}
