package command

import "context"

type Bus interface {
	// Dispatch is the method used to dispatch new commands
	Dispatch(context.Context, Command) error
	// Register is the method used to register a new command handler
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) error
}
