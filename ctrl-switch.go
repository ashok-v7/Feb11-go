package main

import "fmt"

func main() {

	i := 10
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100, 200:
		fmt.Println("i is 100 or 200")
	default:
		fmt.Println("i is neiher 10 ,100, 200")
	}
}
