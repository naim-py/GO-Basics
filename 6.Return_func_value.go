package main

import "fmt"

// Function to add two numbers
func sum(n1 int, n2 int) int {
    return n1 + n2
}

func main() {
    var a, b int

    // Taking input from the user
    fmt.Print("Enter the first number: ")
    fmt.Scanln(&a)

    fmt.Print("Enter the second number: ") // Fixed the message
    fmt.Scanln(&b)

    // Calling the sum function
    result := sum(a, b)

    // Printing the result
    fmt.Println("Sum:", result)
}
