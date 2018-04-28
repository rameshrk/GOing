package main

import "fmt"

func main() {

	m := make(map[string]int)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)

	m["k1"] = 7
	m["k2"] = 7

	fmt.Println(m)
	m["k2"] = 10
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	v := m
	fmt.Println(v)
	fmt.Println(m)
	v["k1"] = 8
	fmt.Println(v)
	fmt.Println(m)

	delete(m, "k2")

	fmt.Println(v)
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	_, prs = m["k1"]
	fmt.Println("prs", prs)

	fmt.Println(n)
}
