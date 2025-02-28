package main

import "fmt"

func main() {
    grade := "B"

    switch grade {
    case "A":
        fmt.Println("Excellent!")
    case "B", "C":
        fmt.Println("Good job!")
    case "D":
        fmt.Println("You passed.")
    case "F":
        fmt.Println("Better luck next time.")
    default:
        fmt.Println("Invalid grade")
    }

    // Using fallthrough to continue execution to the next case
    num := 2
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three (executed due to fallthrough)")
    default:
        fmt.Println("Number not found")
    }
}
