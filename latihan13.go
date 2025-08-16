//if stacment 
package main
import "fmt"

func main() {
	name := "budi"

	if name == "eko" {
		fmt.Println("Hello Eko")
	} else if name == "budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello Unknown")
	}

	if length := len(name); length > 3 {
		fmt.Println("Name is longer than 3 characters")
	} else {
		fmt.Println("Name is 3 characters or shorter")
	}

	

}