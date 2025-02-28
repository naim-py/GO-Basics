package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go hello() // Runs concurrently
    time.Sleep(time.Second) // Wait to allow Goroutine to finish
}
