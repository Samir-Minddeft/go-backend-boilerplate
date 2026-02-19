package routes

import (
	"github.com/Samir-Minddeft/go-backend-boilerplate/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/user/:id/get", controller.GetUser)
	router.GET("/user/list", controller.GetAllUsers)
	router.POST("/user/create", controller.CreateUser)
	router.PUT("/user/:id/update", controller.UpdateUser)
	router.DELETE("/user/:id/delete", controller.DeleteUser)

	return router
}
