package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Bus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

type Handler interface {
	Handle(context.Context, Event) error
}

type Type string

type Event interface {
	Id() string
	AggregateId() string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	eventId     string
	aggregateId string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventId:     uuid.New().String(),
		aggregateId: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (b BaseEvent) Id() string {
	return b.eventId
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateId() string {
	return b.aggregateId
}
