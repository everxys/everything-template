package server

import (
	"everything-template/internal/router/auth"
	"everything-template/internal/router/basic"
	"everything-template/internal/router/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, middleware ...gin.HandlerFunc) {
	basic.Register(r)
	auth.Register(r)
	user.Register(r, middleware...)
}
