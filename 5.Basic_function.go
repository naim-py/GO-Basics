package main
import "fmt"

func sum(n1 int,n2 int){
	add:=n1+n2
	fmt.Println("Sum: ",add)
}

func main(){
	a:=13
	b:=20
	sum(a,b)
}

