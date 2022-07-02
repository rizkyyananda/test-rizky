package controller

import (
	"Test-Rizky/dto"
	helpers "Test-Rizky/helper"
	logger "Test-Rizky/logger/data"
	"Test-Rizky/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController interface {
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
	authService  service.AuthService
}

func NewOrderController(orderService service.OrderService, authService service.AuthService) OrderController {
	return &orderController{
		orderService: orderService,
		authService:  authService,
	}
}

func (c orderController) Add(ctx *gin.Context) {
	//TODO implement me
	var orderDTO dto.OrderDTO

	var response helpers.Response
	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		token := ctx.Request.Header.Get("Access-Token")
		if token != "" {
			validateToken, err := c.authService.ValidateToken(token)
			if validateToken {
				result, err := c.orderService.Save(orderDTO)
				if err != nil {
					response = helpers.ResponseError(err.Error(), http.StatusBadRequest)
					ctx.JSON(http.StatusBadRequest, response)
				} else {
					response = helpers.ResponseSuccess(result)
					ctx.JSON(http.StatusOK, response)
				}
			} else {
				response = helpers.ResponseError(err, http.StatusUnauthorized)
				ctx.JSON(http.StatusOK, response)
			}
		} else {
			response = helpers.ResponseError("token not valid", http.StatusUnauthorized)
			ctx.JSON(http.StatusOK, response)
		}

	}
}

func (c orderController) Update(ctx *gin.Context) {
	//TODO implement me
	//TODO implement me
	var orderDTO dto.OrderDTO

	var response helpers.Response
	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		token := ctx.Request.Header.Get("Access-Token")
		if token != "" {
			validateToken, err := c.authService.ValidateToken(token)
			if validateToken {
				result, err := c.orderService.Save(orderDTO)
				if err != nil {
					response = helpers.ResponseError(err.Error(), http.StatusBadRequest)
					ctx.JSON(http.StatusBadRequest, response)
				} else {
					response = helpers.ResponseSuccess(result)
					ctx.JSON(http.StatusOK, response)
				}
			} else {
				response = helpers.ResponseError(err, http.StatusUnauthorized)
				ctx.JSON(http.StatusOK, response)
			}
		} else {
			response = helpers.ResponseError("token not valid", http.StatusUnauthorized)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c orderController) GetAllData(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	var pageDTO dto.PageDTO
	token := ctx.Request.Header.Get("Access-Token")
	if err := ctx.ShouldBindJSON(&pageDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		if token != "" {
			validateToken, err := c.authService.ValidateToken(token)
			if validateToken {
				result, err := c.orderService.GetAllData(pageDTO)
				if err != nil {
					response = helpers.ResponseError(err.Error(), http.StatusNotFound)
				} else {
					response = helpers.ResponseSuccess(result)
				}
			} else {
				response = helpers.ResponseError(err, http.StatusNotFound)
			}
		} else {
			response = helpers.ResponseError("token not valid", http.StatusNotFound)
		}

		logger.Info("Controller. Get List Donasi Product ", response)
		ctx.JSON(http.StatusOK, response)
	}

}

func (c orderController) GetDetail(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	id := ctx.Param("id")
	token := ctx.Request.Header.Get("Access-Token")
	if token != "" {
		validateToken, err := c.authService.ValidateToken(token)
		if validateToken {
			result, err := c.orderService.GetDetail(id)
			if err != nil {
				response = helpers.ResponseError(err.Error(), http.StatusNotFound)
				ctx.JSON(http.StatusConflict, response)
			} else {
				response = helpers.ResponseSuccess(result)
				ctx.JSON(http.StatusOK, response)
			}
		} else {
			response = helpers.ResponseError(err, http.StatusForbidden)
			ctx.JSON(http.StatusForbidden, response)
		}
	} else {
		response = helpers.ResponseError("token not valid", http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, response)
	}

}

func (c orderController) Delete(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	id := ctx.Param("id")
	result := c.orderService.Delete(id)

	token := ctx.Request.Header.Get("Access-Token")
	if token != "" {
		validateToken, err := c.authService.ValidateToken(token)
		if validateToken {
			if result == "00" {
				response = helpers.ResponseSuccess("delete success")
			} else if result == "01" {
				response = helpers.ResponseError("data not found", http.StatusNotFound)
			} else {
				response = helpers.ResponseError("delete failed", http.StatusNotFound)
			}
		} else {
			response = helpers.ResponseError(err, http.StatusForbidden)
			ctx.JSON(http.StatusForbidden, response)
		}
	} else {
		response = helpers.ResponseError("token not valid", http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, response)
	}

	logger.Info("Controller. Delete", response)
	ctx.JSON(http.StatusOK, response)
}
