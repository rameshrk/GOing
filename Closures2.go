package main

import "fmt"

var DS = func() {
}

var DSS func() = func() {

}

func RegFunc() {
	fmt.Println("reg func")
}

var DS3 = func() { fmt.Println("simple") }

func main() {
	DS()

	DS = func() {
		fmt.Println("abc")
	}

	DSS = func() {
		fmt.Println("abc_")
	}

	DS()
	DSS()

	DS = func() {
		fmt.Println("efg")
	}
	DS = func() {
		fmt.Println("efg_")
	}

	DS()
	DSS()

	//DS2 = RegFunc  //undefined: DS2
	//DS2()

	DS3()
	DS3 = RegFunc
	DS3()

	DS4 := RegFunc
	DS4()

	n := 0
	counter := func() int {
		n += 1
		return n
	}

	fmt.Println(counter())
	n = 10
	fmt.Println(counter())

	newC := newCounter()
	fmt.Println(newC())
	fmt.Println(newC())

}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 2
		return n
	}

}
