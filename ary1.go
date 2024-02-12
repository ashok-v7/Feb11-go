package main

import "fmt"

func main() {
	grades := [5]int{90, 80, 70, 80, 97}
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}

	for i := 0; i < len(grades); i++ {
		fmt.Println(i, "=>", grades[i])
	}

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice := arr[1:8]
	fmt.Println(cap(arr))   // 10
	fmt.Println(cap(slice)) // 9
	fmt.Println("arr =>", arr)
	fmt.Println("slice =>", slice)
}
