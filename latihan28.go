package main

import "fmt"


func Ups(name string) interface{} {
	if name == "Hello World" {
		return 1
	} else if name == "Hello" {
		return 2.5
	} else {
		return true
	}
}

func main() {
	result := Ups("Hello World")
	fmt.Println(result)
}