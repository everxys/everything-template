package server

import (
	"everything-template/internal/vars"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	r := gin.Default()

	RegisterRoutes(r)

	addr := fmt.Sprintf(":%d", vars.Config.App.Port)
	log.Printf("Starting server on %s with env %s", addr, vars.Config.App.Env)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
