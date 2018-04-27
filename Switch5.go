package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Println("write", i, "as ")
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

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("a bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("not know", t)
		}

	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")




}
