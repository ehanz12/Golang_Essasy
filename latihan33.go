package main

import "fmt"

//Pointer in function parameter
//pass by reference

type Mahasiswa struct {
	Nama, Kelas, Alamat string
}

func ChangeAlamat(m *Mahasiswa) { //ini pake pointer agar bisa mengubah nilai asli
	m.Alamat = "Jalan Baru"
}

// func ChangeNama(m Mahasiswa) {  //ini pass by value
// 	m.Nama = "Reihan Aditya Putra"
// }

func main() {
	address := Mahasiswa{"Reihan", "X RPL 2", "Jalan Lama"}
	fmt.Println("Sebelum perubahan:")
	fmt.Println(address.Alamat)
	ChangeAlamat(&address)
	fmt.Println("Setelah perubahan:")
	fmt.Println(address.Alamat)
}