package main

import (
	"fmt"
)

func sum(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	slice := []int{10, 20, 30}
	fmt.Println(sum(slice...))
}