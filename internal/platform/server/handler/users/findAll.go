package users

import (
	"net/http"

	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/gin-gonic/gin"
)

func FindAllController(fetchingUserService fetching.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		elements, err := fetchingUserService.FindAllUsers(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, elements)
	}
}
