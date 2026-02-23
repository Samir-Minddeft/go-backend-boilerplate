package router

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := r.Group("/api/v1")
	userGroup := api.Group("/user")
	authGroup := api.Group("/auth")

	UserRoute(userGroup)
	AuthRoute(authGroup)

}
