package service

import (
	"context"
	"courseWork/internal/types"
	"courseWork/internal/utils"
	"errors"
	"log"
)

type Repository interface {
	AddUser(ctx context.Context, user types.UserLongData) (types.UserResponse, error)
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
