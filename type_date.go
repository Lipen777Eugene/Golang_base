package main

import "fmt"

func main() {
	pointers()
	structs()

	p3 := Point{
		X: 7,
		Y: 9,
		S: "hi",
	}
	p3.method1()
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
	S string
}

func (p Point) method1() {
	fmt.Println(p.X, p.Y, p.S)
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "Point1",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{
		X: 321,
	}
	fmt.Println(p2)

	g := &p1
	fmt.Println(g.X)

}
