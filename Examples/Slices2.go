package main

import "fmt"

func main() {

	nums := []int{10,20,30}
	fmt.Println(nums)
	kvs := map[string]string{"a":"apple", "m":"mango", "b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s  %s", k,v)
	}

	for k := range kvs {
		fmt.Println("key",k)
	}


	for i, c := range "go" {
		fmt.Println(i,c)
		fmt.Println("")
		fmt.Printf("\n%c",c)
	}
}