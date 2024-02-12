package main

import "fmt"

func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	var n1 float64 = 45.89
	var n2 int = int(n1)
	fmt.Printf("%v\n", n2)
}
