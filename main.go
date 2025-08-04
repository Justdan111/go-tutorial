package main

import "fmt"

func main()  {
	fmt.Println("Hello, World!")

	var name string = "John Doe"
	age := 30

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	var x, y int = 5, 10
	sum := x + y
	fmt.Printf("The sum of %d and %d is %d.\n", x, y, sum)
}