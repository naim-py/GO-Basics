package main
import "fmt"

func main(){
	age:=22
	if age >=18{
		fmt.Println("He is adult: ")
	}else if age>=11{
		fmt.Println("He is teenager.")
	}else{
		fmt.Println("You are a child")
	}
	if x := 10; x > 5 {
        fmt.Println("x is greater than 5")
    }

	//switch statements
	day:=3
	switch day {
	case 1:
		fmt.Println("This is saturday")
	case 2:
		fmt.Println("This is Sunday")
	case 3:
		fmt.Println("This is Monday")
	case 4:
		fmt.Println("This is Tuesday")
	case 5:
		fmt.Println("This is Wednesday")
	case 6:
		fmt.Println("This is Thusday")
	case 7:
		fmt.Println("This is Friday")
	default:
		fmt.Println("Invalid day")
	}
}