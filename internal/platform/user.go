// Domain layer

package usuario

import (
	"context"
	"errors"
	"fmt"

	// Trade off / Exception to hexagonal architecture to make our lives easier with uuid
	"github.com/google/uuid"
)

// Sentinel error
var ErrInvalidUserId = errors.New("invalid User ID")

// Value objects strategy
type UserId struct {
	value string
}

// NewUserId instantiate the value object for UserId
func NewUserId(value string) (UserId, error) {
	// Reason for using value object strategy, this check to see if it is getting a valid uuid
	v, err := uuid.Parse(value)
	if err != nil {
		return UserId{}, fmt.Errorf("%w: %s", ErrInvalidUserId, value)
	}
	return UserId{
		value: v.String(),
	}, nil
}

// String type converts the UserID into string
func (id UserId) String() string {
	return id.value
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
	FindAll(ctx context.Context) ([]GetUsersDto, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=storage/storagemocks --name=UserRepository

// Domain entity (data structure that represents a user)
type User struct {
	Id       UserId
	Name     string
	Surname  string
	Password string
	Email    string
}

type GetUsersDto struct {
	Id      string
	Name    string
	Surname string
	Email   string
}

// NewUser creates a new user
func NewUser(id string, name string, surname string, password string, email string) (User, error) {
	idValueObject, err := NewUserId(id)
	if err != nil {
		return User{}, err
	}
	return User{
		Id:       idValueObject,
		Name:     name,
		Surname:  surname,
		Password: password,
		Email:    email,
	}, nil
}

// ID return the user ID
func (u User) GetID() UserId {
	return u.Id
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetSurname() string {
	return u.Surname
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.Password
}
