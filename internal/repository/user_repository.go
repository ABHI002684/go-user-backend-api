package repository

import (
	"context"
	"time"

	"github.com/ABHI002684/go-user-backend-api/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

type UserWithAge struct {
	ID   int32     `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age"`
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (*sqlc.User, error) {
	user, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *UserRepository) GetByID(ctx context.Context, id int32) (*UserWithAge, error) {
	user, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,                 
		Age:  calculateAge(user.Dob),   
	}, nil
}

func (r *UserRepository) List(ctx context.Context) ([]UserWithAge, error) {
	users, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []UserWithAge
	for _, u := range users {
		result = append(result, UserWithAge{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob,
			Age:  calculateAge(u.Dob),
		})
	}
	return result, nil
}

func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (*UserWithAge, error) {
	user, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	return &UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
		Age:  calculateAge(user.Dob),
	}, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
