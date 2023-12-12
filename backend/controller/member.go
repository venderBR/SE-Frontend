package controller

import (
	"net/http"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/entity"
	"golang.org/x/crypto/bcrypt"
)

// POST /member
func CreateMember(c *gin.Context) {
	var member entity.Member
	// var gender entity.Gender
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//hashpassword
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hash password"})
		return
	}

	// สร้าง User
	m := entity.Member{
		Gender_id:   member.Gender_id, // โยงความสัมพันธ์กับ Entity Gender
		Prefix_id:   member.Prefix_id,
		Religion_id: member.Religion_id,
		FirstName:   member.FirstName,     // ตั้งค่าฟิลด์ FirstName
		LastName:    member.LastName,      // ตั้งค่าฟิลด์ LastName
		Email:       member.Email,         // ตั้งค่าฟิลด์ Email
		Password:    string(hashPassword), // ตั้งค่าฟิลด์ Password
		Tel:         member.Tel,           // ตั้งค่าฟิลด์ Phone
		Age:         member.Age,           // ตั้งค่าฟิลด์ Profile
	}

	// var existingUser entity.Member
	// if err := entity.DB().Where("user_email = ?", user.User_email).First(&existingUser).Error; err == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "อีเมลนี้มีผู้ใช้แล้ว"})
	// 	return
	// }

	if err := entity.DB().Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": m})
}

// GET /members
func ListMembers(c *gin.Context) {
	var members []entity.Member
	if err := entity.DB().Preload("Gender").Preload("Religion").Preload("Prefix").Raw("SELECT * FROM members").Find(&members).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": members})
}

func GetMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	if err := entity.DB().Preload("Gender").Preload("Religion").Preload("Prefix").Raw("SELECT * FROM members WHERE id = ?", id).Find(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

//
