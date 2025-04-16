# 005Test
package main

import(
"bufio"
"fmt"
"os"
)

func main() {
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter your full name")
name, _ := reader.ReadString('\n') //includes newline character
fmt.Println("Your name is", name)
// fmt.Print(reflect.Type0f(name))


