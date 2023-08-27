package creating

import (
	"context"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/kit/event"
)

type UserService struct {
	userRepository usuario.UserRepository
	eventBus       event.Bus
}

func NewUserService(userRepository usuario.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (s UserService) CreateUser(ctx context.Context, id, name, surname, password, email string) error {
	user, err := usuario.NewUser(id, name, surname, password, email)
	if err != nil {
		return err
	}
	if err := s.userRepository.Save(ctx, user); err != nil {
		return err
	}
	return s.eventBus.Publish(ctx, user.PullEvents())
}
