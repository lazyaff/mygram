package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = "shinigami"

func GenerateToken(id uint, email string) (string) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secretKey))
	return signedToken
}

func ValidateToken(ctx *gin.Context) (interface{}, error) {
    // Verifikasi apakah ada token
    headerToken := ctx.Request.Header["Authorization"]
    if len(headerToken) == 0 {
        return nil, errors.New("token tidak ditemukan")
    }

    tokenString := strings.TrimSpace(headerToken[0])
    if !strings.HasPrefix(tokenString, "Bearer ") {
        return nil, errors.New("format token tidak valid")
    }

    // Parse token
    stringToken := strings.Split(tokenString, " ")[1]

    token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("token tidak valid")
        }
        return []byte(secretKey), nil
    })
    if err != nil {
        return nil, errors.New("token tidak valid")
    }

    // Verifikasi claims token
    if !token.Valid {
        return nil, errors.New("token tidak valid")
    }

	// Mengembalikan id user
	claims, ok := token.Claims.(jwt.MapClaims)["id"]
	if !ok {
		return nil, errors.New("id user tidak ditemukan")
	}

    return claims, nil
}