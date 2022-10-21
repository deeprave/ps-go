package main

import "fmt"

func main() {
	name, course := "Nigel Poulton", "Go Fundamentals"
	// module, clip := 4, 2
	// courseComplete := false

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("you're currently watching", course)
}

func updateCourse(course *string) string {
	*course = "Getting Started with Docker"

	fmt.Println("Updated course to", *course)
	return *course
}
