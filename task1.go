package main

import "fmt"

func main() {
	var (
		a uint
	)

	for {
		fmt.Scan(&a)
		if a <= 10 {
			continue
		} else if a >= 100 {
			break
		}
		fmt.Println(a)
	}

}
