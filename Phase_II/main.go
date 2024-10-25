package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Niyonkuru", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "NiyonkuruII", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "NiyonkuruIII", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBooksById(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
func getBooksById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not fond")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("books/:id", bookById)
	router.Run("localhost:3000")
}
