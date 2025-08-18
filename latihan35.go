package main

//import in golang
// dengan underscore / _ kita mengimport package tanpa menggunakannya secara langsung
// ini berguna untuk inisialisasi atau side effect dari package tersebut

import (
	_ "Latihan/database" // Adjust the import path as necessary
)

func main() {

	// result := database.GetConnection()
	// fmt.Println(result)
}