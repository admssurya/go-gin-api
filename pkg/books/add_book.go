package books

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/pkg/common/models"
	"net/http"
)

type AddBookRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h handler) AddBook(c *gin.Context) {
	body := AddBookRequestBody{}

	// getting request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
