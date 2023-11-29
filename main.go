package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the GoLang Intro Live Stream!")
	fmt.Println("---------------------------------------")

	// Variables and Data Types
	fmt.Println("\n1. Variables and Data Types")
	// In Go, you declare variables using the 'var' keyword.
	var name string // Declare a string variable named 'name'.
	name = "John"   // Assign a value to 'name'.
	fmt.Println("Name:", name)

	// You can also declare and assign in one line.
	age := 30 // Shorthand declaration and assignment for an integer variable.
	fmt.Println("Age:", age)

	// Go supports various data types, such as int, float64, string, bool, etc.
	var isStudent bool = true
	fmt.Println("Is Student:", isStudent)

	// Constants
	fmt.Println("\n2. Constants")
	const pi float64 = 3.14159265359
	fmt.Println("Value of Pi:", pi)

	// Control Structures
	fmt.Println("\n3. Control Structures")
	// Conditional statements (if-else)
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// Loops (for)
	fmt.Println("\n4. Loops")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i)
	}

	// Functions
	fmt.Println("\n5. Functions")
	greeting := greet("Alice")
	fmt.Println(greeting)

	// Arrays and Slices
	fmt.Println("\n6. Arrays and Slices")
	// Arrays have a fixed size, whereas slices are dynamic.
	numbers := [5]int{1, 2, 3, 4, 5} // Declare an array with 5 integers.
	fmt.Println("Array:", numbers)
	slice := numbers[1:4] // Create a slice from the array (indices 1 to 3).
	fmt.Println("Slice:", slice)

	// Maps
	fmt.Println("\n7. Maps")
	// Maps store key-value pairs.
	studentGrades := make(map[string]int)
	studentGrades["Alice"] = 95
	studentGrades["Bob"] = 87
	fmt.Println("Alice's Grade:", studentGrades["Alice"])
	fmt.Println("Bob's Grade:", studentGrades["Bob"])

	// Structs
	fmt.Println("\n8. Structs")
	// Structs allow you to create custom data types.
	type Person struct {
		Name    string
		Age     int
		IsHuman bool
	}
	person := Person{Name: "Charlie", Age: 25, IsHuman: true}
	fmt.Println("Person:", person)

	// Error Handling
	fmt.Println("\n9. Error Handling")
	// Go uses errors for handling exceptional situations.
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Defer and Panic
	fmt.Println("\n10. Defer and Panic")
	// Defer is used to schedule a function call to run after the surrounding function returns.
	defer fmt.Println("Deferred statement")

	// Here's a scenario where panic is useful. Imagine a function that should never receive a negative input.
	// If it does, we can use panic to indicate a critical error.
	value := -5
	if value < 0 {
		fmt.Println("Received a negative value. This is a critical error!")
		panic("Critical error occurred!")
	}

	// Defer can also be used for clean-up operations like closing a file.
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return // Exit the program if there's an error.
	}
	defer file.Close() // This will close the file when main() exits.

	fmt.Println("Regular statement")

	// Uncomment the following line to see a deferred statement executed even after a panic.
	// defer fmt.Println("Another deferred statement")

	fmt.Println("\n-----------------")
	fmt.Println("End of GoLang Intro")
}

// greet is a simple function that returns a greeting message.
func greet(name string) string {
	return "Hello, " + name + "!"
}

// divide is a function that performs division and handles division by zero error.
func divide(a, b float64) (float64, error) {
	// Check if the divisor (b) is zero.
	if b == 0 {
		// If it is, return an error with a descriptive message.
		// The 'fmt.Errorf' function creates an error value with the given message.
		return 0, fmt.Errorf("division by zero")
	}

	// If the divisor is not zero, perform the division and return the result.
	return a / b, nil
}

