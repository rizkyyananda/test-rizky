package domain

import "time"

type Order struct {
	Id              string    `gorm:"PRIMARY_KEY; NOT NULL" json:"id"`
	CustomerName    string    `gorm:"type:varchar(25)" json:"customerName"`
	CustomerPhone   string    `gorm:"type:varchar(25)" json:"customerPhone"`
	TransactionDate time.Time `gorm:"type:datetime" json:"transactionDate"`
	ProductName     string    `gorm:"type:text;" json:"productName"`
	table           string    `gorm:"-"`
}

func (p Order) TableName() string {
	// double check here, make sure the table does exist!!
	if p.table != "" {
		return p.table
	}
	return "t_order" // default table name
}
