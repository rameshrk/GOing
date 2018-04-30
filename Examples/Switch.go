package main

import (
	"fmt"
	"time"
)
//import "time"

func main() {
	i :=2

	fmt.Println("Write", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3 :
		fmt.Println("three")

	}

	switch time.Now().Weekday(){
	case time.Sunday, time.Saturday:
		fmt.Println("weekend")
	case time.Monday, time.Tuesday :
		fmt.Println("Start")
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int :
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI('c')
	whatAmI("s")
}
