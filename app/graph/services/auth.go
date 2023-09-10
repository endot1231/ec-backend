package services

import (
	"context"
	"errors"
	"time"

	"github.com/endot1231/ec-backend/confings"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/users"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	exec ent.Client
}

func (u *authService) Login(ctx context.Context, role string, email string, password string) (string, error) {
	user, err := u.exec.Users.Query().Where(
		users.Email(email),
	).First(ctx)

	if user == nil {
		return "", errors.New("failed login")
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("failed login")
	}

	claims := jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"Issuer":    "endot1231",
		"Role":      "user",
		"ID":        user.ID,
	}

	mySigningKey := []byte(confings.Get().JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}
