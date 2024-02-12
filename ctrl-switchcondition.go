package main

import "fmt"

func main() {

	a, b := 10, 20
	switch {
	case a+b == 30:
		fmt.Println("a+b == 30")
	case a-b == -10:
		fmt.Println("a-b == -10")
	default:
		fmt.Println("greater than 30")

	}
}
