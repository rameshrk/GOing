package main

import "fmt"

func seq() func() int {
	i := 0
	return func() int {
		i += 1
		return  i
	}
}

func seqArgument(number int) func() int {
	i := number
	return func() int {
		i += 1
		return  i
	}
}



func printC(word string)  func() {
	return func() {
		fmt.Println(word)
	}
}



func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}




func main() {

	sayHello := printC("Hello")
	sayHello() // prints Hello

	saySeq := seq()
	fmt.Println(saySeq())

	saySeqArg := seqArgument(20)
	fmt.Println(saySeqArg())
	fmt.Println("\n")

	fib := makeFibGen()
	fmt.Println(fib)

	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

	fmt.Println("\n")
	fmt.Println(makeFibGen()())
	fmt.Println(makeFibGen()())
	fmt.Println(makeFibGen()())
	fmt.Println(makeFibGen()())


	fmt.Println("\n")
	fib = makeFibGen()
	nextfib := fib()
	fmt.Println(nextfib)
	fmt.Println(nextfib)
	fmt.Println(nextfib)
	fmt.Println(nextfib)
	fmt.Println(nextfib)






}
