package main

import "fmt"

func main() {
	res := plus(1,2)
	fmt.Println("1+2 =", res)
	res = plusplus(1,2,3)
	fmt.Println(res)
	fmt.Println(plusplusplus(1,2,3))
	a, b := vals ()
	fmt.Println(a,b)
	c,_ := vals ()
	fmt.Println(c)
}
func vals() (int,int) {
	return 3,7
}

func plus(a int, b int) int {
	return a+b
}

func plusplus (a,b,c int) int {
	return a+b+c
}


func plusplusplus (a,b,c int) int {
	return a+b+c
}



