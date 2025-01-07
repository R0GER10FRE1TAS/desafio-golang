package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Synopsis string `json:"synopsis"`
}

func connectDB() {
	var err error
	dsn := "host=localhost user=postgres password=087081 dbname=desafio_golang port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	fmt.Println("Database connection successfully established")

	db.AutoMigrate(&Book{})
}

func main() {
	connectDB()
	r := gin.Default()

	r.POST("/books", createBook)

	r.GET("/books", listBooks)

	r.GET("/books/:id", listOneBook)

	r.PUT("/books/:id", updateBook)

	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")
}
func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data"})
		return
	}
	if err := db.Create(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": "Book's record failed"})
		return
	}
	c.JSON(201, gin.H{"message": "Book registered successfully"})
}
func listBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(500, gin.H{"error": "Cannot list the books"})
		return
	}
	c.JSON(200, books)
}
func listOneBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, book)
}
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	var updatedData Book
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data"})
		return
	}
	if updatedData.Title != "" {
		book.Title = updatedData.Title
	}
	if updatedData.Category != "" {
		book.Category = updatedData.Category
	}
	if updatedData.Author != "" {
		book.Author = updatedData.Author
	}
	if updatedData.Synopsis != "" {
		book.Synopsis = updatedData.Synopsis
	}

	if err := db.Save(&book).Error; err != nil {
		c.JSON(500, gin.H{"error": "Book record update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Book record updated successfully"})
}
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Book{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete book record"})
		return
	}
	c.JSON(200, gin.H{"message": "Book record deleted successfully"})
}
