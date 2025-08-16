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

	 // Modify through pointer
	 *ptr = 100
	 fmt.Printf("After *ptr = 100, x = %d\n", x)
	 
	 fmt.Println("\n=== Pointers with Structs ===")
	 
	 // Create person
	 person := Person{Name: "Alice", Age: 25}
	 fmt.Printf("Original person: %+v\n", person)
	 
	 // Get pointer to person
	 personPtr := &person
	 fmt.Printf("Pointer: %p\n", personPtr)
	 fmt.Printf("Value at pointer: %+v\n", *personPtr)
	 
	 // Modify through pointer
	 personPtr.Name = "Alice Smith" // Go automatically dereferences
	 (*personPtr).Age = 26          // Explicit dereferencing
	 
	 fmt.Printf("After modification: %+v\n", person)
	 
	 fmt.Println("\n=== Value vs Pointer Receivers ===")
	 
	 person2 := Person{Name: "Bob", Age: 30}
	 
	 // Both work, Go handles conversion automatically
	 person2.IntroduceValue()
	 person2.IntroducePointer()
	 
	 // Birthday method modifies, so it needs pointer receiver
	 person2.HaveBirthday()
	 fmt.Printf("After birthday: %+v\n", person2)
	 
	 fmt.Println("\n=== Function Parameters: Value vs Pointer ===")
	 
	 person3 := Person{Name: "Carol", Age: 35}
	 fmt.Printf("Before function calls: %+v\n", person3)
	 
	 // Pass by value - original unchanged
	 modifyPersonValue(person3)
	 fmt.Printf("After value function: %+v\n", person3)
	 
	 // Pass by pointer - original changed
	 modifyPersonPointer(&person3)
	 fmt.Printf("After pointer function: %+v\n", person3)
	 
	 fmt.Println("\n=== Pointer to Pointer ===")
	 
	 a := 10
	 ptrA := &a
	 ptrPtrA := &ptrA
	 
	 fmt.Printf("a = %d\n", a)
	 fmt.Printf("ptrA = %p, *ptrA = %d\n", ptrA, *ptrA)
	 fmt.Printf("ptrPtrA = %p, *ptrPtrA = %p, **ptrPtrA = %d\n", 
		 ptrPtrA, *ptrPtrA, **ptrPtrA)
	 
	 fmt.Println("\n=== Nil Pointers ===")
	 
	 var nilPtr *Person
	 fmt.Printf("nilPtr = %v\n", nilPtr)
	 
	 // Check for nil before using
	 if nilPtr == nil {
		 fmt.Println("Pointer is nil, cannot dereference")
	 }
	 
	 // Create new Person with new()
	 newPerson := new(Person) // Returns pointer to zero-value Person
	 fmt.Printf("newPerson = %p, value = %+v\n", newPerson, *newPerson)
	 
	 newPerson.Name = "David"
	 newPerson.Age = 28
	 fmt.Printf("After setting values: %+v\n", *newPerson)
	 
	 fmt.Println("\n=== Slices and Pointers ===")
	 
	 // Slices are reference types (contain pointers internally)
	 slice1 := []int{1, 2, 3}
	 slice2 := slice1 // Both point to same underlying array
	 
	 fmt.Printf("slice1: %v\n", slice1)
	 fmt.Printf("slice2: %v\n", slice2)
	 
	 // Modify through slice2
	 slice2[0] = 999
	 fmt.Printf("After slice2[0] = 999:\n")
	 fmt.Printf("slice1: %v\n", slice1) // Also changed!
	 fmt.Printf("slice2: %v\n", slice2)
	 
	 fmt.Println("\n=== When to Use Pointers ===")
	 fmt.Println("1. When you need to modify the original value")
	 fmt.Println("2. When copying would be expensive (large structs)")
	 fmt.Println("3. When you want to share data between functions")
	 fmt.Println("4. For optional values (nil represents 'no value')")
	 
	 // Example: optional value
	 var optionalAge *int
	 
	 if optionalAge == nil {
		 fmt.Println("Age not provided")
	 } else {
		 fmt.Printf("Age: %d\n", *optionalAge)
	 }
	 
	 // Set optional value
	 age := 25
	 optionalAge = &age
	 fmt.Printf("Age provided: %d\n", *optionalAge)
 }

