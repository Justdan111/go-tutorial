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

// method with pointer reciever - can modify the struct
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy birthday %s! You are now %d years old.\n", p.Name, p.Age)
}

// method to validate email
func (p Person) isEmailValid() bool {
	return len (p.Email) > 0 && contains(p.Email, "@")
}

func main() {

}