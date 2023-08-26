package fetching

import (
	"context"
	"errors"

	"github.com/adnicolas/golang-hexagonal/kit/bus"
)

const UserQueryType bus.Type = "query.fetching.users"

type UserQuery struct{}

func NewUserQuery() UserQuery {
	return UserQuery{}
}

func (c UserQuery) Type() bus.Type {
	return UserQueryType
}

// UserQueryHandler is the query handler responsible for creating users
type UserQueryHandler struct {
	service UserService
}

func NewUserQueryHandler(service UserService) UserQueryHandler {
	return UserQueryHandler{
		service: service,
	}
}

// Handle implements the bus.QueryHandler interface
func (handler UserQueryHandler) Handle(ctx context.Context, qry bus.Query) (bus.QueryResponse, error) {
	// Casting of the generic to the user query
	_, ok := qry.(UserQuery)
	if !ok {
		return nil, errors.New("unexpected query")
	}

	return handler.service.FindAllUsers(ctx)
}
