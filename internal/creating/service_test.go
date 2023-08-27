// Application service unit tests

package creating

import (
	"context"
	"errors"
	"testing"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/platform/storage/storagemocks"
	"github.com/adnicolas/golang-hexagonal/kit/event"
	"github.com/adnicolas/golang-hexagonal/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserService_CreateUser_RepositoryError(t *testing.T) {
	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userPass := "myPass"
	userEmail := "adri@gmail.com"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("usuario.User")).Return(errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userId, userName, userSurname, userPass, userEmail)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_CreateUser_EventsBusError(t *testing.T) {
	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userPass := "myPass"
	userEmail := "adri@gmail.com"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("usuario.User")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userId, userName, userSurname, userPass, userEmail)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

// Happy path
func Test_UserService_CreateUser_Succeed(t *testing.T) {
	userId := "c2f46a2b-9a8e-4614-8809-fedb86acf3b1"
	userName := "Adri"
	userSurname := "Nico"
	userPass := "myPass"
	userEmail := "adri@gmail.com"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("usuario.User")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		evt := events[0].(usuario.UserCreatedEvent)
		return evt.UserName() == userName
	})).Return(nil)

	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userId, userName, userSurname, userPass, userEmail)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
