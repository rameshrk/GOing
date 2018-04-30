package main


import "fmt"

func functions() []func() {
	// pitfall of using loop variables
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)
	result1 := make([]func(), 0)

	// functions are not evaluated, functions definitions are returned
	// since functions refer to i, arr[i]
	// the last values of i, arr[i] are always referred to
	for i := range arr {
		result = append(result, func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) })
	}

	// if desired behaviour is needed, one needs to use them as args
	for i := range arr {
		result1 = append(result, func(index int, val int) { fmt.Printf("index - %d, value - %d\n", index, val) })
	}

	return result
}

func main() {
	fns := functions()
	for f := range fns {
		fns[f]()
	}
}
