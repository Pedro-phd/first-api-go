package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"github.com/pedro-phd/first-api-go/src/configuration/validation"
	"github.com/pedro-phd/first-api-go/src/controller/model/request"
	"github.com/pedro-phd/first-api-go/src/controller/model/response"
	"go.uber.org/zap"
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

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	response := response.UserResponse{
		ID:    "mock_id",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)

}
