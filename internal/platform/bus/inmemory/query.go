package inmemory

import (
	"context"

	"github.com/adnicolas/golang-hexagonal/kit/query"
)

// QueryBus is an in-memory implementation of the query.Bus
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

func (b *QueryBus) Dispatch(ctx context.Context, qry query.Query) (query.QueryResponse, error) {
	handler, ok := b.handlers[qry.Type()]
	if !ok {
		return nil, nil
	}

	// async strategy
	/*go func() {
		err := handler.Handle(ctx, qry)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", qry.Type(), err)
		}

	}()*/

	// sync strategy
	return handler.Handle(ctx, qry)
}

func (b *QueryBus) Register(qryType query.Type, handler query.Handler) {
	b.handlers[qryType] = handler
}
