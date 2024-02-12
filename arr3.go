package main

import "fmt"

func main() {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]
	fmt.Println(slice)                  // 20 30
	fmt.Println(len(slice))             // 2
	fmt.Println(cap(slice))             // 3
	slice = append(slice, 900, -90, 50) // 20 30 900 -90 50
	fmt.Println(slice)                  // 20 30 900 -90 50
	fmt.Println(len(slice))             // 5
	fmt.Println(cap(slice))             // 6
}
