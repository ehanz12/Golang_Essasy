//switch case in golang

package main
import "fmt"

func main() {
	name := "reihan"

	switch name {
	case "eko":
		fmt.Println(("Hello Eko"))
	case "budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hello Unknown")
	}

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("a Name is longer than 3 characters")
	case false:
		fmt.Println("a Name is 3 characters or shorter")
	}
}