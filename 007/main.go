// File: 007/main.go
// ## This program prints the uppercase letters A-Z and lowercase letters a-z

package main

import (
	"fmt"
)

func main() {
	// Uppercase A-Z
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// Lowercase a-z
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
