package http

import (
	"access-platform/internal/controller/http/handlers"
	"access-platform/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(log *zap.Logger, service *service.Service) *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", handlers.Register(log, service))
		authGroup.POST("/login", handlers.Login(log, service))
		authGroup.POST("/add_computer", handlers.AddComputer(log, service))
	}

	mainGroup := router.Group("/api")
	{
		mainGroup.POST("/add_computer", handlers.AddComputer(log, service))
		mainGroup.GET("/get_computer", handlers.GetComputer(log, service))
		mainGroup.POST("/reserve_computer", handlers.ReserveComputer(log, service))
		mainGroup.POST("/relieve_computer", handlers.RelieveComputer(log, service))
		mainGroup.GET("/get_all", handlers.GetAllComputers(log, service))
	}
	return router
}
