package inmemory

import (
	"context"
	"log"

	"github.com/adnicolas/golang-hexagonal/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
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

func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
