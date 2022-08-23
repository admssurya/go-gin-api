package books

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/pkg/common/models"
	"net/http"
)

type UpdateBookRequestBody struct {
	Title       string `json:"string"`
	Author      string `json:"string"`
	Description string `json:"string"`
}

func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := UpdateBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
