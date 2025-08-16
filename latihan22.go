package main

import (
	"fmt"
)

type Blacklist func(string) bool

func HelloWithBlacklist(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Hello, guest! you are blacklisted.")
	} else {
		fmt.Printf("Hello, %s!\n", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "bad_user"
	}

	HelloWithBlacklist("Alice", blacklist)
	HelloWithBlacklist("bad_user", blacklist)
}