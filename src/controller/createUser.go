package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"github.com/pedro-phd/first-api-go/src/configuration/validation"
	"github.com/pedro-phd/first-api-go/src/controller/model/request"
	"github.com/pedro-phd/first-api-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserResquest

	// shouldBindJSON create a object from body of request,
	// if validation of fields it is ok (binding in request struct)

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.String(http.StatusCreated, "")

}
