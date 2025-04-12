package server

import (
	"fmt"

	"everything-template/internal/middleware"
	"everything-template/internal/vars"
	"everything-template/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	RegisterRoutes(r, middleware.AuthMiddleware())

	addr := fmt.Sprintf(":%d", vars.Config.App.Port)
	logger.Infow(fmt.Sprintf("Starting server on %s with env %s", addr, vars.Config.App.Env))

	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
