package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnicolas/golang-hexagonal/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Create(t *testing.T) {
	userRepository := new(storagemocks.UserRepository)
	userRepository.On("Save", mock.Anything, mock.AnythingOfType("usuario.User")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/user", CreateController(userRepository))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createUserReq := createRequest{
			Uuid: "2a6e370c-e015-47c9-9bff-b09bf3da7420",
			Name: "Adri",
		}

		body, err := json.Marshal(createUserReq)
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

	})
}
