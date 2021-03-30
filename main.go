package main

import "fmt"

type Books struct {
	title  string
	author string
	bookId int
}

func main() {
	var Book1 Books

	Book1 = Books{
		title:  "123",
		author: "2",
		bookId: 1,
	}

	fmt.Println(Book1)
}