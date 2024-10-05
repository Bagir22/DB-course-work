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
