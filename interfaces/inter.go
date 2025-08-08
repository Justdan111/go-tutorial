package main

import "fmt"

// interface defines behavior, not data
type Speaker interface {
	Speak() string
}

// interface for things that can move
type Mover interface {
	Move() string
}

// combined interface
type Animal interface {
	Speaker
	Mover
}

// Structs that will implement these interface
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Robot struct {
	Model string
}

// Dog implements speaker interface
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says woof!", d.Name)
}

// Dog implements Mover interface
func (d Dog) Move() string {
	return fmt.Sprintf("%s runs around", d.Name)
}

// Cat implements speaker interface
func (c Cat) Speak() string {
	return  fmt.Sprintf("%s says meow!", c.Name)
}

// cat implements Mover interface
func (c Cat) Move() string {
	return fmt.Sprintf("%s prowls silently", c.Name)
}

// Robot implemets Speaker intertface
func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: BEEP BOOP ", r.Model)
}

// Robot implements mover interface
func (r Robot) Move() string {
	return  fmt.Sprintf("Robot %s moves mechanically", r.Model)
}

// Function that accepts any speaker
func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// function that accepts any Animal (Speaker + Mover)
func describeAnimal(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.Move())
}

// Empty interface - can hold any type
func describe(i interface{}) {
	fmt.Printf("Value: %v, type: %T\n", i, i)
}


func main () {
	// Create instances
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	robot := Robot{Model: "R2D2"}

	// All implement Speaker, so can be passed to makeItSpeak
	 fmt.Println("=== Making them speak ===")
	    makeItSpeak(dog)
		makeItSpeak(cat)
		makeItSpeak(robot)

	// Animals (cat and dog) implement both speaker and mover
	fmt.Println("\n=== Describing animals ===")
	describeAnimal(dog)
	fmt.Println()
	describeAnimal(cat)
	fmt.Println()
	// Robot implements both speaker and mover
	describeAnimal(robot)

	// Empty interface example
	 fmt.Println("\n=== Empty interface example ===")
	 describe(42)
	 describe("Hello, world!")
	 describe(dog)
	 describe([]int{1, 2, 3})

	// slice of interface
	fmt.Println("\n === Slice of Speakers ===")
	speakers := []Speaker{dog, cat, robot}
	 for i, speaker := range speakers {
		fmt.Printf("%d, %s\n", i+1, speaker.Speak())
	 }

	// Type assertion - getting concrete type interface 
	fmt.Println("\n === Type assertions ===") 
	var s Speaker = dog

	// check if it's a Dog
	if d, ok := s.(Dog); ok {
		fmt.Printf("It's a dog named %s!\n", d.Name)
	}

	// Type switch
	switch v := s.(type) {
	case Dog:
		fmt.Printf("Definitely a dog: %s\n", v.Name)
	case Cat:
		fmt.Printf("It's a Cat: %s\n", v.Name)
	case Robot:
		fmt.Printf("It's a robot: %s\n", v.Model)
	default:
		fmt.Printf("Unknown type: %T\n", v)			
	}	
}