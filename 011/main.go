package main

import (
	"fmt"
)

// Define a struct for a student
type Student struct {
	Name   string
	School string
	Year   int
}

func main() {
	// Create a new student instance
	student1 := Student{
		Name:   "John Doe",
		School: "01Founders",
		Year:   2023,
	}

	// Print the student's details
	fmt.Println("Student Information:")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:", student1.Year)

}
