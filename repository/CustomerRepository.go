package repository

import (
	"Test-Rizky/domain"
	"fmt"
	"math"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(domain domain.Customer) (domain.Customer, error)
	FindAllAllWithFilter(customerName string, customerPhone string, page int64,
		limit int64) (result []domain.Customer, totalData int64, currentPage int64, lastPage float64, err error)
	GetWhereId(id string) (domain.Customer, error)
	GetByUserName(userName string) (domain.Customer, error)
	Delete(id string) (domain.Customer, error)
}

type customerConnection struct {
	connection *gorm.DB
}

//make instance
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) Save(add domain.Customer) (domain.Customer, error) {
	if err := db.connection.Save(&add).Error; err != nil {
		return add, err
	}
	return add, nil
}

func (db *customerConnection) FindAllAllWithFilter(customerName string, customerPhone string, page int64,
	limit int64) (result []domain.Customer, totalData int64, currentPage int64, lastPage float64, err error) {
	var data []domain.Customer
	var total int64

	sql := "Select * from m_customer"

	if customerName != "" {
		sql += " where customer_name like '%" + customerName + "%'"
		if customerPhone != "" {
			sql += " and customer_phone like '%" + customerPhone + "%'"
		}
	}

	if customerPhone != "" {
		sql += " where customer_phone like '%" + customerPhone + "%'"
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

func (db *customerConnection) GetWhereId(id string) (domain.Customer, error) {
	var data domain.Customer
	if err := db.connection.Where("id=?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *customerConnection) GetByUserName(userName string) (domain.Customer, error) {
	var data domain.Customer
	if err := db.connection.Where("user_name=?", userName).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (db *customerConnection) Delete(id string) (domain.Customer, error) {
	var data domain.Customer
	if err := db.connection.Where("id=?", id).Delete(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
