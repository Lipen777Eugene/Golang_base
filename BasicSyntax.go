package main

import "fmt"

const pi float32 = 3.14

func main() {
	fmt.Println(pi)
	fmt.Println(test())
	fmt.Println(test1())
	s1, s2 := test()
	fmt.Println(s1, s2)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func test() (a string, b string) {
	a = "Hello"
	b = "World"
	return a, b
}

func test1() string {
	return "My World"
}
