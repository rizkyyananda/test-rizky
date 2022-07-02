package controller

import (
	"Test-Rizky/dto"
	helpers "Test-Rizky/helper"
	logger "Test-Rizky/logger/data"
	"Test-Rizky/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController interface {
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type customerController struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &customerController{
		customerService: customerService,
	}
}

func (c customerController) Add(ctx *gin.Context) {
	//TODO implement me
	var customerDTO dto.CustomerDTO

	var response helpers.Response
	if err := ctx.ShouldBindJSON(&customerDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		result, err := c.customerService.Save(customerDTO)
		if err != nil {
			response = helpers.ResponseError(err.Error(), http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, response)
		} else {
			response = helpers.ResponseSuccess(result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c customerController) Update(ctx *gin.Context) {
	//TODO implement me
	//TODO implement me
	var customerDTO dto.CustomerDTO

	var response helpers.Response
	if err := ctx.ShouldBindJSON(&customerDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		result, err := c.customerService.Save(customerDTO)
		if err != nil {
			response = helpers.ResponseError(err.Error(), http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, response)
		} else {
			response = helpers.ResponseSuccess(result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c customerController) GetAllData(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	var pageDTO dto.PageDTO
	if err := ctx.ShouldBindJSON(&pageDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
	} else {
		result, err := c.customerService.GetAllData(pageDTO)
		if err != nil {
			response = helpers.ResponseError(err.Error(), http.StatusNotFound)
			ctx.JSON(http.StatusOK, response)
		} else {
			response = helpers.ResponseSuccess(result)
			logger.Info("Controller. Get List Donasi Product ", response)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (c customerController) GetDetail(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	id := ctx.Param("id")
	result, err := c.customerService.GetDetail(id)
	if err != nil {
		response = helpers.ResponseError(err.Error(), http.StatusNotFound)
		ctx.JSON(http.StatusConflict, response)
	} else {
		response = helpers.ResponseSuccess(result)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c customerController) Delete(ctx *gin.Context) {
	//TODO implement me
	var response helpers.Response
	id := ctx.Param("id")
	result := c.customerService.Delete(id)

	if result == "00" {
		response = helpers.ResponseSuccess("delete success")
	} else if result == "01" {
		response = helpers.ResponseError("data not found", http.StatusNotFound)
	} else {
		response = helpers.ResponseError("delete failed", http.StatusNotFound)
	}

	logger.Info("Controller. Delete", response)
	ctx.JSON(http.StatusOK, response)
}
