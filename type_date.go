package main

import "fmt"

func main() {
	pointers()
	structs()
}

func pointers() {
	a := "hello world"
	fmt.Println(a)

	p := &a
	fmt.Println(p)
	fmt.Println(*p)

	*p = "hi"
	fmt.Println(p)
	fmt.Println(*p)

	b := 42
	g := &b
	*g = *g / 2
	fmt.Println(b)

}

type Point struct {
	X int
	Y int
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{
		X: 321,
	}
	fmt.Println(p2)
}
