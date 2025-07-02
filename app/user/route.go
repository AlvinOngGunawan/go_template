package user

import (
	"Test_Go/config"
	"Test_Go/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(srv *echo.Echo, handler UserHandler, config config.AppConfig) {
	srv.POST("/register", handler.Register)
	srv.POST("/login", handler.Login)
	userGroup := srv.Group("/api")
	userGroup.Use(middleware.JWTMiddleware(config.JWTSecret))
	userGroup.GET("/profile", handler.Profile)
}
