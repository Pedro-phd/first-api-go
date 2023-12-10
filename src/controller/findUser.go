package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/configuration/rest_err"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {

	err := rest_err.NewBadRequestError("Not exist this route")
	c.JSON(err.Code, err)

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {}
