package fetching

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

func (s UserService) FindAllUsers(ctx context.Context) ([]usuario.GetUsersDto, error) {
	return s.userRepository.FindAll(ctx)
}
