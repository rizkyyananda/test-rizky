package dto

type PageDTO struct {
	Page      int64       `json:"page" form:"page"`
	Limit     int64       `json:"limit" form:"limit"`
	LastPage  float64     `json:"lastPage" form:"lastPage"`
	Data      interface{} `json:"data" form:"data"`
	TotalData int64       `json:"totalData" form:"totalData"`
}
