package server

import (
	"context"
	"errors"
	"everything-template/internal/middleware"
	"everything-template/internal/vars"
	"everything-template/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	RegisterRoutes(r, middleware.AuthMiddleware())

	addr := fmt.Sprintf(":%d", vars.Config.App.Port)
	logger.Infow(fmt.Sprintf("Starting server on %s with env %s", addr, vars.Config.App.Env))

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorw("listen error", "err", err)
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infow("Shutting down server...")

	timeout := vars.Config.App.GracefulTimeout
	if timeout <= 0 {
		timeout = 5
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorw("Server forced to shutdown", "err", err)
	}

	logger.Infow("Server exiting")
}
