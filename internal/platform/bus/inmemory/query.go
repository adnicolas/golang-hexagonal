package inmemory

import (
	"context"

	"github.com/adnicolas/golang-hexagonal/kit/query"
)

type QueryBus struct {
	queryHandlers map[query.Type]query.QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		queryHandlers: make(map[query.Type]query.QueryHandler),
	}
}

func (b *QueryBus) DispatchQuery(ctx context.Context, qry query.Query) (query.QueryResponse, error) {
	handler, ok := b.queryHandlers[qry.Type()]
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

func (b *QueryBus) RegisterQuery(qryType query.Type, handler query.QueryHandler) {
	b.queryHandlers[qryType] = handler
}
