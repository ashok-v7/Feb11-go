package main

import "fmt"

func main() {
	for i := 3; i < 5; i++ {
		fmt.Println(i + (i * 1))

	}

	j := 3
	for j < 5 {
		fmt.Println(j + (j * 1))
		j += 1
	}
}
