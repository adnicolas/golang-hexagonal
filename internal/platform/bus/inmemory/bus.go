package inmemory

import (
	"context"
	"log"

	"github.com/adnicolas/golang-hexagonal/kit/bus"
)

type Bus struct {
	commandHandlers map[bus.Type]bus.CommandHandler
	queryHandlers   map[bus.Type]bus.QueryHandler
}

func NewBus() *Bus {
	return &Bus{
		commandHandlers: make(map[bus.Type]bus.CommandHandler),
		queryHandlers:   make(map[bus.Type]bus.QueryHandler),
	}
}

func (b *Bus) DispatchCommand(ctx context.Context, cmd bus.Command) error {
	handler, ok := b.commandHandlers[cmd.Type()]
	if !ok {
		return nil
	}

	// async strategy
	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		}

	}()

	// sync strategy
	// return handler.Handle(ctx, cmd)

	return nil
}

func (b *Bus) RegisterCommand(cmdType bus.Type, handler bus.CommandHandler) {
	b.commandHandlers[cmdType] = handler
}

func (b *Bus) DispatchQuery(ctx context.Context, qry bus.Query) (bus.QueryResponse, error) {
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

func (b *Bus) RegisterQuery(qryType bus.Type, handler bus.QueryHandler) {
	b.queryHandlers[qryType] = handler
}
