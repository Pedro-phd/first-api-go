package service

import (
	"github.com/pedro-phd/first-api-go/src/configuration/rest_err"
	"github.com/pedro-phd/first-api-go/src/model"
)

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

// constructor - generate user service (use cases)
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
