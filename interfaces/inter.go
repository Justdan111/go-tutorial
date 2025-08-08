package main

import fmt

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


func main () {

}