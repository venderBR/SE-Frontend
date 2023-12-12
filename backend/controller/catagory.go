package controller

import (

	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/entity"

)
// GET check
func ListCatagory(c *gin.Context) {
	var list []entity.Catagory
	if err := entity.DB().Raw("SELECT * FROM catagories").Find(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

// GET by id check
func GetCatagory(c *gin.Context) {
	var choice entity.Catagory
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM catagories WHERE id = ?", id).Find(&choice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": choice})
}