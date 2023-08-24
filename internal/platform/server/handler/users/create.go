package users

import (
	"net/http"

	usuario "github.com/adnicolas/golang-hexagonal/internal/platform"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	// TODO: Usar tipo uuid para robustecer validaciones
	// binding property (validation) offered by Gin
	Uuid     string/*uuid.UUID*/ `json:"uuid" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// TODO: Investigar relaciones entre entidades (1:1, 1:n, m:n)
	//roleId: RoleEnum;
}

func CreateController(userRepository usuario.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		// Pass it by reference (&)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		user := usuario.NewUser(req.Uuid, req.Name, req.Surname, req.Password, req.Email)
		if err := userRepository.Save(ctx, user); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Status(http.StatusCreated)
	}
}
