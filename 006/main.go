// File: 006/main.go
// ## This program prints the even numbers from 0 to 20, inclusive, and then prints "End of program"
package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("End of program")
}
