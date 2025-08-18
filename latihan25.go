package main
import (
	"fmt"
)

//defer digunakan untuk memanggil function ga peduli ada error

func logging() {
	fmt.Println("Logging function called")
}

func runApplication() {
	defer logging() // This will execute after runApplication finishes
	fmt.Println("Running application")
}

//panic digunakan untuk menghentikan eksekusi program dan mencetak pesan kesalahan
func endApp() {
	message := recover() // This will catch the panic
	if message != nil {
		fmt.Println("Recovered from panic:", message)
	}
	fmt.Println("Application ended")
}

func runApp(error bool) {
	defer endApp() // This will execute at the end of runApp
	if error !=  {
		panic("An unexpected error occurred")
	}
	
	fmt.Println("Application is running")
	
	// Uncomment the next line to see panic in action
	// panic("An unexpected error occurred") // This will stop execution and call deferred functions
}

//recover untuk menangkap panic
func main() {
	runApp(true) // Change to true to see panic in action
}