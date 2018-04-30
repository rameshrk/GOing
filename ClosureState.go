package main

import "fmt"

func counter(start int) (func() int, func()) {
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	return ctr, incr
}

func main() {
	c , i := counter(1)
	fmt.Println(c())
	i()
	fmt.Println(c())
	i()
	fmt.Println(c())

}