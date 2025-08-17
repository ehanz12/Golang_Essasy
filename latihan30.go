package main

import (
	"fmt"
	"errors"
)

func Pembagian(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("Pembagian dengan nol tidak diperbolehkan")
	}else {
		result := i / j
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println("Hasil:", hasil)
	}


	// var contoh error = errors.New("Contoh error")
	// fmt.Println("Error:", contoh)
}