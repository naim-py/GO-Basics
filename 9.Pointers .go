package main

import "fmt"

func main() {
    var x int = 10
    var p *int = &x  // Pointer to x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Pointer p stores address:", p)
    fmt.Println("Value stored at pointer p:", *p) // Dereferencing
}
