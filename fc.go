package main

import "fmt"

func adder(x int, y int) (int, int) {
	return x + y, x - y
}

func main() {

	result, diff := adder(2, 3)
	fmt.Println("result : ", result, diff)

}
