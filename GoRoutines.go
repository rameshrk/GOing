package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}

}
func main() {
	f("direct")
	go f("go thread")

	go func(msg string) {
		fmt.Println(msg)
		f("anonymous thrard inside")
	}("anonymous thread")

	fmt.Scanln()
	fmt.Println("done" )
}
