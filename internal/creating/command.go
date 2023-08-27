package creating

import (
	"context"
	"errors"

	"github.com/adnicolas/golang-hexagonal/kit/command"
)

const UserCommandType command.Type = "command.creating.user"

// UserCommand is the command dispatched to create a new user
type UserCommand struct {
	id       string
	name     string
	surname  string
	password string
	email    string
}

func NewUserCommand(id, name, surname, password, email string) UserCommand {
	return UserCommand{
		id:       id,
		name:     name,
		surname:  surname,
		password: password,
		email:    email,
	}
}

func (c UserCommand) Type() command.Type {
	return UserCommandType
}

// UserCommandHandler is the command handler responsible for creating users
type UserCommandHandler struct {
	service UserService
}

func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

// Handle implements the command.CommandHandler interface
func (handler UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	// Casting of the generic to the user command
	createUserCmd, ok := cmd.(UserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return handler.service.CreateUser(
		ctx,
		createUserCmd.id,
		createUserCmd.name,
		createUserCmd.surname,
		createUserCmd.password,
		createUserCmd.email,
	)
}
