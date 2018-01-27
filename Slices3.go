package main
import "fmt"

func main() {

	s:= make ([]string,3)
	fmt.Println("emp",s)


	s[0] = "0"
	s[1] = "1"
	s[2] = "2"

	fmt.Println("set",s)
	fmt.Println("get",s[2])

	s = append(s,"3")
	s = append(s, "4","5")
	fmt.Println("apd",s)


	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy",c,s)

	l := s[2:5]
	fmt.Println("sl1",l)

	l = s[:5]
	fmt.Println("sl2",l)

	l = s[2:]
	fmt.Println("sl3",l)




}