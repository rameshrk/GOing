package main

import "fmt"

type person struct {
	name string
	age int
}


func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "naame"})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"An", 90})
	fmt.Println("\n")

	s := person{name: "Sean", age:9}
	fmt.Println(s.name)
	fmt.Println("\n")
	sp := &s
	fmt.Println(*sp)
	fmt.Println(sp.age)
	fmt.Println(sp.name)
	fmt.Println("\n")
	sp1 := s

	sp. age = 5

	fmt.Println(sp)
	fmt.Println(*sp)
	fmt.Println(sp1)
}