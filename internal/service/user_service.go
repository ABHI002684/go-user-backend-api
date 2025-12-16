package service

import (
	"context"
	"time"

	"github.com/ABHI002684/go-user-backend-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create user
func (s *UserService) Create(ctx context.Context, name string, dob string) (*repository.UserWithAge, error) {
	
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, name, parsedDOB)
	if err != nil {
		return nil, err
	}

	return &repository.UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
		Age:  CalculateAge(user.Dob),
	}, nil
}

// Get user by ID
func (s *UserService) GetByID(ctx context.Context, id int32) (*repository.UserWithAge, error) {
	return s.repo.GetByID(ctx, id)
}

// List users
func (s *UserService) List(ctx context.Context) ([]repository.UserWithAge, error) {
	return s.repo.List(ctx)
}

// Update user
func (s *UserService) Update(ctx context.Context, id int32, name string, dob string) (*repository.UserWithAge, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, id, name, parsedDOB)
}

// Delete user
func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
