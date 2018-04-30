package main

import "fmt"

var n =50
func main(){
	i := 1
	for i<=3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j<=9 ; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println(n)
	for n := 0; n <= 5; n++	{
		if n%2 == 0 {
			continue
		} else if n==5 {
			break
		}

		fmt.Println(n)
	}


}