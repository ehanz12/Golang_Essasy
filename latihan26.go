package main

import "fmt"

// Struct definition

type Customer struct {
	ID   int
	Name string
	Age  int
}

//struct in function method
func (customer Customer) sayHello() {
	fmt.Println("Hello,", customer.Name)
}

func main() {
	customer := Customer{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	fmt.Println("Customer ID:", customer.ID)
	fmt.Println("Customer Name:", customer.Name)
	fmt.Println("Customer Age:", customer.Age)

	reihan := Customer{Name: "Reihan Aditya Putra"}
	reihan.sayHello() // Calling the method to say hello
}