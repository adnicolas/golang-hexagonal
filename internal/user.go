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

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

// Domain entity (data structure that represents a user)
type User struct {
	id UserId
	// TODO: Pass everything to value objects?
	name     string
	surname  string
	password string
	email    string
}

type GetUsersDto struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// NewUser creates a new user
func NewUser(id string, name string, surname string, password string, email string) (User, error) {
	idValueObject, err := NewUserId(id)
	if err != nil {
		return User{}, err
	}
	return User{
		id:       idValueObject,
		name:     name,
		surname:  surname,
		password: password,
		email:    email,
	}, nil
}

// Getters. They are not specific to the language and their use in Go is rare. However, they are the only way to achieve immutable structures

// ID return the user ID
func (u User) GetID() UserId {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetSurname() string {
	return u.surname
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPassword() string {
	return u.password
}
