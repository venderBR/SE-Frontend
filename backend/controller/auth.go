package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut66/team15/entity"
	"github.com/sut66/team15/service"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email    string `json:"email"` //set varible email
	Password string `json:"password"`
}

// logintoken response
type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// get info from user email and password
func Login(c *gin.Context) {
	var payload LoginPayload
	var member entity.Member

	if error := c.ShouldBindJSON(&payload); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}

	// find user from email
	if error := entity.DB().Raw("SELECT * FROM members WHERE email = ?", payload.Email).Scan(&member).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error})
		return
	}

	//check password
	err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password incorrect"})
		return
	}

	//format token
	jwtWrapper := service.JwtWrapper{
		SecretKey:       "ABC",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(member.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error generating token"})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
		ID:    member.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})

}
