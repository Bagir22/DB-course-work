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

type SignupRequest struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DateOfBirth    string `json:"dateOfBirth"`
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
	Password       string `json:"password"`
}

type UserLongData struct {
	Id             int    `json:"id" db:"id"`
	FirstName      string `json:"firstName" db:"firstName"`
	LastName       string `json:"lastName" db:"lastName"`
	Email          string `json:"email" db:"email"`
	Phone          string `json:"phone" db:"phone"`
	DateOfBirth    string `json:"dateOfBirth" db:"dateOfBirth"`
	PassportSerie  string `json:"passportSerie" db:"passportserie"`
	PassportNumber int    `json:"passportNumber" db:"passportnumber"`
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

type Flight struct {
	Id             int    `json:"id" db:"id"`
	Departure      string `json:"departure" db:"departure_airport"`
	DepartureCity  string `json:"departure_city" db:"departure_city"`
	Arrival        string `json:"arrival" db:"arrival_airport"`
	ArrivalCity    string `json:"arrival_city" db:"arrival_city"`
	DepartureDate  string `json:"departure_date" db:"departure_date"`
	ArrivalDate    string `json:"arrival_date" db:"arrival_date"`
	Price          string `json:"price" db:"price"`
	AvailableSeats int    `json:"available_seats" db:"available_seats"`
}

type Seat struct {
	Row    string `json:"row" db:"row"`
	Seat   int    `json:"seat" db:"seat"`
	Status string `json:"status" db:"status"`
}

type BookFlight struct {
	FlightId    int    `json:"flightId" db:"flightid"`
	PassengerId int    `json:"passengerId" db:"passengerid"`
	Status      string `json:"status" db:"status"`
	Row         string `json:"row" db:"row"`
	Seat        int    `json:"seat" db:"seat"`
}

type History struct {
	Id            int    `json:"id" db:"id"`
	Departure     string `json:"departure" db:"departure_airport"`
	DepartureCity string `json:"departure_city" db:"departure_city"`
	Arrival       string `json:"arrival" db:"arrival_airport"`
	ArrivalCity   string `json:"arrival_city" db:"arrival_city"`
	DepartureDate string `json:"departure_date" db:"departure_date"`
	ArrivalDate   string `json:"arrival_date" db:"arrival_date"`
	Price         string `json:"price" db:"price"`
	Status        string `json:"status" db:"status"`
	Row           string `json:"row" db:"row"`
	Seat          int    `json:"seat" db:"seat"`
}
