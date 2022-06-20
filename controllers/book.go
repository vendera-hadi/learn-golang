package controllers

import (
	"errors"
	"example/learn-golang/database"
	"example/learn-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

func New() *BookRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Book{})
	return &BookRepo{Db: db}
}

//create book
func (repository *BookRepo) CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.Create(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

//get books
func (repository *BookRepo) GetBooks(c *gin.Context) {
	var book []models.Book
	err := models.GetBooks(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

//get book by id
func (repository *BookRepo) GetBook(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var book models.Book
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// update book
func (repository *BookRepo) UpdateBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&book)
	err = models.Update(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// delete book
func (repository *BookRepo) DeleteBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.Delete(repository.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// checkout book
func (repository *BookRepo) CheckoutBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	book.Quantity -= 1
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}
	err = models.Update(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// return book
func (repository *BookRepo) ReturnBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	book.Quantity += 1
	err = models.Update(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}
