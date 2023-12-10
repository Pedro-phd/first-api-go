package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/first-api-go/src/controller"
)

// declare routes

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userID", controller.DeleteUser)

}
