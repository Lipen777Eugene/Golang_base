package main

import "fmt"

func main() {
	var a interface{} = 4

	switch a.(type) {
	case string:
		fmt.Println("this is string")
	case float32:
		fmt.Println("this is float")
	case int:
		fmt.Println("this is int")
	default:
		fmt.Printf("unknown type %T\n", a)
	}

}
