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


	s := person{name: "Sean", age:9}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(*sp)
	fmt.Println(sp.age)
	fmt.Println(sp.name)

	sp. age = 5
	fmt.Println(sp)
	fmt.Println(*sp)
}