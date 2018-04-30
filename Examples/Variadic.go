package main

import "fmt"

func main() {

	sum2("a",1,2,3)
	sum2("a",1)
	num := []int {1,2,3,4,5,5}
	nums := []int{1, 2, 3, 4}

	sum(num...)
	sum(nums...)

}



func val(a int , b int ) (int, int) {
	fmt.Println()
	return a, b
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sum2(nums2 string, nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	fmt.Println(nums2)
}
