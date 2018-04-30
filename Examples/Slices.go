package main

import "fmt"

func main() {

	s:= make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len", len(s))

	s=append(s, "d")
	s = append(s,"e", "f")
	fmt.Println("append", s)

	c:= make([]string, len(s))
	copy( c, s)
	fmt.Println(c)
	fmt.Println(s)

	l := s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	t := []string{"g","h","i"}
	fmt.Println("dcl:",t)
	fmt.Println("_____-")
	twoD := make([][]int,3)
	for i:= 0; i<3;i++{
		innerlen := i+1
		twoD[i] = make([]int, innerlen)
		for j:=0;j< innerlen ; j++ {
			twoD[i][j] = i+j
		}

	}
	fmt.Println("2d", twoD)
}