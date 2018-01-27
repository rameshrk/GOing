package main

import "fmt"

func main()  {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)


	g := 0
	fmt.Println(g+10)
	//fmt.Println(g+"10")

	h := "0"
	fmt.Println(h+"10")
	//fmt.Println(h+10)
}