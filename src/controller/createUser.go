package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/configuration/validation"
	"github.com/pedro-phd/first-api-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserResquest

	// shouldBindJSON create a object from body of request,
	// if validation of fields it is ok (binding in request struct)

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
