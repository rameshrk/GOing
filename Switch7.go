package main

import "fmt"

type Page1 interface {
	PrintPage1()
}

type HtmlPage1 struct {
	// Implement Page interface.
	Page1
}

type Image1 interface {
	//PrintImage1() // not required to work
}

type ImagePage1 struct {
	// Implement Image interface.
	//Page    // will work
	Image1
}

func test1(value interface{}) {
	// Use type switch to test interface type.
	// ... The argument is an interface.
	switch value.(type) {
	case nil:
		fmt.Println("Is nil interface")
	case Page1:
		fmt.Println("Is page interface");
	case Image1:
		fmt.Println("Is image interface");
	default:
		fmt.Println("not know", value)
	}
}

func main() {
	// Create class that implements interface and pass to test func.
	item1 := new(HtmlPage1)
	test1(item1)

	item2 := new(ImagePage1)
	test1(item2)
	fmt.Println()

	test1(new (Page1))
	test1(new (Image1))
	test1(new (HtmlPage1))
	test1(new (ImagePage1))


}