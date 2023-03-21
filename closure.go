package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	fmt.Println("hi")
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	orderPrive := totalPrice(1)
	fmt.Println(orderPrive)
	fmt.Println(orderPrive(1))
	fmt.Println(orderPrive(1))

}
