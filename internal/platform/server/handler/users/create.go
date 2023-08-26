package users

import (
	"errors"
	"net/http"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/creating"
	"github.com/adnicolas/golang-hexagonal/kit/command"
	"github.com/gin-gonic/gin"
)

type saveRequest struct {
	// binding property (validation) offered by Gin
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// TODO: Investigar relaciones entre entidades (1:1, 1:n, m:n)
	//roleId: RoleEnum;
}

func CreateController(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req saveRequest
		// Pass it by reference (&)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewUserCommand(
			req.Id,
			req.Name,
			req.Surname,
			req.Password,
			req.Email,
		))

		if err != nil {
			switch {
			case errors.Is(err, usuario.ErrInvalidUserId):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
		ctx.Status(http.StatusCreated)
	}
}
