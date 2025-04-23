package handler

import (
	"net/http"
	"strconv"
	"training-project/internal/dto"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(route *gin.RouterGroup, categoryService service.CategoryService) *CategoryHandler {
	categoryHandler := &CategoryHandler{categoryService: categoryService}
	categoryAPI := route.Group("/categories")
	{
		categoryAPI.GET("/", categoryHandler.GetAllCategories)
		categoryAPI.GET("/id/:id", categoryHandler.GetCategoryByID)
		categoryAPI.POST("/", categoryHandler.CreateCategory)
		categoryAPI.PUT("/id/:id", categoryHandler.UpdateCategoryById)
		categoryAPI.DELETE("/id/:id", categoryHandler.DeleteCategoryById)
	}
	return categoryHandler
}

func (categoryHandler *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	categories, err := categoryHandler.categoryService.FindAllCategories(rqCtx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToCategoryResDTOs(categories),
	})
}

func (categoryHandler *CategoryHandler) GetCategoryByID(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	category, err := categoryHandler.categoryService.FindCategoryById(rqCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToCategoryResDTO(category),
	})
}

func (categoryHandler *CategoryHandler) CreateCategory(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	var categoryCreateReqDTO dto.CategoryCreateReqDTO
	if err := ctx.ShouldBind(&categoryCreateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := categoryHandler.categoryService.CreateCategory(rqCtx, &categoryCreateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create category successful",
	})
}

func (categoryHandler *CategoryHandler) UpdateCategoryById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	var categoryUpdateReqDTO dto.CategoryUpdateReqDTO
	if err := ctx.ShouldBind(&categoryUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := categoryHandler.categoryService.UpdateCategoryById(rqCtx, int64(id), &categoryUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update category successful",
	})
}

func (categoryHandler *CategoryHandler) DeleteCategoryById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	if err := categoryHandler.categoryService.DeleteCategoryById(rqCtx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete category successful",
	})
}
