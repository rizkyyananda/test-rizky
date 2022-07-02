package service

import (
	"Test-Rizky/domain"
	"Test-Rizky/dto"
	logger "Test-Rizky/logger/data"
	"Test-Rizky/repository"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type OrderService interface {
	Save(dto dto.OrderDTO) (interface{}, error)
	GetAllData(dataDto dto.PageDTO) (result dto.PageDTO, err error)
	GetDetail(id string) (domain.Order, error)
	Delete(id string) string
}

type orderService struct {
	orderRepository repository.OrderRepository
}

//create new instance

func NewOrderService(check repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: check,
	}
}

func (repo orderService) Save(dto dto.OrderDTO) (interface{}, error) {
	//TODO implement me
	data := domain.Order{}

	checkOrder, err := repo.orderRepository.GetWhereId(dto.Id)
	if err != nil {
		return checkOrder, err
	}

	dt := time.Now()
	timeStart := dt.Format("2006-01-02 15:04:05")

	date, _ := time.Parse("2006-01-02 15:4:5", timeStart)

	if checkOrder.Id != "" {
		data.Id = checkOrder.Id
		data.CustomerPhone = dto.CustomerPhone
		data.CustomerName = dto.CustomerName
		data.TransactionDate = date
		data.ProductName = dto.ProductName
	} else {
		data.Id = uuid.NewString()
		data.CustomerPhone = dto.CustomerPhone
		data.CustomerName = dto.CustomerName
		data.TransactionDate = date
		data.ProductName = dto.ProductName
	}

	data, err = repo.orderRepository.Save(data)

	if err != nil {
		return data, err
	}

	return data, err
}

func (repo orderService) GetAllData(dataDto dto.PageDTO) (result dto.PageDTO, err error) {
	//TODO implement me
	var detailData = dto.OrderDTO{}

	dataByte, _ := json.Marshal(dataDto.Data)
	err = json.Unmarshal(dataByte, &detailData)
	if err != nil {
		return result, err
	}
	data, total, page, lastPage, err := repo.orderRepository.FindAllAllWithFilter(detailData.CustomerName, detailData.CustomerPhone, detailData.ProductName, dataDto.Page, dataDto.Limit)
	if err != nil {
		logger.Error("GetAllData", errors.New("Error get data "+err.Error()))
		return result, err
	}
	result.Page = page
	result.Data = data
	result.TotalData = total
	result.LastPage = lastPage

	return result, nil
}

func (repo orderService) GetDetail(id string) (domain.Order, error) {
	//TODO implement me
	data, err := repo.orderRepository.GetWhereId(id)
	if err != nil {
		logger.Error("GetDetail", errors.New("Error get data "+err.Error()))
		return data, err
	}

	if data.Id == "" {
		logger.Error("GetDetailOrder", errors.New("data not found"))
		return data, errors.New("data not found")
	} else {
		return data, nil
	}
}

func (repo orderService) Delete(id string) string {
	//TODO implement me
	data, err := repo.orderRepository.GetWhereId(id)
	if err != nil {
		logger.Error("Delete", errors.New("Error get data "+err.Error()))
		return "01"
	}
	logger.Info("GetDetailData", data)

	if data.Id == "" {
		logger.Error("GetDetailData", errors.New("data not found"))
		return "01"

	} else {
		result, err := repo.orderRepository.Delete(id)
		if err != nil {
			logger.Error("Delete", errors.New("Error delete data "+err.Error()))
			return "01"
		}
		if result.Id == "" {
			logger.Info("delete ", result)
			return "00"
		} else {
			return "02"
		}
	}
}
