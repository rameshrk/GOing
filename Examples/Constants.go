package main

import "fmt"
import "math"


const s string = "constant"

func main(){
	fmt.Println(s)
	const n = 50000000
	const d = 3e20 / 50000000
	fmt.Println(n)
	fmt.Println(d/n)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println(s)
	//s = "abc"
}