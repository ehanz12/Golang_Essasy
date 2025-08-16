//tipe data number golang
//ada 3 tipe data number di golang yaitu integer, unsigned integer, dan float
//int8 = -128 to 127
//int16 = -32768 to 32767
//int32 = -2147483648 to 2147483647
//int64 = -9223372036854775808 to 9223372036854775807

//jadi kita gunakan tipe data integer / int untuk menyimpan bilangan bulat

//untuk yang ke dua
//uint8 = 0 to 255
//uint16 = 0 to 65535
//uint32 = 0 to 4294967295
//uint64 = 0 to 18446744073709551615

//jadi kita gunakan tipe data unsigned integer / uint untuk menyimpan bilangan bulat positif

//untuk yang ke tiga
//float32 = 1.401298464324817e-45 to 3.4028234663852886e+38
//float64 = 5.0e-324 to 1.7976931348623157e+308
//complex64 = complex numbers with float32 real and imaginary parts
//complex128 = complex numbers with float64 real and imaginary parts

//jadi kita gunakan tipe data float untuk menyimpan bilangan pecahan

//ALIAS / nama lain dari tipe data

//int = Minimal int32
//uint = Minimal uint32
//byte = uint8
//rune = int32

//contoh penggunaan tipe data number
package main
import "fmt"

func main() {
	//tipe data integer
	fmt.Println("Tipe data integer: ", 123) // int
	fmt.Println("Tipe data integer: ", -123) // int
	fmt.Println("Tipe data float: ", 1.2) // float
}