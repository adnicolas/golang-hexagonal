// Application service unit tests

package fetching

import (
	"context"
	"errors"
	"testing"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserService_FindAllUsers_RepositoryError(t *testing.T) {
	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).Return([]usuario.GetUsersDto{}, errors.New("something unexpected happened"))

	userService := NewUserService(userRepositoryMock)

	_, err := userService.FindAllUsers(context.Background())

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

// Happy path
func Test_UserService_FindAllUsers_Succeed(t *testing.T) {
	var users []usuario.GetUsersDto

	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userEmail := "adri@gmail.com"

	user := usuario.GetUsersDto{
		Id:      userId,
		Name:    userName,
		Surname: userSurname,
		Email:   userEmail,
	}

	users = append(users, user)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).Return(users, nil)

	userService := NewUserService(userRepositoryMock)

	_, err := userService.FindAllUsers(context.Background())

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
