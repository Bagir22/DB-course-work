package utils

import (
	"courseWork/internal/types"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateUser(user types.UserLongData) bool {
	if user.FirstName == "" || user.LastName == "" ||
		user.Email == "" || user.Phone == "" || user.DateOfBirth == "" ||
		user.PassportNumber == 0 || user.PassportSerie == "" || user.Password == "" {
		return false
	}

	return true
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &types.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(types.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
