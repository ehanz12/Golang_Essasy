// array in golang 
package main
import "fmt"

func main() {
	// declare and initialize an array
	// var arr [3]string 
	// arr[0] = "apple"
	// arr[1] = "banana"
	// arr[2] = "cherry"

	 values := [...]string{
		"orange",
		"grape",
		"kiwi",
		"melon",
		}
	// print the array
	fmt.Println("Array elements:")
	//untuk mencetak seluruh elemen array
	//atau bisa dibilang juga looping
	for i := 0; i < len(values); i++ {
		fmt.Printf("Element at index %d: %s\n", i, values[i])
	}

	//slice adalah potongan array
	var slice = values[1:]
	fmt.Println(len(slice)) // panjang slice
	fmt.Println(cap(slice)) // slice
	fmt.Println("Slice elements:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Element at index %d: %s\n", i, slice[i])
	}

	slice2 := values[2:4]
	fmt.Println("Slice2 elements:", slice2)
	slice3 := append(slice2, "papaya")
	fmt.Println("Slice3 elements:", slice3)
	slice3[1] = "blueberry" // mengubah elemen kedua
	fmt.Println("Updated Slice3 elements:", slice3)
	fmt.Println("Updated values array:", values) // menampilkan array yang telah diubah

	//untuk membuat slice baru
	newSlice := make([]string, 2, 5) // membuat slice baru dengan panjang 2
	newSlice[0] = "peach"
	newSlice[1] = "plum"
	fmt.Println("New Slice elements:", newSlice)

	//untuk mencopy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Copied Slice elements:", copySlice)

	//perbedaan antara array dan slice
	// array memiliki ukuran tetap, sedangkan slice dapat berubah ukurannya
	
}