package dto

import "time"

type OrderDTO struct {
	Id              string    `json:"id" form:"id"`
	CustomerName    string    `json:"customerName" from:"customerName"`
	CustomerPhone   string    `json:"customerPhone" from:"customerPhone"`
	TransactionDate time.Time `json:"transactionDate" from:"transactionDate"`
	ProductName     string    `json:"productName" from:"productName"`
}
