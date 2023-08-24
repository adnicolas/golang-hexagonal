package usuario

import (
	"context"
)

// Domain entity (data structure that represents a user)
type User struct {
	uuid     string //uuid.UUID
	name     string
	surname  string
	password string
	email    string
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

// NewUser creates a new user
func NewUser(uuid string /*uuid.UUID*/, name string, surname string, password string, email string) User {
	return User{
		uuid:     uuid,
		name:     name,
		surname:  surname,
		password: password,
		email:    email,
	}
}

// ID return the user UUID
func (u User) ID() string /*uuid.UUID*/ {
	return u.uuid
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
