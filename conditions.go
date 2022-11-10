package main

import "fmt"

func main() {
	var x = 6
	var y = 5
	// If / Else If / Else
	if x == 5 && y == 3 {
		fmt.Print("x==5 and y==3")
	} else if x <= 5 || y < 3 {
		fmt.Print("x<=5 or y<3")
	} else {
		fmt.Print("Whoops!")
	}

	// If with short declaration
	if n := 10; n%2 == 0 {
		fmt.Printf("\n%d is even\n", n)
	}

	// Switch
	var n = 3
	switch n {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("Blah")
	}

	// Switch with declaration
	switch n := 2; n {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("Blah")
	}

	// Switch with same cases
	switch n := 2; n {
	case 1, 2:
		fmt.Println("A")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("Blah")
	}

	// Switch with no expression
	switch {
	case n == 1:
		fmt.Println("A")
	case n == 2:
		fmt.Println("B")
	case n == 3:
		fmt.Println("C")
	default:
		fmt.Println("Blah")
	}
}
