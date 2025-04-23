package handler

import (
	"net/http"
	"strconv"
	"training-project/internal/dto"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
)

type CartItemHandler struct {
	CartItemService service.CartItemService
}

func NewCartItemHandler(route *gin.RouterGroup, cartItemService service.CartItemService) *CartItemHandler {
	cartItemHandler := &CartItemHandler{CartItemService: cartItemService}
	cartItemAPI := route.Group("/cart-items")
	{
		cartItemAPI.GET("/", cartItemHandler.GetAllCartItems)
		cartItemAPI.GET("/id/:id", cartItemHandler.GetCartItemById)
		cartItemAPI.GET("/user/:userId", cartItemHandler.GetAllCartItemsByUserId)
		cartItemAPI.POST("/", cartItemHandler.CreateCartItem)
		cartItemAPI.PUT("/id/:id", cartItemHandler.UpdateCartItemById)
		cartItemAPI.DELETE("/id/:id", cartItemHandler.DeleteCartItemById)
	}
	return cartItemHandler
}

func (cartItemHandler *CartItemHandler) GetAllCartItems(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	cartItems, err := cartItemHandler.CartItemService.FindAllCartItems(rqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToCartItemResDTOs(cartItems),
	})
}

func (cartItemHandler *CartItemHandler) GetCartItemById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	cartItem, err := cartItemHandler.CartItemService.FindCartItemById(rqCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToCartItemResDTO(cartItem),
	})
}

func (cartItemHandler *CartItemHandler) GetAllCartItemsByUserId(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "userId is not valid",
		})
		return
	}
	cartItems, err := cartItemHandler.CartItemService.FindAllCartItemsByUserId(rqCtx, int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToCartItemResDTOs(cartItems),
	})
}

func (cartItemHandler *CartItemHandler) CreateCartItem(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	var cartItemCreateReqDTO dto.CartItemCreateReqDTO
	if err := ctx.ShouldBind(&cartItemCreateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := cartItemHandler.CartItemService.CreateCartItem(rqCtx, &cartItemCreateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create cart item successful",
	})
}

func (cartItemHandler *CartItemHandler) UpdateCartItemById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	var cartItemUpdateReqDTO dto.CartItemUpdateReqDTO
	if err := ctx.ShouldBind(&cartItemUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := cartItemHandler.CartItemService.UpdateCartItemById(rqCtx, int64(id), &cartItemUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update cart item successful",
	})
}

func (cartItemHandler *CartItemHandler) DeleteCartItemById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	if err := cartItemHandler.CartItemService.DeleteCartItemById(rqCtx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete cart item successful",
	})
}
