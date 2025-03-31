package server

import (
	"everything-template/internal/router/basic"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	basic.Register(r)
}
