package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

// calculatePercentage calculates the average of the grades as a percentage.
func (s *Student) calculatePercentage() float64 {

	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	average := float64(sum) / float64(len(s.grades))
	return average // This is already a percentage of the scale, assuming 100 is the max grade.
}

func main() {
	s := Student{name: "Joe", grades: []int{90, 75, 80}}
	s.displayName()
	fmt.Printf("%.2f%%\n", s.calculatePercentage())
}
