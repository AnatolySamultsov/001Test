# print only even numbers from 0 to 20
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
