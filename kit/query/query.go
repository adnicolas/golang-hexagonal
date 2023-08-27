package query

import "context"

type Bus interface {
	DispatchQuery(context.Context, Query) (QueryResponse, error)
	RegisterQuery(Type, QueryHandler)
}

//go:generate mockery --case=snake --outpkg=querymocks --output=querymocks --name=Bus

type Type string

type QueryResponse interface{}

type Query interface {
	Type() Type
}

type QueryHandler interface {
	Handle(context.Context, Query) (QueryResponse, error)
}
