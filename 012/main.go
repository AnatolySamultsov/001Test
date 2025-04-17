package main

import "fmt"

type bookStruct struct {
	Title     string
	Author    string
	Data_type string
}

func main() {

	book1 := bookStruct{
		Title:     "The Great Gatsby",
		Author:    "F. Scott Fitzgerald",
		Data_type: "1925",
	}

	fmt.Println("Book Information:")
	fmt.Println("Title:", book1.Title)
	fmt.Println("Author:", book1.Author)
	fmt.Println("Data type:", book1.Data_type)
}
