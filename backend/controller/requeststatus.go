package controller

import (

	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/entity"

)
// GET check
func ListBooks_requeststatus(c *gin.Context) {
	var status []entity.Book_request_status
	if err := entity.DB().Raw("SELECT * FROM book_request_statuses").Find(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET by id check
func GetBooks_requeststatus(c *gin.Context) {
	var status entity.Book_request_status
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_request_statuses WHERE id = ?", id).Find(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}