package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [5]int
	bArray := [5] int{1, 2, 3}
	s := make([]string, 3)
	t := []string{"g", "h", "i"}

	fmt.Println(a)

	a[1] = 100
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println()

	//i := make ([]int,2)
	s[0] = "a"
	s[1] = "bArray"
	s[2] = "c"

	//a = a.append(s, int)

	fmt.Println(bArray)
	fmt.Println(s)
	//fmt.Println(append(s, bArray))

	fmt.Println()

	var twoD [2][3] int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	myStringSlice := []string{"first", "second", "third"}
	fmt.Println(myStringSlice)
	myStringSlice = append(myStringSlice, []string{"fourth", "fift"}...)
	fmt.Println(myStringSlice)
	myStringSlice = append(myStringSlice, "fourth", "fifth")
	fmt.Println(myStringSlice)

	myArray := []string{"first", "second", "third"}
	fmt.Println(myArray)
	myArray = append(myArray, []string{"fourth", "fift"}...)
	fmt.Println(myArray)
	myArray = append(myArray, "fourth", "fifth")
	fmt.Println(myArray)

	s = make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("sl2", c)

	fmt.Println(s[:5])
	fmt.Println(s[2:])
	fmt.Println(t)
	fmt.Println()

	t2 := []string{"g", "h", "i"}
	fmt.Println(t2)
	t3 := [3]string{"g", "h", "i"}
	fmt.Println(t3)
	twoA := make([][]int, 2)
	twoB := [2][3] int{}
	fmt.Println(twoA)
	fmt.Println(twoB)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			innerLen := i + 1
			twoA[i] = make([]int, innerLen)
			twoB[i][j] = i + j
		}
	}
	fmt.Println(twoA)
	fmt.Println(twoB)

	GetTypeArray(t2)

}

func GetTypeArray(arr []interface{}) reflect.Type {
	return reflect.TypeOf(arr[0])
}