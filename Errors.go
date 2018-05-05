package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("error 42")
	}

	return arg + 3, nil

}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "f2 error 42 struct"}
	}
	return arg + 3, nil
}

func f3(arg int) (int, *argError) {
	if arg == 42 {
		return -1, &argError{arg, "f3 error 42 struct"}
	}
	return arg + 3, &argError{0, ""}
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d %s", e.arg, e.prob)

}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}

	}
	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	fmt.Println("", )

	_, e1 := f2(42)
	ae1, okk := e1.(*argError)
	fmt.Println(ae1)
	fmt.Println(okk)


	_, e2 := f3(42)
	ae2, okkk := e2.(*argError)
	fmt.Println(ae2)
	fmt.Println(okkk)

}
