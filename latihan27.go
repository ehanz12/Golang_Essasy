package main
import "fmt"

//interface in golang 

type HasName interface {
	GetName() string
}

type Customer2 struct {
	Name string
}


func Say(h HasName) {
	fmt.Println("Name:", h.GetName())
}

func (c Customer2) GetName() string {
	return c.Name
}
func main() {
	customer := Customer2{Name: "John Doe"}
	Say(customer)
}
