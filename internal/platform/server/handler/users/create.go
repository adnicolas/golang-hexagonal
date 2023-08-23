package users

import (
	"context"
	"net/http"

	mooc "github.com/adnicolas/golang-hexagonal/internal/platform"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRequest struct {
	// binding property (validation) offered by Gin
	Uuid     uuid.UUID `json:"uuid" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Surname  string    `json:"surname" binding:"required"`
	//roleId: RoleEnum;
}

func CreateController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		// Pass it by reference (&)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		user := mooc.NewUser(req.Uuid, req.Email, req.Name, req.Surname, req.Password)
		Save(ctx, user)
		/*if err := Save(ctx, user); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}*/
		ctx.Status(http.StatusCreated)
	}
}

// Save persist the user on the DB
func Save(ctx context.Context, user mooc.User) {}
