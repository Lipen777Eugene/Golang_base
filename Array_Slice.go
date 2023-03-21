package main

import "fmt"

func main() {
	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	a := animalsArr[0:2]
	fmt.Println(a)
	var b []string = animalsArr[1:3]
	fmt.Println(b)

	b[0] = "123"

	fmt.Println(a)
	fmt.Println(animalsArr)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[:5]
	fmt.Println(t)
	t = s[5:]
	fmt.Println(t)

}

func arrays() {
	// array
	var a [2]string
	a[0] = "hi"
	a[1] = "hello"
	fmt.Println(a, a[0], a[1])

	number := [...]int{1, 2, 3}
	fmt.Println(number)
	number[2] = 4
	fmt.Println(number)
}

func slices() {
	//slices

	letters := []string{"a", "b"}
	letters[0] = "1"
	letters = append(letters, "d")
	fmt.Println(letters)
	letters = append(letters, "e", "f")
	fmt.Println(letters)

	createSlice := make([]string, 3)
	createSlice[0] = "a"
	createSlice[1] = "b"
	createSlice[2] = "c"
	fmt.Println(createSlice)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "d")
	fmt.Println(createSlice)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "e", "f", "g")
	fmt.Println(createSlice)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
}
