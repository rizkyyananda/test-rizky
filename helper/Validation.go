package helpers

import (
	logger "Test-Rizky/logger/data"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidationBindJsonField(err error) interface{} {
	var ve validator.ValidationErrors
	var response Response
	if errors.As(err, &ve) {
		out := make([]AppError, len(ve))
		for i, fe := range ve {
			out[i] = AppError{402, nil, fe.Field(), getErrorMsg(fe)}
			response = ResponseError(getErrorMsg(fe), http.StatusBadRequest)
		}
	}
	return response
}

func ValidateField(data interface{}, field string, context *gin.Context) {
	logger.Info("Request data ", data)
	var response Response
	if data == nil || data == "" {
		logger.Info("masuk sini bos", data)
		response = ResponseError(field+" can't be empty", http.StatusBadRequest)
		context.JSON(http.StatusBadRequest, response)
		context.Abort()
	}

}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required " + fe.Field()
	case "lte":
		return "Should be less than" + fe.Field()
	case "gte":
		return "Should be greater than" + fe.Field()
	case "email":
		return "failed email format"
	case "unique":
		return fe.Field() + "already exist"
	}
	return "Unknown error"
}
