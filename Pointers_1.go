package main

import "fmt"

func zeroval(ival int){
	ival =0
}

func zeroptr(ival *int){
	*ival = 10
}

func main() {
	i := 0

	zeroval(i)
	fmt.Println("", i)

	zeroptr(&i)
	fmt.Println("", i)

}