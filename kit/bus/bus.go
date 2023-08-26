package bus

import "context"

type Bus interface {
	DispatchCommand(context.Context, Command) error
	RegisterCommand(Type, CommandHandler)
	DispatchQuery(context.Context, Query) (QueryResponse, error)
	RegisterQuery(Type, QueryHandler)
}

//go:generate mockery --case=snake --outpkg=busmocks --output=busmocks --name=Bus

type Type string

type Command interface {
	Type() Type
}

type CommandHandler interface {
	Handle(context.Context, Command) error
}

type QueryResponse interface{}

type Query interface {
	Type() Type
}

type QueryHandler interface {
	Handle(context.Context, Query) (QueryResponse, error)
}
