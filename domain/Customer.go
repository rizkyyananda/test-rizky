package domain

type Customer struct {
	Id              string `gorm:"PRIMARY_KEY; NOT NULL" json:"id"`
	CustomerName    string `gorm:"type:varchar(25)" json:"customerName"`
	CustomerPhone   string `gorm:"type:varchar(15); unique" json:"customerPhone"`
	CustomerAddress string `gorm:"type:text;" json:"customerAddress"`
	UserName        string `gorm:"type:varchar(25); unique" json:"userName"`
	Password        string `gorm:"type:text;" json:"password"`
	Token           string `gorm:"type:text;" json:"token"`
	table           string `gorm:"-"`
}

func (p Customer) TableName() string {
	// double check here, make sure the table does exist!!
	if p.table != "" {
		return p.table
	}
	return "m_customer" // default table name
}
