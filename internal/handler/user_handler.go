package handler

import (
	"net/http"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userServ service.UserService
}

func NewUserHandler(userServ service.UserService) *UserHandler {
	return &UserHandler{userServ: userServ}
}

func (userHandler *UserHandler) GetAllUsers(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	users, err := userHandler.userServ.GetAllUsers(rqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
