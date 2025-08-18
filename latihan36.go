//mengambil data argument
//menggunakan package os

package main

import (
	"fmt"
	"os"
)

func main() {
	osArgs := os.Args
	for i, arg := range osArgs {
		fmt.Printf("Argument %d: %s\n", i, arg)
	}

	fmt.Println("Total arguments:",osArgs)

	osHost, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	fmt.Println("Hostname:", osHost)
}