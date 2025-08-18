//type declarations 
package main
import "fmt"

// This program demonstrates the use of variables in Go.
func main() {
	// Declare and initialize a variable
	type name string
	var myName name = "Reihan Aditya Putra X RPL 2"
	fmt.Println("Nama: ", myName) // Reihan Aditya Putra X RPL 2
	// Declare a new type based on int
	type age int
	var myAge age = 16
	fmt.Println("Umur: ", myAge) // 16
}