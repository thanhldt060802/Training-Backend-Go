package handler

import (
	"net/http"
	"strconv"
	"training-project/internal/dto"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService service.InvoiceService
}

func NewInvoiceHandler(route *gin.RouterGroup, invoiceService service.InvoiceService) *InvoiceHandler {
	invoiceHandler := &InvoiceHandler{invoiceService: invoiceService}
	invoiceAPI := route.Group("/invoices")
	{
		invoiceAPI.GET("/", invoiceHandler.GetAllUsers)
		invoiceAPI.GET("/id/:id", invoiceHandler.GetInvoiceByID)
		invoiceAPI.GET("/user/:userId", invoiceHandler.GetAllInvoicesByUserId)
		invoiceAPI.POST("/", invoiceHandler.CreateInvoice)
		invoiceAPI.PUT("/id/:id", invoiceHandler.UpdateInvoiceById)
		invoiceAPI.DELETE("/id/:id", invoiceHandler.DeleteInvoiceById)
	}
	return invoiceHandler
}

func (invoiceHandler *InvoiceHandler) GetAllUsers(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	invoices, err := invoiceHandler.invoiceService.FindAllInvoices(rqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToInvoiceResDTOs(invoices),
	})
}

func (invoiceHandler *InvoiceHandler) GetInvoiceByID(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	invoice, err := invoiceHandler.invoiceService.FindInvoiceById(rqCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToInvoiceResDTO(invoice),
	})
}

func (invoiceHandler *InvoiceHandler) GetAllInvoicesByUserId(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "userId is not valid",
		})
		return
	}
	invoices, err := invoiceHandler.invoiceService.FindAllInvoicesByUserId(rqCtx, int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": dto.ToInvoiceResDTOs(invoices),
	})
}

func (invoiceHandler *InvoiceHandler) CreateInvoice(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	var invoiceCreateReqDTO dto.InvoiceCreateReqDTO
	if err := ctx.ShouldBind(&invoiceCreateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := invoiceHandler.invoiceService.CreateInvoice(rqCtx, &invoiceCreateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": " create invoice successfully",
	})
}

func (invoiceHandler *InvoiceHandler) UpdateInvoiceById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	var invoiceUpdateReqDTO dto.InvoiceUpdateReqDTO
	if err := ctx.ShouldBind(&invoiceUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := invoiceHandler.invoiceService.UpdateInvoiceById(rqCtx, int64(id), &invoiceUpdateReqDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update invoice successfully",
	})
}

func (invoiceHandler *InvoiceHandler) DeleteInvoiceById(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not valid",
		})
		return
	}
	if err := invoiceHandler.invoiceService.DeleteInvoiceById(rqCtx, int64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete invoice successfully",
	})
}
