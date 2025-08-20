package main

import (
	"flag"
	"fmt"
)

// import "fmt"

func main() {
	var host *string = flag.String("host", "localhost", "Put your host here")
	var user *string = flag.String("user", "root", "Put your user here")
	var password *string = flag.String("password", "Rhanss12345", "Put your password here")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}