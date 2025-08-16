package main
import (
	"fmt"
)

func main() {
	counter := 0
	increment := func (){
		counter := 1
		fmt.Println("Incrementing counter")
		counter++
		fmt.Println("Counter value inside increment:", counter) // Output: Counter value inside increment: 2
	}

	increment()
	increment()
	fmt.Println("Counter value:", counter) // Output: Counter value: 2
}