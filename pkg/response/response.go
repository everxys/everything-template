package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	ErrCode   Code        `json:"errcode"`
	RequestID string      `json:"requestid"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// JSON 基础返回
func JSON(ctx *gin.Context, status int, errcode Code, message string, data interface{}) {
	if message == "" {
		message = CodeMap[errcode]
	}
	ctx.JSON(status, JSONResponse{
		ErrCode:   errcode,
		Message:   message,
		RequestID: ctx.GetHeader("X-Request-Id"), // Gin uses GetHeader for requests
		Data:      data,
	})
}

// SuccessJSON 成功返回
func SuccessJSON(ctx *gin.Context, message string, data interface{}) {
	if message == "" {
		message = Success.Msg()
	}
	JSON(ctx, http.StatusOK, Success, message, data)
}

// BadRequestException 400错误
func BadRequestException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	JSON(ctx, http.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	JSON(ctx, http.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[Failed]
	}
	JSON(ctx, http.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	JSON(ctx, http.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	JSON(ctx, http.StatusInternalServerError, InternalErr, message, nil)
}

// CustomError 自定义错误返回
func CustomError(code int, message string) gin.H {
	return gin.H{"errcode": code, "message": message}
}
