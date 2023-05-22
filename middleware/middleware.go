package middleware

import (
	"fmt"
	"net/http"

	token "github.com/pomntv/Neversitup_E_Commerce/tokens"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ClientToken := c.Request.Header.Get("token")
		ClientToken := c.Request.Header.Get("Authorization")
		fmt.Println("c.Request.Header :", c.Request.Header)
		fmt.Println("ClientToken :", ClientToken)

		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		fmt.Printf("err: %v\n", err)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
