package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Description string
	AuthorID int
}

func main() {

	// Create a new book
	book := Book{
		Title: "The Go Programming Language",
		Description: "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go.",
		AuthorID: 1,
	}

	fmt.Println("Book", book)

	bookJson, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}



	fmt.Println("Book JSON", string(bookJson))


	// Unmarshal the JSON back to a book
	var book2 Book
	err = json.Unmarshal(bookJson, &book2)

	fmt.Println("Book 2", book2.Title)
	

}