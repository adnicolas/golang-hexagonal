package query

import "context"

type Bus interface {
	// Dispatch is the method used to dispatch new queries
	Dispatch(context.Context, Query) (QueryResponse, error)
	// Register is the method used to register a new command handler
	Register(Type, Handler)
}

type QueryResponse interface{}

//go:generate mockery --case=snake --outpkg=querymocks --output=querymocks --name=Bus

type Type string

type Query interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Query) (QueryResponse, error)
}
