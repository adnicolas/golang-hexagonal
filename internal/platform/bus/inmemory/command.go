package inmemory

import (
	"context"
	"log"

	"github.com/adnicolas/golang-hexagonal/kit/command"
)

type CommandBus struct {
	commandHandlers map[command.Type]command.CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		commandHandlers: make(map[command.Type]command.CommandHandler),
	}
}

func (b *CommandBus) DispatchCommand(ctx context.Context, cmd command.Command) error {
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

func (b *CommandBus) RegisterCommand(cmdType command.Type, handler command.CommandHandler) {
	b.commandHandlers[cmdType] = handler
}
