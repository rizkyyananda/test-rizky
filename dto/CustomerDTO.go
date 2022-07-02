package dto

type CustomerDTO struct {
	Id              string `json:"id" form:"id"`
	CustomerName    string `json:"customerName" from:"customerName"`
	CustomerPhone   string `json:"customerPhone" from:"customerPhone"`
	CustomerAddress string `json:"customerAddress" from:"customerAddress"`
	UserName        string `json:"userName" from:"userName" binding:"required"`
	Password        string `json:"password" from:"password" binding:"required"`
	Token           string `json:"token" from:"token"`
}
