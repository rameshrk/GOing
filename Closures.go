package main

import "fmt"

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt)

	fmt.Println(fmt.Println(plus))
	fmt.Println(fmt.Println(plus))

	plus()
	plus()

	nextEven := closure2()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	nextEven1 := closure1()
	fmt.Println(nextEven1()) // 0
	fmt.Println(nextEven1()) // 2
	fmt.Println(nextEven1()) // 4

	outer("Hello")

	fooInMain := outerReturn("Hello___")
	fooInMain()
	outerReturn("abc")()

	ctr_, incr_ := counter(2)
	ctr1, incr1 := counter(12)
	fmt.Println(ctr_())
	fmt.Println(ctr1())
	fmt.Println(incr_)
	fmt.Println(incr1)

	incr_()
	incr_()
	incr1()
	fmt.Println(ctr_())
	fmt.Println(ctr1())

	ctr2, incr2 := counter(20)
	fmt.Println(ctr2())
	fmt.Println(incr2)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}

}

func plus1() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("--", increment())
	fmt.Println("--", increment())

}
func closure1() func() uint {
	i := uint(0)
	fmt.Println("closure1", i)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func closure2() func() uint {
	i := uint(0)
	fmt.Println("closure2", i)
	return func() (ret uint) {
		ret = 2
		i += 2
		return

	}
}

func outer(name string) {
	text := "Modified" + name
	foo := func() {
		fmt.Println(text)
	}
	foo()
}

func outerReturn(name string) func() {
	text := "Modified" + name

	foo := func() {
		fmt.Println(text)
	}
	return foo
}

func counter(start int) (func() int, func()) {
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	return ctr, incr

}
