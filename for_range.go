package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}

	for i, l := range arr {
		fmt.Println(i, l)
	}

	for _, j := range arr {
		fmt.Println(j)
	}
}
