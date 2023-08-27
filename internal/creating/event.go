package creating

import (
	"context"
	"errors"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/increasing"
	"github.com/adnicolas/golang-hexagonal/kit/event"
)

type IncreaseUsersCounterOnUserCreated struct {
	increasingService increasing.UserCounterIncreaserService
}

func NewIncreaseUsersCounterOnUserCreated(increaserService increasing.UserCounterIncreaserService) IncreaseUsersCounterOnUserCreated {
	return IncreaseUsersCounterOnUserCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseUsersCounterOnUserCreated) Handle(_ context.Context, evt event.Event) error {
	// Casts to domain event
	userCreatedEvt, ok := evt.(usuario.UserCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(userCreatedEvt.Id())
}
