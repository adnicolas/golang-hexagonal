package command

import "context"

type Bus interface {
	DispatchCommand(context.Context, Command) error
	RegisterCommand(Type, CommandHandler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

type Type string

type Command interface {
	Type() Type
}

type CommandHandler interface {
	Handle(context.Context, Command) error
}
