package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/configuration/rest_err"
)

func FindUserByID(c *gin.Context) {

	err := rest_err.NewBadRequestError("Not exist this route")
	c.JSON(err.Code, err)

}

func FindUserByEmail(c *gin.Context) {}
