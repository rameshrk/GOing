package main

import "fmt"

var DoStuff1 func() = func() {

}


func RefFunc() { fmt.Println("reg")}




func returnInt() func() int {
	return func() int{
		return 1
	}
}


func returnMessage() func(message string) string {
	return func(message string) string {
		fmt.Println("inside")
		return message
	}
}


func returnMessage1() func(string) string {
	return func(message string) string {
		fmt.Println("inside")
		return message
	}
}


func main() {
	DoStuff1()
	DoStuff1 = RefFunc
	DoStuff1()
	fmt.Println(returnInt())
	fmt.Println(returnInt()())
    result := returnInt()
	fmt.Println(result())

    fmt.Println(returnMessage()("abc"))
	fmt.Println(returnMessage1()("abc"))

}