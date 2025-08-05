//variable
// This program demonstrates the use of variables in Go.
// It declares a variable, assigns a value, and prints it to the console.

package main
import "fmt"
func main() {
	//deklarasi variabel
	var name string = "Reihan Aditya Putra X RPL 2"
	var age int = 16

	//print variabel
	fmt.Println("Nama: ", name) // Reihan Aditya Putra X RPL 2
	fmt.Println("Umur: ", age)   // 16

	nilai := 90 // shorthand declaration
	boolean := true // shorthand declaration
	fmt.Println("Boolean: ", boolean) // true
	fmt.Println("Nilai: ", nilai) // 90

	//gabungan variable
	var (
		namaLengkap string = "Reihan Aditya Putra X RPL 2"
		kelas       string = "X RPL 2"
		umur        int    = 16
	)

	fmt.Println("Nama Lengkap: ", namaLengkap) // Reihan Aditya Putra X RPL 2
	fmt.Println("Kelas: ", kelas)               // X RPL 2
	fmt.Println("Umur: ", umur)                 // 16
	fmt.Println("Gabungan: ", namaLengkap, kelas, umur) // Reihan Ad
}