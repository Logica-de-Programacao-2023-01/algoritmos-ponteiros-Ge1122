package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}
func main() {
	livro := Book{
		Title:  "HP",
		Author: "J. K. Rowling",
		Price:  100,
	}
	fmt.Println(livro)

	tenPercentDiscount(&livro)
	fmt.Print(livro)
}
func tenPercentDiscount(livro *Book) {
	livro.Price = livro.Price - (livro.Price * 0.1)
}

