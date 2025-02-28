package main

import "fmt"

func main() {
    // Array
    var arr = [3]int{10, 20, 30}
    fmt.Println("Array:", arr)

    // Slice
    slice := []string{"Go", "Python", "Java"}
    slice = append(slice, "C++") // Add an element
    fmt.Println("Slice:", slice)
}
