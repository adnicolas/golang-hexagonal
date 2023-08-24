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
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=storage/storagemocks --name=UserRepository

// Domain entity (data structure that represents a user)
type User struct {
	id       UserId
	name     string
	surname  string
	password string
	email    string
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

// ID return the user ID
func (u User) ID() UserId {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Surname() string {
	return u.surname
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}
