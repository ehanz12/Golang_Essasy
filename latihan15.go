package main
import "fmt"

func main() {
	count := 1

	for count < 10 {
		if count%2 == 0 {
			fmt.Println(count, "is even")
		} else {
			fmt.Println(count, "is odd")
		}
		count++
	}

	//for dengan stacment
	fmt.Println("Using for loop with a condition:")

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	//for dengan range
	fmt.Println("Using for loop with range:")
	numbers := []int {1, 2, 3, 4, 5}
	for index, value := range numbers {
		if value%2 == 0 {
			fmt.Println("Index:", index, "Value:", value, "is even")
		} else {
			fmt.Println("Index:", index, "Value:", value, "is odd")
		}
	}
	//for dengan kondisi
	fmt.Println("Using for loop with a condition and out of range check:")
	for i := 0; i <= len(numbers); i++ {
		if i < len(numbers) {
			fmt.Println("Index:", i, "Value:", numbers[i])
		} else {
			fmt.Println("Index:", i, "is out of range")
		}
	}

	//for map
	fmt.Println("Using for loop with a map:")
	
	person := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range person {
		if age%2 == 0 {
			fmt.Println(name, "is", age, "years old and even")
		} else {
			fmt.Println(name, "is", age, "years old and odd")
		}
	}

	sayHallo()
}