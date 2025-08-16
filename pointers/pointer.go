package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// function with value receiver
func (p Person) IntroduceValue() {
	fmt.Printf("Hi, I'm %s (value receiver)\n", p.Name)
}

// Method with pointer receiver - works with original
func (p *Person) IntroducePointer() {
    fmt.Printf("Hi, I'm %s (pointer receiver)\n", p.Name)
}

// Method that modifies - needs pointer receiver
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

// Function that takes a value (copy)
func modifyPersonValue(p Person) {
    p.Name = "Modified Name"
    p.Age = 999
    fmt.Printf("Inside function (value): %+v\n", p)
}

// Function that takes a pointer (original)
func modifyPersonPointer(p *Person) {
    p.Name = "Modified Name"
    p.Age = 999
    fmt.Printf("Inside function (pointer): %+v\n", *p)
}

func main() {

	fmt.Println("=== Understanding Pointers ===")
    
    // Basic variables and pointers
    x := 42
    fmt.Printf("x = %d\n", x)
    
    // Get pointer to x
    ptr := &x
    fmt.Printf("ptr = %p (points to x)\n", ptr)
    fmt.Printf("*ptr = %d (value at pointer)\n", *ptr)

}