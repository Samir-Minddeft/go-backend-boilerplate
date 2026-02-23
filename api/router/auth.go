package router

import (
	"github.com/Samir-Minddeft/go-backend-boilerplate/api/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/login", controller.Login)
}
