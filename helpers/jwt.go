package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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