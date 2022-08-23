package books

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/pkg/common/models"
	"net/http"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.Status(http.StatusOK)
}
