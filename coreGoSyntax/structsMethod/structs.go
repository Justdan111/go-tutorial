package main

import "fmt"

// define a struct -like class in other languages
type Person struct {
	Name string
	Age  int
	Email string
}

// method with a receiver
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s I am %d years old.", p.Name, p.Age)
}

// method with pointer reciever - can modify the struct
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

// method to validate email
func (p Person) isEmailValid() bool {
	return len(p.Email) > 0 && contains(p.Email, "@")
}

// Helper function 
func contains (s, substr string) bool {
	for i := 0; i <= len(s) - len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return  true
		}
	}
	return  false
}

// Another struct for demonstrating composition
type Employee struct {
	Person
	Company string
	Salary float64
}

// Method for Employee 
func (e Employee) GetDetails() string {
	return  fmt.Sprintf("%s works at %s with salary $%.2f", e.Name, e.Company, e.Salary)
}


func main() {
	//Create a Person
	person1 := Person{
		Name: "Dan",
		Age: 25,
		Email: "dan@email.com",
	}

	// Call Method

	fmt.Println(person1.Introduce())
		fmt.Printf("Email valids: %t\n", person1.isEmailValid())

		// Modify using pointer method
		person1.HaveBirthday()

		// Another way to create struct
		person2 :=Person{}
		person2.Name = "John"
		person2.Age = 30
		person2.Email = "john@gmail.com"

		fmt.Println(person2.Introduce())

		// Create an Employee (embedded struct)
		emp := Employee{
			Person: Person{
				Name:  "Alice",
				Age:   28,
				Email: "Al@gmail.com",
			},
			Company: "TechCorp",
			Salary: 75000.00,
		}

	// Can access Person methods directly
	fmt.Println(emp.Introduce())
	fmt.Println(emp.GetDetails())

	// Can access Person fields directly too
	fmt.Printf("Employee Email: %s\n", emp.Email)

	// Slice of structs
    team := []Person{person1, person2}
	fmt.Println("\n Intoduction:")
	for _, member := range team {
		fmt.Println(member.Introduce())
	}

}