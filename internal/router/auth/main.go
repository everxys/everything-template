package auth

import (
	"everything-template/internal/app/controller/authcontroller"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	router := r.Group("/auth")
	{
		router.POST("/register", authcontroller.Register)
		router.POST("/login", authcontroller.Login)
	}
}
