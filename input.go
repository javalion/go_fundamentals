package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Process stdin - Reader (String - Line)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("Hi, " + text + ". How are you?")

	// Process stdin - Reader (Character)
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Continue (Y/N) >")
	char, _, err := reader2.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	switch char {
	case 'Y':
		fmt.Println("Continuing")
		break
	case 'N':
		fmt.Println("Exiting")
		break
	}

	// Process stdin - Scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println("You Typed: " + scanner.Text())
	}
}
