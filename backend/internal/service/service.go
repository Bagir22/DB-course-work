package service

import (
	"context"
	"courseWork/internal/types"
	"courseWork/internal/utils"
	"errors"
	"log"
	"time"
)

type Repository interface {
	AddUser(ctx context.Context, user types.UserLongData) (types.UserResponse, error)
	CheckUserExist(email, password string) (types.UserShortData, error)
	GetFlights(dep, des string, departureDate time.Time) ([]types.Flight, error)
	GetUserByEmail(email string) (types.UserLongData, error)
	GetSeatsForFlight(flightId int) ([]types.Seat, error)
	AddFlightBooking(booking types.BookFlight) error
	UpdateUser(updateData types.UserLongData) error
	GetUserIdByEmail(email string) (int, error)
	GetPassengerHistory(email, status, city string, date *time.Time) ([]types.History, error)
	IsFlightBookedByUser(flightId int, userId int) (bool, error)
	CancelFlightByID(flightId int, userId int) error
	GetAllFlights() ([]types.FlightControl, error)
	GetFlightById(flightId int) (types.FlightControl, error)
	CreateFlight(flight types.FlightCreate) error
	UpdateFlight(flight types.FlightControl) error
	DeleteFlight(id int) error
	GetAirlinesAircrafts() ([]types.AirlineAircrafts, error)
	GetAirports() ([]types.Airport, error)
}

type Service struct {
	repo Repository
}

func InitService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddUser(ctx context.Context, user types.UserLongData) (types.UserResponse, error) {
	validate := utils.ValidateUser(user)
	if !validate {
		return types.UserResponse{}, errors.New("Can't validate user")
	}

	pwd, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Err password hashing ", err)
	}

	user.Password = pwd

	userResponse, err := s.repo.AddUser(ctx, user)
	if err != nil {
		return types.UserResponse{}, err
	}

	return userResponse, nil
}

func (s *Service) CheckUserExist(email string, password string) (types.UserShortData, error) {
	user, err := s.repo.CheckUserExist(email, password)

	if err != nil {
		log.Println(err)
		return types.UserShortData{}, err
	}

	validate := utils.VerifyPassword(password, user.Password)
	if !validate {
		return types.UserShortData{}, errors.New("Invalid password")
	}

	return user, nil
}

func (s *Service) GetFlights(dep, des string, departureDate time.Time) ([]types.Flight, error) {
	return s.repo.GetFlights(dep, des, departureDate)
}

func (s *Service) GetUserByEmail(email string) (types.UserLongData, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *Service) GetSeatsForFlight(flightId int) ([]types.Seat, error) {
	return s.repo.GetSeatsForFlight(flightId)
}

func (s *Service) AddFlightBooking(booking types.BookFlight) error {
	return s.repo.AddFlightBooking(booking)
}

func (s *Service) GetUserIdByEmail(email string) (int, error) {
	return s.repo.GetUserIdByEmail(email)
}

func (s *Service) UpdateUser(updateData types.UserLongData) error {
	log.Println(updateData)
	err := s.repo.UpdateUser(updateData)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPassengerHistory(email, status, city string, date *time.Time) ([]types.History, error) {
	return s.repo.GetPassengerHistory(email, status, city, date)
}

func (s *Service) IsFlightBookedByUser(flightId int, userId int) (bool, error) {
	return s.repo.IsFlightBookedByUser(flightId, userId)
}

func (s *Service) CancelFlightByID(flightId int, userId int) error {
	err := s.repo.CancelFlightByID(flightId, userId)
	return err
}

func (s *Service) GetAllFlights() ([]types.FlightControl, error) {
	return s.repo.GetAllFlights()
}

func (s *Service) GetFlightById(flightId int) (types.FlightControl, error) {
	return s.repo.GetFlightById(flightId)
}

func (s *Service) CreateFlight(flight types.FlightCreate) error {
	return s.repo.CreateFlight(flight)
}

func (s *Service) UpdateFlight(flight types.FlightControl) error {
	return s.repo.UpdateFlight(flight)
}

func (s *Service) GetAirlinesAircrafts() ([]types.AirlineAircrafts, error) {
	return s.repo.GetAirlinesAircrafts()
}

func (s *Service) GetAirports() ([]types.Airport, error) {
	return s.repo.GetAirports()
}

func (s *Service) DeleteFlight(id int) error {
	return s.repo.DeleteFlight(id)
}
