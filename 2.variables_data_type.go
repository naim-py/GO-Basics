package main

import (
	"fmt"
	"reflect"
)

func main(){
	//1.Varible Declaration with explicit type
	var name string= "naim Hossain"
	var age int = 30
	var height float64 =5.9
	var isEmployed bool=true

	//2.Varible Declaration with type inference with :=
	country := "USA"		// Automatically inferred as string
	salary := 50000.50 		// Automatically inferred as float64

	//3.Multiple variable declaration
	var x, y, z int=10, 20, 30
	a, b, c :="Go", 3.15, false

	// 4. Constant Declaration
	const pi float64 = 3.14159
	const company = "Tech Corp"

	// 5. Zero Value Variables (default values)
	var defaultInt int      // 0
	var defaultFloat float64 // 0.0
	var defaultBool bool    // false
	var defaultString string // ""

	// Printing values with their types
	fmt.Println("Name: ",name,"Type: ",reflect.TypeOf(name))
	fmt.Println("Age:", age, "Type:", reflect.TypeOf(age))
	fmt.Println("Height:", height, "Type:", reflect.TypeOf(height))
	fmt.Println("Employed:", isEmployed, "Type:", reflect.TypeOf(isEmployed))
	fmt.Println("Country:", country, "Type:", reflect.TypeOf(country))
	fmt.Println("Salary:", salary, "Type:", reflect.TypeOf(salary))
	fmt.Println("Multiple Variables: ",x,y,z,a,b,c)
	fmt.Println("Constants:", pi, company)
	fmt.Println("Default Values:", defaultInt, defaultFloat, defaultBool, defaultString)
}