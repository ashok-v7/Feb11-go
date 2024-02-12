package main

import "fmt"

func mofy(s *string) {
	*s = "world"
}

func main() {
	a := "hello"
	fmt.Println(a)
	mofy(&a)
	fmt.Println(a)
}

//Slices are passed by reference, by default.
//Maps, as well, are passed by reference, by default
