package main

import "fmt"

func main() {
	
	// Basic types and variables
	var name string = "John Doe"
	var age int = 25
	var height float64 = 5.6
	var isStudents bool = true

	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Height: %.2f (type: %T)\n", height, height)
	fmt.Printf("Is Student: %t (type: %T)\n", isStudents, isStudents)

	// Arrays - fixed size
	var scores [5]int = [5]int{90, 85, 88, 92, 95}	
	fmt.Printf("Scores array: %v\n", scores,)

	// Slices - dynamic size
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("Fruits slice: %v\n", fruits)

	// Add to slice
	fruits = append(fruits, "grape")
	fmt.Printf("After adding grape: %v\n", fruits)
}