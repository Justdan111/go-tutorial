package main

import (
	"errors"
	"fmt"
    "strconv"
)

// Custom error type
type ValidationError struct {
	Field string
	Message string
}

// Implement the error interface
func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", v.Field, v.Message)
}

// user struct for validation
type User struct {
	Name string
	Email string
	Age int
}

// function that returns an error
func divide (a, b float64) (float64, error) {
	if b == 0 {
		return 0,
	errors.New("division by Zero")	
	}
	return  a / b, nil
}

// Function with custom error
func validateUser (u User) error {
	if u.Name == " " {
		return ValidationError{
			Field: "name",
			Message: "name cannot be empty",
		}
	}
	if u.Age < 0 || u.Age > 150 {
		return ValidationError{
			Field: "age",
			Message: "age must be between 0 and 150",
		}
	}
	if len(u.Email) == 0 {
		return  ValidationError{
			Field: "email",
			Message: "email is required",
		}
	}
	return nil // No error
}

// Function that might return multiple types of errors
func parseAndValidateAge(ageStr string) (int, error) {
	// First, try to parse the string
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0,
	fmt.Errorf("failed to parse age '%s': %w", ageStr, err)	
	}

	// then validate the parsed age
	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}
	if age > 150 {
	   return 0, errors.New("age cannot be greater than 150")
	}   
	return age, nil
}

// demostration of error wrapping and unwrapping
func processUserData(name, email, ageStr string) (*User, error) {
	age, err := parseAndValidateAge(ageStr)
	if err != nil {
		return nil,
		fmt.Errorf("error processing user data: %w", err)
	}

	User := User{
		Name: name,
		Email: email,
		Age: age,
	}

	if err := validateUser(User);
	err != nil {
		return nil,
		fmt.Errorf("user valodation failed: %w", err)
	}
	return &User, nil
}

func main() {
    fmt.Println("=== Basic Error Handling ===")

	//bassic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
  
	fmt.Println("\n === Custom Error Types ===")

	// Invalid user
    invalidUser := User{Name: "", Email: "", Age: -5}
    if err := validateUser(invalidUser); err != nil {
        fmt.Printf("Validation error: %v\n", err)
        
        // Type assertion to get specific error details
        if validationErr, ok := err.(ValidationError); ok {
            fmt.Printf("Field with error: %s\n", validationErr.Field)
        }
    }

	fmt.Println("\n === Error Wrapping ===")

	// Successful user processing
	user, err := processUserData("Dan", "slime@gmail.com", "25")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Processed user: %+v\n", user)
	}

	// Error in Parsing
	_, err = processUserData("Dan", "slime@gmail.com", "invalid_age")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n=== Error Checking Patterns ===")

	// Multiple error checks - common Go pattern
	ages := []string{"25", "-5", "invalid_age", "200"}

	for _, ageStr := range ages {
        age, err := parseAndValidateAge(ageStr)
        switch {
        case err != nil:
            fmt.Printf("❌ '%s': %v\n", ageStr, err)
        default:
            fmt.Printf("✅ '%s': valid age %d\n", ageStr, age)
        }
    }

	fmt.Println("\n=== Panic and Recover (Use Sparingly!) ===")

	// Demonstrate panic and recover
    func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from panic: %v\n", r)
            }
        }()
        
        fmt.Println("About to panic...")
        panic("Something went terribly wrong!")
        fmt.Println("This will never print")
    }()
    
    fmt.Println("Program continues after recovery")
}