package main

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
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}