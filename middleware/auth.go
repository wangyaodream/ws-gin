package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "valid-token" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "unauthorized",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}
