package main

import "fmt"

// define a struct -like class in other languages
type Person struct {
	Name string
	Age  int
	Email string
}

// method with a receiver
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func main() {

}