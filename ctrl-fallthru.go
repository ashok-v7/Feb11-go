package main

import "fmt"

func main() {
	i := 10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")
	}
}

//The fallthrough keyword is used in switch-case to force the execution flow to fall through the successive case block
