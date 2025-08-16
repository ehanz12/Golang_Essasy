package main

import (
	"fmt"
)
//type declaration
type FilterFunc func(string) string


func HelloWithFilter(name string, filter FilterFunc) string {
	return "Hello " + filter(name)
}

func spam(name string) string {
	if name == "Reihan" {
		return "Spam detected!"
	}
	return "No unknown"
}

func main() {
	fmt.Println(HelloWithFilter("Anjing", spam))
}