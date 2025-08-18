//conversion
package main
import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println("Nilai 32: ", nilai32) // 10000
	fmt.Println("Nilai 64: ", nilai64) // 10000
	fmt.Println("Nilai 8: ", nilai8) // -128

	var name = "Reihan Aditya Putra X RPL 2"
	var e byte = name[0] // mengambil byte pertama dari string
	fmt.Println("Byte pertama: ", e) // 82
	fmt.Println("Karakter pertama: ", string(e)) // R
}