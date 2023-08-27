package users

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/kit/query/querymocks"

	"github.com/adnicolas/golang-hexagonal/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_FindAll(t *testing.T) {
	queryBus := new(querymocks.Bus)

	gin.SetMode(gin.TestMode)
	r := gin.New()

	t.Run("it returns an empty array of usuario.GetUsersDto when there are no users", func(t *testing.T) {
		queryBus.On(
			"DispatchQuery",
			mock.Anything,
			mock.AnythingOfType("fetching.UserQuery"),
		).Return([]usuario.GetUsersDto{}, nil)
		r.GET("/users", FindAllController(queryBus))

		userRepository := new(storagemocks.UserRepository)
		userRepository.On("FindAll", mock.Anything).Return([]usuario.GetUsersDto{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response []usuario.GetUsersDto
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		assert.Equal(t, []usuario.GetUsersDto{}, response)
	})

	t.Run("it returns an array of usuario.GetUsersDto when there are users", func(t *testing.T) {
		user, _ := usuario.NewUser("8a1c5cdc-ba57-445a-994d-aa412d23723f", "Adri", "Nico", "randomPassword", "adrian@gmail.com")
		var users = []usuario.User{user}

		var responseUsers = convertToGetUsersDto(users)

		queryBus.On(
			"DispatchQuery",
			mock.Anything,
			mock.AnythingOfType("fetching.UserQuery"),
		).Return(responseUsers, nil)
		r.GET("/users", FindAllController(queryBus))

		userRepository := new(storagemocks.UserRepository)
		userRepository.On("FindAll", mock.Anything).Return(responseUsers, nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response []usuario.GetUsersDto
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		assert.Equal(t, responseUsers, response)
	})
}

func convertToGetUsersDto(users []usuario.User) []usuario.GetUsersDto {
	var response []usuario.GetUsersDto

	if len(users) == 0 {
		return response
	}

	for _, user := range users {
		response = append(response, usuario.GetUsersDto{
			Id:      user.GetID().String(),
			Name:    user.GetName(),
			Surname: user.GetSurname(),
			Email:   user.GetEmail(),
		})
	}

	return response
}
