package user

import (
	"Test_Go/config"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type UserModule struct {
	UserService    UserService
	UserHandler    UserHandler
	UserRepository UserRepository
}

func NewUserModule(db *sqlx.DB, config config.AppConfig, redis *redis.Client, endpoint *echo.Echo) *UserModule {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(config, userService)
	InitRoutes(endpoint, userHandler, config)

	return &UserModule{
		UserService:    userService,
		UserHandler:    userHandler,
		UserRepository: userRepository,
	}
}
