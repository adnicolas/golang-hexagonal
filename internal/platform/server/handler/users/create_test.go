// Infrastructure layer tests example

package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnicolas/golang-hexagonal/kit/command/commandmocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Create(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"DispatchCommand",
		mock.Anything,
		mock.AnythingOfType("creating.UserCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/user", CreateController(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		saveUserReq := saveRequest{
			Id:   "2a6e370c-e015-47c9-9bff-b09bf3da7420",
			Name: "Adri",
		}

		body, err := json.Marshal(saveUserReq)
		// If the condition is not met, the require implies that the execution of the test will stop at this point
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		result := recorder.Result()
		defer result.Body.Close()

		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		saveUserReq := saveRequest{
			Id:       "c2f46a2b-9a8e-4614-8809-fedb86acf3b1",
			Name:     "Adri",
			Surname:  "Nico",
			Password: "2023Password!",
			Email:    "adrian.nicolas@geograma.com",
		}

		body, err := json.Marshal(saveUserReq)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		result := recorder.Result()
		defer result.Body.Close()

		assert.Equal(t, http.StatusCreated, result.StatusCode)
	})

	t.Run("given a not valid ID it returns 400", func(t *testing.T) {
		saveUserReq := saveRequest{
			Id:       "invalid-id",
			Name:     "Adri",
			Surname:  "Nico",
			Password: "2023Password!",
			Email:    "adrian.nicolas@geograma.com",
		}

		body, err := json.Marshal(saveUserReq)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		result := recorder.Result()
		defer result.Body.Close()

		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}
