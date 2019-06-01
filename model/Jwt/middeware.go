package Jwt

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//VerifyJWToken ...
func VerifyJWToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token can not be null",
		})
		return
	}
	rawToken := string(token[len("Tin "):])
	userID, _, err := VerificationToken(rawToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token is in valid",
		})
		return
	}
	fmt.Println(userID)
	c.Set("user_id", userID)
}
