package main

import "fmt"

const pi float32 = 3.14

func main() {
	defer fmt.Println(pi)
	fmt.Println(test())
	fmt.Println(test1())
	s1, s2 := test()
	fmt.Println(s1, s2)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if i == 7 {
			fmt.Println(fmt.Sprintf("It is seven (%d)", i))
		}
	}
	fmt.Println(sum)
	fmt.Println(isTest(3))
}

func test() (a string, b string) {
	a = "Hello"
	b = "World"
	return a, b
}

func test1() string {
	return "My World"
}

func isTest(a int) (bool, int) {
	if a == 2 {
		return true, 1
	}
	return false, 2
}
