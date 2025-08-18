package main

import (
	"fmt"
)

func GoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodbye := GoodBye
	fmt.Println(goodbye("Reihan"))
}