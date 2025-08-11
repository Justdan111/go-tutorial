package main

import (
	"errors"
	"fmt"
    "strconv"
)

// Custom error type
type ValidationError struct {
	Field string
	Message string
}

// Implement the error interface
func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", v.Field, v.Message)
}

// user struct for validation
type User struct {
	Name string
	Email string
	Age int
}

// function that returns an error
func divide (a, b float64) (float64, error) {
	if b == 0 {
		return 0,
	errors.New("division by Zero")	
	}
	return  a / b, nil
}


func main() {

}