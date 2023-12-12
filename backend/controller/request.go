package controller

import (

	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/entity"

)

// POST /request check
func CreateBooks_request(c *gin.Context) {
	var bookRequest entity.Book_request
	// var status entity.Book_request_status
	// var member entity.Member
	// var admin entity.Admin
	// var category entity.Catagory

	// Bind JSON payload to bookRequest struct
	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create BookRequest instance
	newBookRequest := entity.Book_request{
		Book_request_status_id:	bookRequest.Book_request_status_id,
		Member_id:   			bookRequest.Member_id, 
		Catagory_id:   			bookRequest.Catagory_id,
		Admin_id: 				bookRequest.Admin_id,
		BookRequest_title:   	bookRequest.BookRequest_title,     
		BookRequest_reason:    	bookRequest.BookRequest_reason,      
		BookRequest_writer:     bookRequest.BookRequest_writer,
		BookRequest_publisher:  bookRequest.BookRequest_publisher,
	}
	

	// Create the record in the database
	if err := entity.DB().Create(&newBookRequest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newBookRequest})
}


// GET by id  check
func GetBooks_request(c *gin.Context) {
	var book entity.Book_request
	id := c.Param("id")
	if err := entity.DB().Preload("Book_request_status").Preload("Member").Preload("Admin").Preload("Catagory").Raw("SELECT * FROM book_requests WHERE id = ?", id).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET  check
func ListBooks_request(c *gin.Context) {
	var book []entity.Book_request
	if err := entity.DB().Preload("Book_request_status").Preload("Member").Preload("Admin").Preload("Catagory").Raw("SELECT * FROM book_requests").Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE by id  check
func DeleteBooks_request(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_requests WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateBooks_request(c *gin.Context) {
	var book entity.Book_request
	var result entity.Book_request

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา book ด้วย id
	if tx := entity.DB().Where("id = ?", book.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if err := entity.DB().Save(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}