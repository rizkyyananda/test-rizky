package helpers

type Response struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
	Info   interface{} `json:"info"`
}

type AppError struct {
	Code        int         `json:"code,omitempty"`
	Object      interface{} `json:"object"`
	Field       interface{} `json:"field"`
	MessageData interface{} `json:"messageData"`
}

//ResponseSuccess
func ResponseSuccess(data interface{}) Response {
	res := Response{
		Status: true,
		Data:   data,
		Error:  nil,
		Info:   nil,
	}
	return res
}

//ResponseError
func ResponseError(messageData interface{}, code int) Response {
	error := AppError{
		Code:        code,
		Object:      nil,
		Field:       nil,
		MessageData: messageData,
	}
	res := Response{
		Status: false,
		Error:  error,
	}
	return res
}
