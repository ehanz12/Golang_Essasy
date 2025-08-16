// type data map in golang
package main
import "fmt"

func main() {
	// declare and initialize a map
	// var m map[string]int
	// m = make(map[string]int)
	// m["apple"] = 5
	// m["banana"] = 10
	// m["cherry"] = 15

	m := map[string]int{
		"orange": 20,
		"grape":  25,
		"kiwi":   30,
	}

	// print the map
	fmt.Println("Map elements:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// accessing a value by key
	fmt.Println("Value for 'grape':", m["grape"])

	// adding a new key-value pair
	m["melon"] = 35
	fmt.Println("After adding 'melon':", m)

	// deleting a key-value pair
	delete(m, "kiwi")
	fmt.Println("After deleting 'kiwi':", m)

	// checking if a key exists
	if val, exists := m["apple"]; exists {
		fmt.Println("Value for 'apple':", val)
	} else {
		fmt.Println("'apple' does not exist in the map")
	}

	var makeMap = make(map[string]int) // creating an empty map
	makeMap["peach"] = 40
	fmt.Println("Newly created map:", makeMap)
	// copying a map
	copyMap := make(map[string]int)
	fmt.Println("Copied map elements:", copyMap)
	// maps are unordered collections of key-value pairs
	fmt.Println("Original map after modifications:", m)
	fmt.Println("Newly created map after copying:", makeMap)
	fmt.Println("Copied map after copying:", copyMap)
}