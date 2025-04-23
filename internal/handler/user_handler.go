package handler

import (
	"context"
	"net/http"
	"strconv"
	"training-project/internal/dto"
	"training-project/internal/service"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(api huma.API, userService service.UserService) *UserHandler {
	userHandler := &UserHandler{userService: userService}
	// userAPI := route.Group("/users")
	// {
	// 	userAPI.GET("/", userHandler.GetAllUsers)
	// 	userAPI.GET("/id/:id", userHandler.GetUserByID)
	// 	userAPI.GET("/username/:username", userHandler.GetUserByUsername)
	// 	userAPI.POST("/", userHandler.CreateUser)
	// 	userAPI.PUT("/id/:id", userHandler.UpdateUserById)
	// 	userAPI.DELETE("/id/:id", userHandler.DeleteUserById)
	// }
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users",
		Summary:     "Sumary ...",
		Description: "Description ...",
	}, userHandler.GetAllUsers)
	return userHandler
}

func (userHandler *UserHandler) GetAllUsers(ctx context.Context, req *struct{}) (*struct {
	Body struct {
		Data []dto.UserResDTO `json:"data"`
	}
}, error) {
	users, err := userHandler.userService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// res := &dto.GetAllUsersResponse{}
	// res.Body.Data = dto.ToUserResDTOs(users)
	// return res, nil

	res := struct {
		Body struct {
			Data []dto.UserResDTO `json:"data"`
		}
	}{}
	res.Body.Data = dto.ToUserResDTOs(users)
	return &res, nil
}

func (userHandler *UserHandler) GetUserByID(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
	}
	user, err := userHandler.userService.FindUserById(rqCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToUserResDTO(user),
	})
}

func (userHandler *UserHandler) GetUserByUsername(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	username := ctx.Param("username")
	user, err := userHandler.userService.FindUserByUsername(rqCtx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToUserResDTO(user),
	})
}

func (userHandler *UserHandler) CreateUser(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	var userCreateReqDTO dto.UserCreateReqDTO
	if err := ctx.ShouldBind(&userCreateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := userHandler.userService.CreateUser(rqCtx, &userCreateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create user successful",
	})
}

func (userHandler *UserHandler) UpdateUserById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
	}
	var userUpdateReqDTO dto.UserUpdateReqDTO
	if err := ctx.ShouldBind(&userUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := userHandler.userService.UpdateUserById(rqCtx, int64(id), &userUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update user successful",
	})
}

func (userHandler *UserHandler) DeleteUserById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
	}
	if err := userHandler.userService.DeleteUserById(rqCtx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete user successful",
	})
}
