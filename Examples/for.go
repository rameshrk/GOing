package main

import "fmt"

func main() {
	nums := []int{10,20,30}
	sum := 0

	for index := range nums {
		fmt.Println(index)
		sum += index
	}

	fmt.Println("sum:", sum)

	for _, value := range nums {
		fmt.Println(value)
		sum += value
	}

	fmt.Println("sum:", sum)

	for  index, value := range nums {
		if value == 3 {
			fmt.Println("index:", index)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	sum = 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
