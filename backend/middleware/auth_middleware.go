package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    // "github.com/dgrijalva/jwt-go"
    "books-manage-system/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        token = strings.TrimPrefix(token, "Bearer ")
        user, err := utils.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    }
}

func AuthAdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, ok := c.Get("user")
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid context"})
            c.Abort()
        }
        if user.(*utils.Claims).Role != utils.Admin {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
            c.Abort()
            return
        }
        c.Next()
    }
}
