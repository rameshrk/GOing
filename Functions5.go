package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusp(a, b, c int) int {
	return a + b + c
}

func plusplus(a, b, c int) (int, int, int) {
	return a, b, c
}

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusp(1, 2, 3))
	fmt.Println(plusplus(1, 2, 3))
	a, _, c := plusplus(2, 34, 4)
	fmt.Println(a, c)
	fmt.Println(sum(1, 2, 3, 4, ))
	fmt.Println(sum(1, 2, 3, sum(1, 23, 4)))

	nums := []int{1, 2, 3, 4}
	//nums1 := [4]int{1, 2, 3, 4}
	//fmt.Println(nums...)
	//fmt.Println(nums1...)
	fmt.Println(sum(nums...))

	//fmt.Println(sum(nums))
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
