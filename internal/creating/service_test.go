// Application service unit tests

package creating

import (
	"context"
	"errors"
	"testing"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_UserService_CreateUser_RepositoryError(t *testing.T) {
	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userPass := "myPass"
	userEmail := "adri@gmail.com"

	user, err := usuario.NewUser(userId, userName, userSurname, userPass, userEmail)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, user).Return(errors.New("something unexpected happened"))

	userService := NewUserService(userRepositoryMock)

	err = userService.CreateUser(context.Background(), userId, userName, userSurname, userPass, userEmail)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

// Happy path
func Test_UserService_CreateUser_Succeed(t *testing.T) {
	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userPass := "myPass"
	userEmail := "adri@gmail.com"

	user, err := usuario.NewUser(userId, userName, userSurname, userPass, userEmail)
	require.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, user).Return(nil)

	userService := NewUserService(userRepositoryMock)

	err = userService.CreateUser(context.Background(), userId, userName, userSurname, userPass, userEmail)

	// Without this line the test would pass even though it was not calling the repo.
	// With this line we make sure that the repo call is checked and with the user that is passed in the following line:
	// userRepositoryMock.On("Save", mock.Anything, user).Return(nil)
	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
