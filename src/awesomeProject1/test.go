package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("" , sum)

//	for key, value := range oldMap {
//		newMap[key] = value
//	}
}