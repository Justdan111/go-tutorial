package main

import "fmt"

// Functions with parameters and return values
func greet (name string) string {
	return "Hello, " + name + "!"
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// function with name return value
func calculateGrade(score int) (grade string, passed bool) {
	if score >= 90 {
		grade = "A"
		passed = true
	} else if score >= 80 {
		grade = "B"
		passed = true
	} else if score >= 70 {
		grade = "C"
		passed = true
	} else if score >= 60 {
		grade = "D"
		passed = true
	} else {
		grade = "F"
		passed = false
	}
	return
}
func main() {
	// function calls
	message := greet("Go Developer")
    fmt.Println(message)	

	// multiple return values
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 divided by 3 is %.2f\n", result)
	}

	// try divide by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// Named returns
	grade, passed := calculateGrade(85)
	fmt.Printf("Score: 85, Grade: %s, Passed: %t\n", grade, passed)

}