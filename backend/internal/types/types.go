package types

import (
	"github.com/dgrijalva/jwt-go"
	"mime/multipart"
	"time"
)

var JwtSecret = []byte("SomeSecretKey")

type Response struct {
	Message     string `json:"message"`
	Description any    `json:"description"`
}

type UserShortData struct {
	Id       int
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string
	IsAdmin  bool
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
	Id             int    `json:"id" db:"id" form:"id"`
	FirstName      string `json:"firstName" db:"firstName" form:"firstName"`
	LastName       string `json:"lastName" db:"lastName" form:"lastName"`
	Email          string `json:"email" db:"email" form:"email"`
	Phone          string `json:"phone" db:"phone" form:"phone"`
	DateOfBirth    string `json:"dateOfBirth" db:"dateOfBirth" form:"dateOfBirth"`
	PassportSerie  string `json:"passportSerie" db:"passportserie" form:"passportSerie"`
	PassportNumber string `json:"passportNumber" db:"passportnumber" form:"passportNumber"`
	Password       string `json:"password" db:"password" form:"password"`
	Image          string `json:"image" db:"image" form:"image"`
}

type UserLongDataFromFront struct {
	Id             int                   `json:"id" db:"id" form:"id"`
	FirstName      string                `json:"firstName" db:"firstName" form:"firstName"`
	LastName       string                `json:"lastName" db:"lastName" form:"lastName"`
	Email          string                `json:"email" db:"email" form:"email"`
	Phone          string                `json:"phone" db:"phone" form:"phone"`
	DateOfBirth    string                `json:"dateOfBirth" db:"dateOfBirth" form:"dateOfBirth"`
	PassportSerie  string                `json:"passportSerie" db:"passportserie" form:"passportSerie"`
	PassportNumber string                `json:"passportNumber" db:"passportnumber" form:"passportNumber"`
	Password       string                `json:"password" db:"password" form:"password"`
	Image          *multipart.FileHeader `json:"image" db:"image" form:"image"`
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

type FlightControl struct {
	Id                 int       `json:"id"`
	AircraftId         int       `json:"aircraft_id"`
	AircraftName       string    `json:"aircraft_name"`
	AirlineId          int       `json:"airline_id"`
	AirlineName        string    `json:"airline_name"`
	DepartureId        int       `json:"departure_id"`
	DepartureAirport   string    `json:"departure_airport"`
	DepartureCity      string    `json:"departure_city"`
	DepartureCountry   string    `json:"departure_country"`
	DestinationId      int       `json:"destination_id"`
	DestinationAirport string    `json:"destination_airport"`
	DestinationCity    string    `json:"destination_city"`
	DestinationCountry string    `json:"destination_country"`
	DepartureDateTime  time.Time `json:"departure_datetime"`
	ArrivalDateTime    time.Time `json:"arrival_datetime"`
	Price              float64   `json:"price"`
}

type AirlineAircrafts struct {
	AirlineId    int    `json:"airline_id" db:"airline_id"`
	AirlineName  string `json:"airline_name" db:"airline_name"`
	AircraftId   int    `json:"aircraft_id" db:"aircraft_id"`
	AircraftName string `json:"aircraft_name" db:"aircraft_name"`
}

type Airport struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type FlightCreate struct {
	AircraftId        int       `json:"aircraft_id" db:"aircraftid"`
	AirlineId         int       `json:"airline_id" db:"airlineid"`
	DepartureId       int       `json:"departure_id" db:"departureid"`
	DestinationId     int       `json:"arrival_id" db:"arrivalid"`
	DepartureDateTime time.Time `json:"departure_datetime" db:"destinationdatetime"`
	ArrivalDateTime   time.Time `json:"arrival_datetime" db:"arrivaldatetime"`
	Price             float64   `json:"price" db:""`
}
