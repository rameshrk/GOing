
package main

import (
"fmt"
"math"
)

var c, python, java bool
var i, j int = 1,2

func main() {
	var java1=true
	var c, python,jk = true , false, "no"
	fmt.Println(math.Pi)
	fmt.Println(add(1,1))
	fmt.Println(subtract(11,2))
	fmt.Println(swap(11,2))
	fmt.Println(split(7))
	fmt.Println(i, j, jk,c, python, java, j)
	fmt.Println(c, python, java1 )
	main2()
}

func add (x int, y int) int {
	return x+y
}

func subtract (x , y int) int {
	return x-y
}

func swap (x , y int) (int, int) {
	return y, x
}

func split(sum int) (x,y int) {
	x= sum+4
	y = sum -x
	return x, y
}


func main2 () {
	var i, j int =1, 2
	k := 4
	c, python,jk := true , false, "no"
	fmt.Println(i, j, k, c, python, jk)

}