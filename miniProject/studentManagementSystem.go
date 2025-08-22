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

// Method with POINTER receiver (modifies original)
func (s *Student) UpdateGrade(newGrade string) {
	s.Grade = newGrade
}

// Function: pass by VALUE (does not modify original)
func ChangeNameValues(s Student, newName string) {
	s.Name = newName  // only changes the copy
}

// Function: pass by POINTER (modifies original)
func ChangeNamePointer(s *Student, newName string) {
	s.Name = newName  // changes the original
}