package main

import "fmt"

// Struct
type Person struct {
    Name string
    Age  int
}

// Method to print details
func (p Person) display() {
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
}

func main() {
    p1 := Person{"Alice", 25}
    p1.display()
}
