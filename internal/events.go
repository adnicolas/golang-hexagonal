package usuario

import "github.com/adnicolas/golang-hexagonal/kit/event"

const UserCreatedEventType event.Type = "events.user.created"

type UserCreatedEvent struct {
	// Composition
	event.BaseEvent
	id       string
	name     string
	surname  string
	password string
	email    string
}

func NewUserCreatedEvent(id, name, surname, password, email string) UserCreatedEvent {
	return UserCreatedEvent{
		id:        id,
		name:      name,
		surname:   surname,
		password:  password,
		email:     email,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e UserCreatedEvent) Type() event.Type {
	return UserCreatedEventType
}

func (e UserCreatedEvent) UserID() string {
	return e.id
}

func (e UserCreatedEvent) UserName() string {
	return e.name
}

func (e UserCreatedEvent) UserSurname() string {
	return e.surname
}

func (e UserCreatedEvent) UserPassword() string {
	return e.password
}

func (e UserCreatedEvent) UserEmail() string {
	return e.email
}
