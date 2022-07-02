package service

import (
	"Test-Rizky/dto"
	"Test-Rizky/jwt"
	logger "Test-Rizky/logger/data"
	"Test-Rizky/repository"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type AuthService interface {
	Login(loginDTO dto.LoginDTO) (bool, error)
	ValidateToken(tokenRequest string) (bool, string)
	GenerateToken(userName string) (dto.CustomerDTO, error)
}

type authService struct {
	customerRepository repository.CustomerRepository
}

//create new instance
func NewAuthService(
	customerRepository repository.CustomerRepository,
) AuthService {
	return &authService{
		customerRepository: customerRepository,
	}
}

func (service *authService) Login(loginDTO dto.LoginDTO) (bool, error) {
	getByUsername, err := service.customerRepository.GetByUserName(loginDTO.UserName)
	fmt.Println("getByUsername : ", getByUsername)
	if err != nil {
		return false, err
	}

	passHex := sha256.Sum256([]byte(loginDTO.Password))
	passResult := hex.EncodeToString(passHex[:])

	if passResult != getByUsername.Password {
		return false, errors.New("invalid password")
	}

	return true, nil
}

func (service *authService) GenerateToken(userName string) (dto.CustomerDTO, error) {
	fmt.Println("userName from client : ", userName)
	var customerDTO dto.CustomerDTO

	customerData, err := service.customerRepository.GetByUserName(userName)
	fmt.Println("customerData : ", customerData)
	if err != nil {
		return customerDTO, err
	}

	generatedToken, err := jwt.GenerateToken(customerData.UserName, customerData.Password)

	if err != nil {
		return customerDTO, err
	}

	customerData.Token = generatedToken

	saveData, err := service.customerRepository.Save(customerData)
	if err != nil {
		logger.Error("error save generate token", err)
		return customerDTO, err
	}
	logger.Info("response save generate token to client ", saveData)
	customerDTO.CustomerName = saveData.CustomerName
	customerDTO.CustomerPhone = saveData.CustomerPhone
	customerDTO.UserName = saveData.UserName
	customerDTO.Token = saveData.Token

	return customerDTO, nil
}

func (service *authService) ValidateToken(tokenRequest string) (bool, string) {
	var claimData jwt.Claims
	splittedAccessKey := strings.Split(tokenRequest, ".")
	bodyTokenBase64Value := splittedAccessKey[1]
	fmt.Println("ValidateAccessToken. Body token ", bodyTokenBase64Value)
	bodyTokenJsonByte, _ := base64.RawStdEncoding.DecodeString(bodyTokenBase64Value)
	bodyTokenJson := string(bodyTokenJsonByte)
	fmt.Println("ValidateAccessToken. Body token in string json ", bodyTokenJson)

	err := json.Unmarshal(bodyTokenJsonByte, &claimData)
	if err != nil {
		fmt.Println("Error while parsing json token body ", err)
		return false, err.Error()
	}

	fmt.Println("ValidateAccessToken. username from token ", claimData.Username)

	customerDataDB, err := service.customerRepository.GetByUserName(claimData.Username)
	fmt.Println("ValidateAccessToken. data from DB ", customerDataDB)
	if err != nil {
		return false, err.Error()
	}

	if customerDataDB.Token != tokenRequest {
		return false, "Access key not valid"
	}

	isValid, message := jwt.ValidateAccessToken(tokenRequest, customerDataDB.Password)

	return isValid, message
}
