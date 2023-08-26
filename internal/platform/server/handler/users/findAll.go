package users

import (
	"net/http"

	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/adnicolas/golang-hexagonal/kit/bus"
	"github.com/gin-gonic/gin"
)

func FindAllController(myBus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		elements, err := myBus.DispatchQuery(ctx, fetching.NewUserQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, elements)
	}
}
