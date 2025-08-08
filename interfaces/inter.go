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

func main () {

}