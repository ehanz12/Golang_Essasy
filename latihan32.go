package main

import "fmt"

type Kota struct {
	City, Provinsi, Kec string
}

//pass by value adalah untuk menduplikat data
// jika kita mengubah nilai dari result2, itu tidak akan mempengaruhi result
//kalo pointer by referense jadi tidak menduplikat data

func main() {
	result := Kota{"Jakarta", "DKI Jakarta", "Kecamatan Menteng"}
	// result2 := &result // by reference pointer
	result2 := result // pass by value

	result2.City = "Bandung"
	fmt.Println("Kota Asli:", result.City) // "Jakarta"
	fmt.Println("Kota Baru:", result2.City) // "Bandung"
}