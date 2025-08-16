package main
import "fmt"
// This program demonstrates the use of variables in Go.
func main() {
	var a = 10 + 10
	var f = 10
	a += 10 // This line will cause a compilation error because f is not declared
	// This line will cause a compilation error because f is not declared
	fmt.Println("Penjumlahan: ", f) // 20
	var b = 20 - 5
	var c = 5 * 4
	var d = 20 / 4
	var e = 20 % 3
	fmt.Println("Penjumlahan: ", a) // 20
	fmt.Println("Pengurangan: ", b) // 15
	fmt.Println("Perkalian: ", c) // 20
	fmt.Println("Pembagian: ", d) // 5
	fmt.Println("Sisa Bagi: ", e) // 2
}