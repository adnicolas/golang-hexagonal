package users

import (
	"net/http"

	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/adnicolas/golang-hexagonal/kit/query"
	"github.com/gin-gonic/gin"
)

func FindAllController(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		elements, err := queryBus.Dispatch(ctx, fetching.NewUserQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, elements)
	}
}
