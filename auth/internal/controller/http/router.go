package http

import (
	"access-platform/internal/controller/http/auth"
	"access-platform/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(log *zap.Logger, service *service.AuthService) *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", auth.Register(log, service))
		authGroup.POST("/login", auth.Login(log, service))
	}

	return router
}
