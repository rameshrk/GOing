package main

import "fmt"

func main()	{
	nums := []int{20,30,40}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	indexsum :=0
	sum = 0
	for  index, num := range nums {
		indexsum+=index
		sum+=num

	}
	println("****" , indexsum , sum)

	fmt.Println("Outside sum:", sum)

	for i, num :=range nums {
		fmt.Println("\n\nIn loop. Value: ", num)
		fmt.Println("in loop before. Sum:", sum)
		sum := sum+ 1
		fmt.Println("in loop after. Sum:",sum)

		if num ==3 {
			//fmt.Println("index", i)
		}
		fmt.Println(len(nums))
		if i == len(nums)-1 {
			fmt.Println("sum",sum)
		}
	}

	for i :=0; i <10; i++ {
		sum:= sum+ 1
			fmt.Println("Sum in second for loop",sum)
	}

	for i :=0; i <10; i++ {
		sum = sum+ 1

		fmt.Println("Sum in third for loop",sum)

	}

	for  i, num := range nums {
		if num == 3{
			fmt.Println("index", i)
		}
	}


	fmt.Println("Sum", sum)

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range kvs {
		fmt.Println("%s -> %s\n",k,v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for index,c := range "go" {
		fmt.Println(index,c)
	}

}