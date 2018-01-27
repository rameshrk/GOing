package main

import (
	"fmt"
)

func main() {

	functionarray := functions()
	fmt.Println(functionarray)

	for f:= range functionarray {
		functionarray[f]()
	}


}
func functions() []func() {
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func() { fmt.Printf("\nindex - %d, value - %d", i, arr[i]) } )
		index := i
		val := arr[i]
		result = append(result, func() { fmt.Printf("\nSecond time : index - %d, value - %d\n", index, val) })
		fmt.Println(result)

	}
	return result
}
