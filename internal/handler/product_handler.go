package handler

import (
	"net/http"
	"strconv"
	"training-project/internal/dto"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(route *gin.RouterGroup, productService service.ProductService) *ProductHandler {
	productHandler := &ProductHandler{productService: productService}
	productAPI := route.Group("/products")
	{
		productAPI.GET("/", productHandler.GetAllProducts)
		productAPI.GET("/id/:id", productHandler.GetProductByID)
		productAPI.POST("/", productHandler.CreateProduct)
		productAPI.PUT("/id/:id", productHandler.UpdateProductById)
		productAPI.DELETE("/id/:id", productHandler.DeleteProductById)
	}
	return productHandler
}

func (productHandler *ProductHandler) GetAllProducts(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	products, err := productHandler.productService.FindAllProducts(rqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToProductResDTOs(products),
	})
}

func (productHandler *ProductHandler) GetProductByID(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	product, err := productHandler.productService.FindProductById(rqCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToProductResDTO(product),
	})
}

func (productHandler *ProductHandler) CreateProduct(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	var productCreateReq dto.ProductCreateReqDTO
	if err := ctx.ShouldBind(&productCreateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := productHandler.productService.CreateProduct(rqCtx, &productCreateReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product create successful",
	})
}

func (productHandler *ProductHandler) UpdateProductById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	var productUpdateReq dto.ProductUpdateReqDTO
	if err := ctx.ShouldBind(&productUpdateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := productHandler.productService.UpdateProductById(rqCtx, int64(id), &productUpdateReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product update successful",
	})
}

func (productHandler *ProductHandler) DeleteProductById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	if err := productHandler.productService.DeleteProductById(rqCtx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product delete successful",
	})
}
