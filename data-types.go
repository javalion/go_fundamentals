package main

import "fmt"

func main() {
	var firstName = "John"
	fmt.Println("String: ", firstName)
	var age = 30
	fmt.Println("Integer: ", age)
	var price = 19.99
	fmt.Println("Float: ", price)
	var isHappy = true
	fmt.Println("Boolean: ", isHappy)
	var myArray = [3]int{1, 2, 3}
	fmt.Println("Array: ", myArray)
	type person struct {
		name string
		age  int
	}
	me := person{name: "Mary", age: 42}
	fmt.Println("Struct:", me)
	m := make(map[int]string)
	m[1] = "AAA"
	m[2] = "BBB"
	fmt.Println("Map:", m)
}
