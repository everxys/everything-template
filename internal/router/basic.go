package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBasic(r *gin.Engine) {
	router := r.Group("/")
	{
		router.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		router.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}
}
