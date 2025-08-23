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

func main() {
	// 1 Basic Variables + pointers
	x := 42
	prt := &x
	fmt.Println("Original x:", x)
	*prt = 100
	fmt.Println("Modified x through pointer:", x)

	// 2 Structs and Pointers
	student1 := Student{"John", 20, "B"}
	studentPtr := &student1
	fmt.Println("\nAccess via struct pointer:", studentPtr.Name)

	// 3. Methods: value vs pointer receiver
	student1.PrintInfo()
	student1.UpdateGrade("A+")
	student1.PrintInfo()

	// 4 Functions: pass by value vs pointer
	ChangeNameValues(student1, "Mike")
	student1.PrintInfo() // Name should still be "John"

	ChangeNamePointer(&student1, "Mike")
	student1.PrintInfo() // Name should now be "Mike"
	
	// 5. nil pointer check
	var student2 *Student
	if student2 == nil {
		fmt.Println("\nstudent2 is nil, creating new Student")
		student2 = new(Student)
		student2.Name = "Peter"
		student2.Age = 22
		student2.Grade = "C"
	}
	student2.PrintInfo()
	
	// 6. Slices with pointers (reference behavior)
	students := []Student{
		{"Clark", 21, "B"},
		{"Diana", 23, "A"},
	}
	fmt.Println("\nOriginal Student slice", students)

	// Modify slice element (affects original, since slices are references)
	students[0].Grade = "A+"
	fmt.Println("Modified Student slice", students)

	// 7. Pointer to Pointer (rarely used)
	prtToPtr := &studentPtr
	(**prtToPtr).Grade = "B-"
	fmt.Println("\nAfter modifying via pointer to pointer:", student1)
}