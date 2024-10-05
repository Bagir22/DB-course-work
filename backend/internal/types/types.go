package types

import "github.com/dgrijalva/jwt-go"

var JwtSecret = []byte("SomeSecretKey")

type Response struct {
	Message     string `json:"message"`
	Description any    `json:"description"`
}

type UserShortData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLongData struct {
	FirstName      string `json:"firstName" db:"firstName"`
	LastName       string `json:"lastName" db:"lastName"`
	Email          string `json:"email" db:"email"`
	Phone          int    `json:"phone" db:"phone"`
	DateOfBirth    string `json:"dateOfBirth" db:"dateOfBirth"`
	PassportSerie  string `json:"passportSerie" db:"passportSerie"`
	PassportNumber int    `json:"passportNumber" db:"passportNumber"`
	Password       string `json:"password" db:"password"`
}

type UserResponse struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
