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
		if err := rows.Scan(&flight.ID, &flight.Departure, &flight.Arrival, &flight.DepartureDate, &flight.Price); err != nil {
			return []types.Flight{}, err
		}
		flights = append(flights, flight)
	}

	if err := rows.Err(); err != nil {
		return []types.Flight{}, err
	}

	return flights, nil
}
