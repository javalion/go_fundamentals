package main

import "fmt"

func main() {
	// Basic For Loop
	fmt.Println("Basic For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Basic While Loop (using for construct)
	fmt.Println("Basic While Loop")
	n := 1
	for n < 5 {
		fmt.Println(n)
		n++
	}

	// Loop through array
	fmt.Println("Loop through array")
	strAry := []string{"Sunday", "Monday", "Tuesday"}
	for idx, val := range strAry {
		fmt.Println(idx, val)
	}

	// Loop through map keys
	fmt.Println("Loop through map")
	myMap := make(map[string]int)
	myMap["key1"] = 1
	myMap["key2"] = 2
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
