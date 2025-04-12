package router

import (
	"everything-template/pkg/response"
	"everything-template/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

type SleepRequest struct {
	Mill int `form:"mill" binding:"required,min=0"`
}

func RegisterTest(r *gin.Engine) {
	router := r.Group("/test")
	{
		router.GET("/sleep", func(c *gin.Context) {
			var req SleepRequest
			if err := util.Validate.BindAndValidate(c, &req); err != nil {
				response.BadRequestException(c, "error"+err.Error())
				return
			}
			time.Sleep(time.Duration(req.Mill) * time.Millisecond)
			response.SuccessJSON(c, "sleep done", nil)
		})
	}
}
