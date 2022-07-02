package service

import (
	"Test-Rizky/domain"
	"Test-Rizky/dto"
	logger "Test-Rizky/logger/data"
	"Test-Rizky/repository"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type CustomerService interface {
	Save(dto dto.CustomerDTO) (interface{}, error)
	GetAllData(dto dto.PageDTO) (result dto.PageDTO, err error)
	GetDetail(id string) (domain.Customer, error)
	Delete(id string) string
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

//create new instance

func NewCustomerService(check repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: check,
	}
}

func (repo customerService) Save(dto dto.CustomerDTO) (interface{}, error) {
	//TODO implement me
	data := domain.Customer{}

	checkOrder, err := repo.customerRepository.GetWhereId(dto.Id)
	if err != nil {
		return checkOrder, err
	}

	passHex := sha256.Sum256([]byte(dto.Password))
	passResult := hex.EncodeToString(passHex[:])

	if checkOrder.Id != "" {
		data.Id = checkOrder.Id
		data.CustomerPhone = dto.CustomerPhone
		data.CustomerName = dto.CustomerName
		data.CustomerAddress = dto.CustomerAddress
		data.UserName = dto.UserName
		data.Password = passResult
	} else {
		data.Id = uuid.NewString()
		data.CustomerPhone = dto.CustomerPhone
		data.CustomerName = dto.CustomerName
		data.CustomerAddress = dto.CustomerAddress
		data.UserName = dto.UserName
		data.Password = passResult
	}

	data, err = repo.customerRepository.Save(data)

	if err != nil {
		return data, err
	}

	return data, err
}

func (repo customerService) GetAllData(dataDto dto.PageDTO) (result dto.PageDTO, err error) {
	//TODO implement me

	var detailData = dto.CustomerDTO{}

	dataByte, _ := json.Marshal(dataDto.Data)
	err = json.Unmarshal(dataByte, &detailData)
	if err != nil {
		return result, err
	}
	data, total, page, lastPage, err := repo.customerRepository.FindAllAllWithFilter(detailData.CustomerName, detailData.CustomerPhone, dataDto.Page, dataDto.Limit)
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

func (repo customerService) GetDetail(id string) (domain.Customer, error) {
	//TODO implement me
	data, err := repo.customerRepository.GetWhereId(id)
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

func (repo customerService) Delete(id string) string {
	//TODO implement me
	data, err := repo.customerRepository.GetWhereId(id)
	if err != nil {
		logger.Error("Delete", errors.New("Error get data "+err.Error()))
		return "01"
	}
	logger.Info("GetDetailData", data)

	if data.Id == "" {
		logger.Error("GetDetailData", errors.New("data not found"))
		return "01"

	} else {
		result, err := repo.customerRepository.Delete(id)
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
