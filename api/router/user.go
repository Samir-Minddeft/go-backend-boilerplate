package router

import (
	"github.com/Samir-Minddeft/go-backend-boilerplate/api/controller"
	"github.com/Samir-Minddeft/go-backend-boilerplate/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {

	r.GET("/:id/get", middleware.AuthRole("admin"), controller.GetUser)
	r.GET("/list", middleware.AuthRole("admin"), controller.GetAllUsers)
	r.POST("/create", middleware.AuthRole("admin"), controller.CreateUser)
	r.PUT("/:id/update", middleware.AuthRole("admin"), controller.UpdateUser)
	r.DELETE("/:id/delete", middleware.AuthRole("admin"), controller.DeleteUser)
}
