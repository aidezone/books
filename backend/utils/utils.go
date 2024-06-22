package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "fmt"
)

// 定义声明结构体
type Claims struct {
    UserID   uint   `json:"user_id"`
    Username string `json:"username"`
    Role     Role `json:"role"`
    jwt.StandardClaims
}

var jwtSecret = []byte("books-system-jwt-token")

func GenerateToken(userID uint, role Role) (string, error) {
    // 设置 JWT 过期时间
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID:  userID,
        Role:    role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    // 解析 token
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, fmt.Errorf("invalid token signature")
        }
        return nil, fmt.Errorf("invalid token")
    }
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    return claims, nil

}
