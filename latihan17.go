package main 
import "fmt"

func sayHallo(FirstName string, Age int16){
	for i:= 0; i < 10; i++ {
		fmt.Println("Hello Word ", FirstName, "You are ", Age, " years old")
	}
}

func main() {
	sayHallo("Azka Kontol", 19)
}

