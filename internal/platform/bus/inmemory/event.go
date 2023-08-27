package inmemory

import (
	"context"

	"github.com/adnicolas/golang-hexagonal/kit/event"
)

// In-memory implementation of the event.Bus
type EventBus struct {
	handlers map[event.Type][]event.Handler
}

// Initializes a new EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

// event.Bus publish interface sync implementation
func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, evt)
			// concurrent (async) strategy
			// go handler.Handle(ctx, evt)
		}
	}

	return nil
}

// event.Bus subscribe interface implementation
func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}
