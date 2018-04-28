package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	sumi := 0
	for i, num := range nums {
		sumi += i
		sum += num
	}
	fmt.Println(sumi, sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s - %s\n", k, v)
	}

	for k, v := range kvs {
		fmt.Printf("%s %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
