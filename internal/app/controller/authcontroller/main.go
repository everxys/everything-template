package authcontroller

import (
	"strconv"
	"time"

	"everything-template/internal/app/entity"
	"everything-template/internal/vars"
	"everything-template/pkg/logger"
	"everything-template/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	logger.Infow("Received a registration request")
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		response.BadRequestException(c, "Failed to parse request body")
		return
	}

	logger.Infow("User name: ", data["name"], "email", data["email"])

	var existingUser entity.User
	if err := vars.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		response.BadRequestException(c, "Email already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		response.InternalServerException(c, "Failed to hash password")
		return
	}

	user := &entity.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}
	if err := vars.DB.Create(&user).Error; err != nil {
		response.InternalServerException(c, "Failed to create user")
		return
	}

	logger.Infow("User registered successfully")
	response.SuccessJSON(c, "User registered successfully", nil)
}

func Login(c *gin.Context) {
	logger.Infow("Received a Login request")

	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		response.BadRequestException(c, "Failed to parse request body")
		return
	}

	var user entity.User
	vars.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		logger.Infow("User not found")
		response.UnauthorizedException(c, "Invalid credentials")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		logger.Infow("Invalid Password:", err)
		response.UnauthorizedException(c, "Invalid credentials")
		return
	}

	logger.Infow("Generating JWT token")

	expirationTime := time.Now().Add(vars.Config.Auth.TokenExpiry * time.Hour)
	claims := jwt.MapClaims{
		"sub":   strconv.Itoa(int(user.ID)),
		"exp":   expirationTime.Unix(),
		"name":  user.Name,
		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(vars.Config.Auth.SecretKey))
	if err != nil {
		logger.Errorw("Error generating token:", err)
		response.InternalServerException(c, "Failed to generate token")
		return
	}

	c.SetCookie("jwt", tokenString, vars.Config.Auth.CookieMaxAge, "/", "", true, true)

	logger.Infow("Authentication successful, returning")

	responseData := map[string]interface{}{
		"token":      tokenString,
		"expires_at": expirationTime.Unix(),
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	}

	response.SuccessJSON(c, "Login successful", responseData)
}
