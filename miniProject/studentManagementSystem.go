package main

import (
	"fmt"
)

// Student struct to hold student information
type Student struct {
	Name  string
	Age   int
	Grade string
}

// Method with VALUE receiver (does not modify original)
func (s Student) PrintInfo() {
	fmt.Printf("Student info -> Name: %s, Age: %d, Grade: %s\n", s.Name, s.Age, s.Grade)
}

// 