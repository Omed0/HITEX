package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func CheckToken(c *gin.Context) {
	
	
	secretKey := os.Getenv("SECRET_KEY")

	// Get the token from the request header
	token := c.GetHeader("Authorization")

	begin_of_token := strings.ToLower(token[0:7])

	if  begin_of_token != "bearer " {
		c.JSON(400, gin.H{
			"messege": "Authorization header format must be 'Bearer [token]'",
		})
		return
		
	}

	// Remove the "bearer" keyword from the token string
	 token = token[7:]

	secret := []byte(secretKey)

	// Parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	// Check if the token is valid
	if err != nil || !parsedToken.Valid {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the token is expired
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || time.Now().Unix() > int64(claims["exp"].(float64)) {
		c.AbortWithStatusJSON(401, gin.H{"error": "Token expired"})
		return
	}

	c.Next()
 
}
