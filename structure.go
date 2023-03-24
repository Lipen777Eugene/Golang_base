package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) movePoint2(x, y int) Point {
	p.X *= x
	p.Y *= y
	return p
}

func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePoint3() {
	p.X += 10
	p.Y += 20
}

func main() {
	p := Point{X: 1, Y: 2}
	fmt.Println(p)
	fmt.Println(movePoint(p, 1, 1))
	p.movePoint2(2, 2)
	fmt.Println(p)
	p.movePoint3()
	fmt.Println(p)
}
