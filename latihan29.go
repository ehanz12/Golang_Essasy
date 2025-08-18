package main

import "fmt"

func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("John Doe")
	person2 := NewMap("")
	fmt.Println(person2) // This will print: map[]
	fmt.Println(person)
}