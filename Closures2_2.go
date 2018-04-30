package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {
	// returns an anonymous function
	return func(message string) {
		fmt.Println(message)
	}
}

func outer(name string) {

}

func main() {
	// named function
	printMessage("Hello function!")

	// anonymous function declared and called
	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	func(message string) {
		fmt.Println(message)
	}("abc")

	output := func(message string) string {
		fmt.Println(message)
		return "from return"
	}("to print from argument")

	fmt.Println(output)

	output1 := func(message string) string {
		fmt.Println(message)
		return "from return output1"
	}

	fmt.Println(output1("output1"))

	// gets anonymous function and calls it
	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller!")

}
