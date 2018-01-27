package main

import "fmt"

func main() {
	fmt.Println(val(1,2))
	fmt.Println(val(1,2))
	c, d := val(5,6)
	fmt.Println(c,d)
	_,e := val(9,10)
	fmt.Println(e)
}



func val(a int , b int ) (int, int) {
	return a, b
}
