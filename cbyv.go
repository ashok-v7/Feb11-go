package main

import "fmt"

func modify(s string) {
	s = "world"
}

func mint(j int) int {
	j += 10
	return j
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)

	j := 10
	fmt.Println("j: ", j)
	res := mint(j)
	fmt.Println("j: ", res)
}
