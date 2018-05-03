package main

import "fmt"

type rect struct {
	width, height int
}

type circle struct {
	radius int
}

func (c *circle) area() int {
	return c.radius * c.radius * 2 * 3
}
func (c *circle) perimemter() int {
	return 2 * c.radius * 3
}

func (r *rect) area() int {
	r.height = r.height * 2
	return r.width * r.height
}
//func (r rect) area() int {
//	r.height = r.height * 2
//	return r.width * r.height
//}
func (r rect) perim() int {
	r.height = r.height * 2
	return 2*r.width + 2*r.height
}

func color() int {
	return 2
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("perim:", r.perim())
	fmt.Println(r.width, r.height)
	fmt.Println("area: ", r.area())
	fmt.Println(r.width, r.height)

	fmt.Println("", )
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println(r.width, r.height)
	fmt.Println("", )

	r.height = 10
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	color()

	c := circle{radius: 5}
	fmt.Println(c.area())
	fmt.Println(c.perimemter())

}
