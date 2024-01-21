package secrets

import (
	"time"

	"github.com/endot1231/ec-backend/configs"
	"github.com/golang-jwt/jwt/v5"
)

func JwtGenerate(role string, userId string, now time.Time) (string, error) {
	claims := jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(now.Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(now),
		"NotBefore": jwt.NewNumericDate(now),
		"Issuer":    "endot1231",
		"Role":      role,
		"ID":        userId,
	}

	mySigningKey := []byte(configs.Get().JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}
