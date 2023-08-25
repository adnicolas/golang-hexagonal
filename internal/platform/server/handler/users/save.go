package users

import (
	"net/http"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
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

func CreateSaveController(userRepository usuario.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req saveRequest
		// Pass it by reference (&)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		user, err := usuario.NewUser(req.Id, req.Name, req.Surname, req.Password, req.Email)
		if err != nil {
			/*if errors.Is(err, usuario.ErrInvalidUserId) {
				log.Println(err.Error())
			}*/
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := userRepository.Save(ctx, user); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Status(http.StatusCreated)
	}
}
