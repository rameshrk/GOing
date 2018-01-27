package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write  ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")

	default:
		fmt.Println("weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("noon")
	default:
		fmt.Println("afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("unknow type", t)

		}
	}
	//whatAmI1 := func(i interface{}) {
	//	switch i.(type) {
	//	case bool:
	//		return "bool"
	//	case int:
	//		return "int"
	//	default:
	//		return "string"
	//
	//	}
	//}

	whatAmI2 := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("unknow type")

		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	whatAmI2(true)
	whatAmI2(1)
	whatAmI2("hey")

}
