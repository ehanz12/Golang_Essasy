package main

import "fmt"

func random() interface{} {
	
	return "reihte"
}

func main() {
	result := random()
// 	resultString := result.(string)
// 	fmt.Println(resultString) // "oke"

// 	resultInt := result.(int) //panic
// 	fmt.Println(resultInt)
	switch v := result.(type) {
	case string:
		fmt.Println(v, "string") // "oke"
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println("Tipe tidak dikenal")
	}

}