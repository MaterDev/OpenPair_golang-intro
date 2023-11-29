package main

import "fmt"

// main function - entry point of every Go program
func main() {

    // Basic printing in Go
    fmt.Println("Hello, GoLang Live Stream!")

    // Variables in Go
    // Declare a variable with 'var'
    var name string = "GoLang Basics"
    fmt.Println("Topic:", name)

    // Short variable declaration
    // Type is inferred by the compiler
    count := 10
    fmt.Println("Count:", count)

    // Basic data types in Go: int, float64, bool, string
    var age int = 30
    var accuracy float64 = 99.9
    var isValid bool = true
    var greeting string = "Hello, viewers!"

    // Printing variables with different data types
    fmt.Printf("Age: %d, Accuracy: %.2f, Valid: %t, Greeting: %s\n", age, accuracy, isValid, greeting)

    // If-else statement in Go
    if count > 5 {
        fmt.Println("Count is greater than 5")
    } else {
        fmt.Println("Count is less than or equal to 5")
    }

    // Loops in Go - using for
    // There's no while loop in Go, for is used instead
    for i := 0; i < 5; i++ {
        fmt.Println("Loop iteration:", i)
    }

    // Functions in Go
    // Here's a simple function to add two numbers
    result := addNumbers(2, 3)
    fmt.Println("2 + 3 =", result)
}

// Function declaration in Go
// func functionName(param1 type, param2 type) returnType
func addNumbers(a int, b int) int {
    return a + b
}
