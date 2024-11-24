package postgres

import (
	"context"
	"courseWork/Database/Queries"
	"courseWork/internal/config"
	"courseWork/internal/types"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Db struct {
	db *sqlx.DB
}

func InitDb(db *sqlx.DB) *Db {
	return &Db{
		db: db,
	}
}

func InitConn(cfg *config.Config) (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"db", cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDatabase, "disable")

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Db) AddUser(ctx context.Context, user types.UserLongData) (types.UserResponse, error) {
	_, err := d.db.Exec(Queries.InsertUserQuery, user.FirstName, user.LastName,
		user.Email, user.Phone, user.DateOfBirth, user.PassportSerie, user.PassportNumber, user.Password)

	if err != nil {
		log.Println("Add user to db err: ", err)
		return types.UserResponse{}, err
	}

	return types.UserResponse{user.Email, user.Password}, nil
}

func (d *Db) CheckUserExist(email string, password string) (types.UserShortData, error) {
	var user types.UserShortData

	err := d.db.QueryRow(Queries.CheckUserExistQuery, email).
		Scan(&user.Id, &user.Email, &user.Password, &user.Image, &user.IsAdmin)
	if err != nil {
		return types.UserShortData{}, err
	}

	return user, nil
}

func (d *Db) GetFlights(dep, des string, depDate time.Time) ([]types.Flight, error) {
	var flights []types.Flight

	rows, err := d.db.Query(Queries.GetFlights, dep, des, depDate)
	if err != nil {
		return []types.Flight{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var flight types.Flight
		if err := rows.Scan(
			&flight.Id,
			&flight.Departure,
			&flight.DepartureCity,
			&flight.Arrival,
			&flight.ArrivalCity,
			&flight.DepartureDate,
			&flight.ArrivalDate,
			&flight.Price,
			&flight.AvailableSeats,
		); err != nil {
			return []types.Flight{}, err
		}
		flights = append(flights, flight)
	}

	if err := rows.Err(); err != nil {
		return []types.Flight{}, err
	}

	return flights, nil
}

func (d *Db) GetUserByEmail(email string) (types.UserLongData, error) {
	var user types.UserLongData

	err := d.db.QueryRow(Queries.GetUserLongByEmailQuery, email).
		Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.DateOfBirth,
			&user.PassportSerie, &user.PassportNumber, &user.Password, &user.Image)
	if err != nil {
		return types.UserLongData{}, err
	}

	return user, nil
}

func (d *Db) GetSeatsForFlight(flightId int) ([]types.Seat, error) {
	var seats []types.Seat

	rows, err := d.db.Query(Queries.GetSeatsForFlightQuery, flightId)
	if err != nil {
		return []types.Seat{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var seat types.Seat
		if err := rows.Scan(
			&seat.Row,
			&seat.Seat,
			&seat.Status,
		); err != nil {
			return []types.Seat{}, err
		}
		seats = append(seats, seat)
	}

	return seats, nil
}

func (d *Db) AddFlightBooking(booking types.BookFlight) error {

	_, err := d.db.Exec(Queries.InsertBooking, booking.FlightId, booking.PassengerId, booking.Status, booking.Row, booking.Seat)
	if err != nil {
		log.Printf("Failed to add flight booking: %v", err)
		return fmt.Errorf("Failed to add flight booking: %v", err)
	}

	return nil
}

func (d *Db) GetUserIdByEmail(email string) (int, error) {
	var id int

	err := d.db.QueryRow(Queries.GetUserIdByEmail, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *Db) UpdateUser(user types.UserLongData) error {
	var err error
	if user.Image != "" {
		_, err = d.db.Exec(Queries.UpdateUserByIdWithImage, user.FirstName, user.LastName,
			user.Email, user.Phone, user.DateOfBirth, user.PassportSerie, user.PassportNumber, user.Image, user.Id)
	} else {
		_, err = d.db.Exec(Queries.UpdateUserByIdWithoutImage, user.FirstName, user.LastName,
			user.Email, user.Phone, user.DateOfBirth, user.PassportSerie, user.PassportNumber, user.Id)
	}

	if err != nil {
		log.Println("Add user to db err: ", err)
		return err
	}

	return nil
}

func (d *Db) GetPassengerHistory(email, status, city string, date *time.Time) ([]types.History, error) {
	var history []types.History

	statusParam := sql.NullString{}
	if status != "" {
		statusParam.String = status
		statusParam.Valid = true
	}

	cityParam := sql.NullString{}
	if city != "" {
		cityParam.String = city
		cityParam.Valid = true
	}

	var dateParam sql.NullTime
	if date != nil {
		dateParam.Time = *date
		dateParam.Valid = true
	}

	log.Printf("Query parameters: email=%s, status=%v, city=%v, date=%v", email, statusParam, cityParam, dateParam)

	rows, err := d.db.Query(Queries.GetHistory, email, statusParam, cityParam, dateParam)

	if err != nil {
		return nil, fmt.Errorf("error querying passenger history: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var h types.History
		if err := rows.Scan(
			&h.Id,
			&h.Departure,
			&h.DepartureCity,
			&h.Arrival,
			&h.ArrivalCity,
			&h.DepartureDate,
			&h.ArrivalDate,
			&h.Price,
			&h.Status,
			&h.Row,
			&h.Seat,
		); err != nil {
			return []types.History{}, err
		}
		history = append(history, h)
	}

	return history, nil

}

func (db *Db) IsFlightBookedByUser(flightId int, userId int) (bool, error) {
	var count int
	err := db.db.QueryRow(Queries.IsBookedFlightForUser, flightId, userId).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *Db) CancelFlightByID(flightId int, userID int) error {
	_, err := db.db.Exec(Queries.CancelBooking, flightId, userID)
	if err != nil {
		return fmt.Errorf("could not update flight status: %v", err)
	}
	return nil
}

func (db *Db) GetAllFlights() ([]types.FlightControl, error) {
	rows, err := db.db.Query(Queries.GetFlightList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []types.FlightControl
	for rows.Next() {
		var flight types.FlightControl
		err := rows.Scan(
			&flight.Id,
			&flight.AircraftId,
			&flight.AircraftName,
			&flight.AirlineId,
			&flight.AirlineName,
			&flight.DepartureId,
			&flight.DepartureAirport,
			&flight.DepartureCity,
			&flight.DepartureCountry,
			&flight.DestinationId,
			&flight.DestinationAirport,
			&flight.DestinationCity,
			&flight.DestinationCountry,
			&flight.DepartureDateTime,
			&flight.ArrivalDateTime,
			&flight.Price,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	return flights, nil
}

func (db *Db) GetFlightById(flightId int) (types.FlightControl, error) {
	log.Println(flightId)
	var flight types.FlightControl
	err := db.db.QueryRow(Queries.GetFullFlightById, flightId).
		Scan(&flight.Id,
			&flight.AircraftId,
			&flight.AircraftName,
			&flight.AirlineId,
			&flight.AirlineName,
			&flight.DepartureId,
			&flight.DepartureAirport,
			&flight.DepartureCity,
			&flight.DepartureCountry,
			&flight.DestinationId,
			&flight.DestinationAirport,
			&flight.DestinationCity,
			&flight.DestinationCountry,
			&flight.DepartureDateTime,
			&flight.ArrivalDateTime,
			&flight.Price)
	if err != nil {
		return types.FlightControl{}, err
	}

	return flight, nil
}

func (db *Db) GetAirlinesAircrafts() ([]types.AirlineAircrafts, error) {
	rows, err := db.db.Query(Queries.GetAirlinesAircrafts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []types.AirlineAircrafts
	for rows.Next() {
		var d types.AirlineAircrafts
		err := rows.Scan(
			&d.AirlineId,
			&d.AirlineName,
			&d.AircraftId,
			&d.AircraftName,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func (db *Db) UpdateFlight(flight types.FlightControl) error {
	_, err := db.db.Exec(Queries.UpdateFlight, flight.AircraftId, flight.DepartureId, flight.DestinationId,
		flight.DepartureDateTime, flight.ArrivalDateTime, flight.Price, flight.Id)
	return err
}

func (db *Db) DeleteFlight(id int) error {
	_, err := db.db.Exec(Queries.DeleteFlight, id)
	return err
}

func (db *Db) GetAirports() ([]types.Airport, error) {
	rows, err := db.db.Query(Queries.GetAirports)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airports []types.Airport
	for rows.Next() {
		var airport types.Airport
		err := rows.Scan(
			&airport.Id,
			&airport.Name,
		)
		if err != nil {
			return nil, err
		}
		airports = append(airports, airport)
	}
	return airports, nil
}

func (db *Db) CreateFlight(flight types.FlightCreate) error {
	_, err := db.db.Exec(Queries.CreateFlight, flight.AircraftId, flight.DepartureId, flight.DestinationId,
		flight.DepartureDateTime, flight.ArrivalDateTime, flight.Price)
	return err
}
