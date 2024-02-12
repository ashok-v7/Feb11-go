// package main

// import "fmt"

// type Student struct {
// 	name   string
// 	rollNo int
// 	marks  []int
// 	grades map[string]int
// }

// func main() {

// 	s1 := Student{
// 		name:   "priyanka",
// 		rollNo: 1,
// 		marks:  []int{90, 80, 70},
// 		grades: map[string]int{
// 			"A": 90,
// 			"B": 80,
// 			"C": 70},
// 	}
// 	fmt.Println("%+v", s1)
// }

package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {

	s1 := Student{
		name:   "priyanka",
		rollNo: 1,
		marks:  []int{90, 80, 70},
		grades: map[string]int{
			"A": 90,
			"B": 80,
			"C": 70},
	}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%+v\n", s1)
	// If the value is a struct, the %+v variant will include the struct's field names.
	// The %#v variant prints a Go syntax representation of the value
	fmt.Println(s1.name)
}
