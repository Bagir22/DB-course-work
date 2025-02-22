package utils

import (
	"courseWork/Database/Queries"
	"courseWork/internal/types"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
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
		user.PassportNumber == "" || user.PassportSerie == "" || user.Password == "" {
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

	log.Println("Token: ", tokenString)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*types.Claims, error) {
	claims := &types.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return types.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func MarkFlightsAsDone(db *sql.DB) error {
	res, err := db.Exec(Queries.MarkFlightsAsDone)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Updated %d flight bookings to 'done'", rowsAffected)
	return nil
}

func ConvertToUserLongData(userFromFront types.UserLongDataFromFront) (types.UserLongData, error) {
	user := types.UserLongData{
		Id:             userFromFront.Id,
		FirstName:      userFromFront.FirstName,
		LastName:       userFromFront.LastName,
		Email:          userFromFront.Email,
		Phone:          userFromFront.Phone,
		DateOfBirth:    userFromFront.DateOfBirth,
		PassportSerie:  userFromFront.PassportSerie,
		PassportNumber: userFromFront.PassportNumber,
		Password:       userFromFront.Password,
	}
	if userFromFront.Image != nil && userFromFront.Image.Filename != "" {
		user.Image = userFromFront.Image.Filename
	}

	return user, nil
}
