package main

import "fmt"

// Function to print a welcome message
func printWelcome() {
    fmt.Println("Welcome to the Go language.")
}

// Function to get the user's name
func getUserName() string {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    return name
}

// Function to get two numbers from the user
func getTwoNum() (int, int) { // Fixed return type
    var num1, num2 int
    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)
    return num1, num2 // Added return statement
}

// Function to add two numbers
func add(n1 int, n2 int) int {
    return n1 + n2
}

// Function to display the result
func display(name string, sum int) {
    fmt.Println("Hello", name)
    fmt.Println("Summation =", sum)
}

func main() {
    printWelcome()
    name := getUserName()
    n1, n2 := getTwoNum()
    sum := add(n1, n2)
    display(name, sum)
}
