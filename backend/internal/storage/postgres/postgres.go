package postgres

import (
	"context"
	"courseWork/Database/Queries"
	"courseWork/internal/config"
	"courseWork/internal/types"
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

	err := d.db.QueryRow(Queries.CheckUserExistQuery, email).Scan(&user.Email, &user.Password)
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
			&user.PassportSerie, &user.PassportNumber, &user.Password)
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
	_, err := d.db.Exec(Queries.UpdateUserById, user.FirstName, user.LastName,
		user.Email, user.Phone, user.DateOfBirth, user.PassportSerie, user.PassportNumber, user.Id)

	if err != nil {
		log.Println("Add user to db err: ", err)
		return err
	}

	return nil
}

func (d *Db) GetPassengerHistory(email string) ([]types.History, error) {
	var history []types.History

	rows, err := d.db.Query(Queries.GetHistory, email)
	if err != nil {
		return []types.History{}, err
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
