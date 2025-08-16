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

func main() {

}