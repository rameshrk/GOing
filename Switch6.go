package main

import "fmt"

type Page interface {
	PrintPage()
}

type HtmlPage struct {
	// Implement Page interface.
	Page
}

type Image interface {
	PrintImage() // not required to work
}

type ImagePage struct {
	// Implement Image interface.
	//Page    // will work
	Image
}

func test(value interface{}) {
	// Use type switch to test interface type.
	// ... The argument is an interface.
	switch value.(type) {
	case nil:
		fmt.Println("Is nil interface")
	case Page:
		fmt.Println("Is page interface");
	case Image:
		fmt.Println("Is image interface");
	default:
		fmt.Println("not know", value)
	}
}

func main() {
	// Create class that implements interface and pass to test func.
	item1 := new(HtmlPage)
	test(item1)

	item2 := new(ImagePage)
	test(item2)
	fmt.Println()

	test(new (Page))
	test(new (Image))
	test(new (HtmlPage))
	test(new (ImagePage))


}