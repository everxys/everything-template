package middleware

import (
	"strings"
	"time"

	"everything-template/internal/vars"
	"everything-template/pkg/logger"
	"everything-template/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			var err error
			tokenString, err = c.Cookie("jwt")
			if err != nil {
				response.UnauthorizedException(c, "Unauthorized: No valid authentication token")
				c.Abort()
				return
			}
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(vars.Config.Auth.SecretKey), nil
		})
		if err != nil {
			logger.Errorw("Token parsing failed", "error", err)
			response.UnauthorizedException(c, "Unauthorized: "+err.Error())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
				response.UnauthorizedException(c, "Unauthorized: Token expired")
				c.Abort()
				return
			}

			if sub, ok := claims["sub"].(string); ok {
				c.Set("user_id", sub)
			}
			if name, ok := claims["name"].(string); ok {
				c.Set("user_name", name)
			}
			if email, ok := claims["email"].(string); ok {
				c.Set("user_email", email)
			}

			c.Next()
		} else {
			response.UnauthorizedException(c, "Unauthorized: Invalid token")
			c.Abort()
		}
	}
}
