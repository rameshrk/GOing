package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Alice", 30})
	fmt.Println(person{name: "Fr"})
	fmt.Println(&person{age: 50, name: "Fr"})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	s.name = "Sean1"
	fmt.Println(s.name)


	sp := &s
	fmt.Println(sp.name, sp.age)

	sp.age = 51
	fmt.Println(sp.name,sp.age)

}
