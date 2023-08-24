package usuario

import (
	"context"
)

//go:generate mockery --case=snake --outpkg=storagemocks --output=storage/storagemocks --name=UserRepository

// Domain entity (data structure that represents a user)
type User struct {
	id       string
	name     string
	surname  string
	password string
	email    string
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

// NewUser creates a new user
func NewUser(id string, name string, surname string, password string, email string) User {
	return User{
		id:       id,
		name:     name,
		surname:  surname,
		password: password,
		email:    email,
	}
}

// ID return the user ID
func (u User) ID() string {
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
