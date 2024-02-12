package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)

	var s1 string = "200abc"
	i2, err := strconv.Atoi(s1)
	fmt.Printf("%v, %T \n", i2, i2) // 0 , int
	fmt.Printf("%v, %T", err, err)  //strconv.Atoi: parsing "200a": invalid syntax
}
