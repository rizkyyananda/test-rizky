package repository

import (
	"Test-Rizky/domain"
	"fmt"
	"math"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(domain domain.Order) (domain.Order, error)
	FindAllAllWithFilter(customerName string, customerPhone string, productName string, page int64,
		limit int64) (result []domain.Order, totalData int64, currentPage int64, lastPage float64, err error)
	GetWhereId(id string) (domain.Order, error)
	Delete(id string) (domain.Order, error)
}

type orderConnection struct {
	connection *gorm.DB
}

//make instance
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) Save(add domain.Order) (domain.Order, error) {
	if err := db.connection.Save(&add).Error; err != nil {
		return add, err
	}
	return add, nil
}

func (db *orderConnection) FindAllAllWithFilter(customerName string, productName string, customerPhone string, page int64,
	limit int64) (result []domain.Order, totalData int64, currentPage int64, lastPage float64, err error) {
	var data []domain.Order
	var total int64

	sql := "Select * from t_order"

	if customerName != "" {
		sql += " where customer_name like '%" + customerName + "%'"
		if customerPhone != "" {
			sql += " and customer_phone like '%" + customerPhone + "%'"
		}
		if productName != "" {
			sql += " and product_name like '%" + productName + "%'"
		}
	} else {
		if productName != "" {
			sql += " where product_name like '%" + productName + "%'"
		}
	}
	db.connection.Count(&total)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	sql = fmt.Sprintf("%s limit %d offset %d", sql, limit, (page-1)*limit)
	lastPage = math.Ceil(float64(total / limit))
	if err := db.connection.Raw(sql).Scan(&data).Error; err != nil {
		return data, total, page, lastPage, err
	}
	return data, total, page, lastPage, nil
}

func (db *orderConnection) GetWhereId(id string) (domain.Order, error) {
	var data domain.Order
	if err := db.connection.Where("id=?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *orderConnection) Delete(id string) (domain.Order, error) {
	var data domain.Order
	if err := db.connection.Where("id=?", id).Delete(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
