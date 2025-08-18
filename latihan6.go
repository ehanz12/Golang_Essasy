package main 
import "fmt"
// This program demonstrates the use of variables in Go.
func main() {
	const name = "Reihan Aditya Putra X RPL 2"
	const age = 16
	// Print constants
	name = 100 // This line will cause a compilation error because constants cannot be reassigned
	fmt.Println("Nama: ", name) // Reihan Aditya Putra X RPL
}