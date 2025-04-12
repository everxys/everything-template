package server

import (
	"everything-template/internal/router"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, middleware ...gin.HandlerFunc) {
	router.RegisterBasic(r)
	router.RegisterAuth(r)
	router.RegisterTest(r)
	router.RegisterUser(r, middleware...)
}
