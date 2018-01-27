package main

import "fmt"

func main() {
	res := plus(1,2)
	fmt.Println("1+2 =", res)
	res = plusplus(1,2,3)
	fmt.Println(res)
}


func plus(a int, b int) int {
	return a+b
}

func plusplus (a,b,c int) int {
	return a+b+c
}

