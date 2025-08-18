package main

import (
	"fmt"
)
func FactorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func FactorialRecursive(value int) int {
	if value == 0 {
		return 1
	}
	return value * FactorialRecursive(value-1)
}

func main() {
	fmt.Println(FactorialLoop(0)) // Output: 1
	fmt.Println(FactorialRecursive(0)) // Output: 1
	fmt.Println(FactorialLoop(5)) // Output: 120
	fmt.Println(FactorialRecursive(5)) // Output: 120
}