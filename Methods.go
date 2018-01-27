package main

import "fmt"

type rect struct {
	 width, height int
}

func (r* rect) area() int {
	return r.width * r.height
}


func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}


func main() {
	r := rect{10,20}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())88


	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())

	rp.width =40
	rp.height = 30
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())

	fmt.Println(r.area())
	fmt.Println(r.perimeter())


}