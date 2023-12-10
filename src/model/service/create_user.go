package service

import (
	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"github.com/pedro-phd/first-api-go/src/configuration/rest_err"
	"github.com/pedro-phd/first-api-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	return nil
}
