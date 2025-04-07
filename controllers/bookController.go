package controllers

import (
	"net/http"

	"github.com/adasarpan404/bookstoreapi/database"
	"github.com/adasarpan404/bookstoreapi/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := database.DB.Delete(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
