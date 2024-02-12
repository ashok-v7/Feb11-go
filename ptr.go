// A pointer is a variable that holds memory address of another variable.
// They point where the memory is allocated and provide ways to find or even change the value located at the memory location.
package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))
}
