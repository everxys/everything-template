package userController

import (
	"everything-template/internal/app/entity"
	"everything-template/internal/vars"
	"everything-template/pkg/logger"
	"everything-template/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func User(c *gin.Context) {
	logger.Infow("Request to get user...")

	authHeader := c.GetHeader("Authorization")
	var tokenString string

	if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		var err error
		tokenString, err = c.Cookie("jwt")
		if err != nil {
			response.UnauthorizedException(c, "Unauthorized: No valid authentication token")
			return
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(vars.Config.Auth.SecretKey), nil
	})

	if err != nil {
		response.UnauthorizedException(c, "Unauthorized: "+err.Error())
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
			response.UnauthorizedException(c, "Unauthorized: Token expired")
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			response.InternalServerException(c, "Invalid subject in token")
			return
		}

		id, err := strconv.Atoi(sub)
		if err != nil {
			response.InternalServerException(c, "Invalid user ID in token")
			return
		}

		var user entity.User
		if err := vars.DB.Where("id = ?", id).First(&user).Error; err != nil {
			response.InternalServerException(c, "Failed to retrieve user")
			return
		}

		userResponse := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}

		response.SuccessJSON(c, "User data retrieved successfully", userResponse)
	} else {
		response.UnauthorizedException(c, "Unauthorized: Invalid token")
	}
}

func Logout(c *gin.Context) {
	logger.Infow("Received a logout request")

	c.SetCookie("jwt", "", -1, "/", "", true, true)

	response.SuccessJSON(c, "Logout successful", nil)
}
