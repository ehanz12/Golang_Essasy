package main

import "fmt"

//pointer method in golang
//struct method struct function
type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Married " + m.Name
}

func main(){
	man := Man{Name: "John"}
	man.Married()
	fmt.Println(man.Name)
}