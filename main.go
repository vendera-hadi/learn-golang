package main

import (
	"net/http"

	"example/learn-golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8081")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "pong"})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	bookRepo := controllers.New()
	r.GET("/ping", ping)
	r.GET("/books", bookRepo.GetBooks)
	r.GET("/books/:id", bookRepo.GetBook)
	r.POST("/books", bookRepo.CreateBook)
	r.DELETE("/books/:id", bookRepo.DeleteBook)
	r.PUT("/books/:id", bookRepo.UpdateBook)
	r.PATCH("/checkout/:id", bookRepo.CheckoutBook)
	r.PATCH("/return/:id", bookRepo.ReturnBook)

	return r
}
