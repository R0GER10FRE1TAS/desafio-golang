package main


import(
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID		uint	`json:"id"`
	Title	string	`json:"title"`
	Category string	`json:"category"`
	Author	string	`json:"author"`
	Synopsis string	`json:"synopsis"`
}

var books []Book

func main() {
	r := gin.Default()

	r.POST("/books", func(c *gin.Context){
		var newBook Book
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		newBook.ID = uint(len(books) + 1)
		books = append(books, newBook)
		
		c.JSON(http.StatusCreated, newBook)
	})

	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.GET("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 || id > len(books) + 1{
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, books[id-1])
	})

	r.PUT("/books/:id", func(c *gin.Context) {
		id, err :=  strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 || id > len(books) + 1{
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		var updatedFields Book
		if err := c.ShouldBindJSON(&updatedFields); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, book := range books {
			if book.ID == uint(id) {
				// Atualiza apenas os campos enviados
				if updatedFields.Title != "" {
					books[i].Title = updatedFields.Title
				}
				if updatedFields.Category != "" {
					books[i].Category = updatedFields.Category
				}
				if updatedFields.Author != "" {
					books[i].Author = updatedFields.Author
				}
				if updatedFields.Synopsis != "" {
					books[i].Synopsis = updatedFields.Synopsis
				}
	
				c.JSON(http.StatusOK, books[i])
				return
			}
		}
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 || id > len(books) + 1{
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		books = append(books[:id-1], books[id:]...)
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})

	r.Run()
}