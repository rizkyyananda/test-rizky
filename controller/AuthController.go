package controller

import (
	"Test-Rizky/dto"
	helpers "Test-Rizky/helper"
	"Test-Rizky/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	GenerateToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(check service.AuthService) AuthController {
	return &authController{
		authService: check,
	}
}

func (controller *authController) GenerateToken(context *gin.Context) {
	var loginDTO dto.LoginDTO
	var response helpers.Response
	if err := context.ShouldBindJSON(&loginDTO); err != nil {
		var data = helpers.ValidationBindJsonField(err)
		context.JSON(http.StatusBadRequest, data)
	} else {
		isValidLogin, _ := controller.authService.Login(loginDTO)
		if isValidLogin {
			result, status := controller.authService.GenerateToken(loginDTO.UserName)
			if status != nil {
				response = helpers.ResponseError(status.Error(), http.StatusBadRequest)
				context.JSON(http.StatusBadRequest, response)
			} else {
				response = helpers.ResponseSuccess(result)
				context.JSON(http.StatusOK, response)
			}
		} else {
			response = helpers.ResponseError("Wrong combination Username and Password", http.StatusUnauthorized)
			context.JSON(http.StatusUnauthorized, response)
		}
	}
}
