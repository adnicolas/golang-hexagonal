package users

import (
	"net/http"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/gin-gonic/gin"
)

func CreateFindAllController(userRepository usuario.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		elements, err := userRepository.FindAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, elements)
	}
}