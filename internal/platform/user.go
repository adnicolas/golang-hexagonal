package mooc

import "github.com/google/uuid"

// Domain entity (data structure that represents a user)
type User struct {
	uuid     uuid.UUID
	email    string
	password string
	name     string
	surname  string
}

// NewUser creates a new user
func NewUser(uuid uuid.UUID, name string, surname string, email string, password string) User {
	return User{
		uuid:     uuid,
		name:     name,
		surname:  surname,
		password: password,
		email:    email,
	}
}

// ID return the user unique identifier
func (u User) ID() uuid.UUID {
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
