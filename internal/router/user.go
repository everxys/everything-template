package router

import (
	"everything-template/internal/app/controller/userController"
	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.Engine, middleware ...gin.HandlerFunc) {
	router := r.Group("/user")
	{
		router.Use(middleware...)

		router.GET("/", userController.User)
		router.POST("/logout", userController.Logout)
	}
}
