package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 12
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println(len(m))

	v1 = m["k1"]
	fmt.Println("v1", v1)

	delete(m,"k2")
	fmt.Println("map", m)

	_, prs := m["k2"]
	fmt.Println("prs",prs)

	m["k2"] = 5
	_, prs = m["k2"]

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println()

	n = map[string]int {"foo":1,"bar":2}
	fmt.Println("map", n)



}