package creating

import (
	"context"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
)

type UserService struct {
	userRepository usuario.UserRepository
}

func NewUserService(userRepository usuario.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) CreateUser(ctx context.Context, id, name, surname, password, email string) error {
	user, err := usuario.NewUser(id, name, surname, password, email)
	if err != nil {
		return err
	}
	return s.userRepository.Save(ctx, user)
}
