package server

import (
	"Test_Go/config"
	"Test_Go/handler"
	"Test_Go/repository"
	"Test_Go/service/user"
	"database/sql"
	"github.com/caarlos0/env/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ServerAttribute struct {
	Config     config.AppConfig
	Server     *echo.Echo
	DB         *sql.DB
	Services   Services
	Repository Repository
	Handler    Handler
	Endpoint   Endpoint
}

func (s *ServerAttribute) LoadConfig() {
	if err := env.Parse(&s.Config); err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
}

type Services struct {
	userService user.UserService
}

type Repository struct {
	userRepository repository.UserRepository
}

type Handler struct {
	userHandler     handler.UserHandler
	categoryHandler handler.CategoryHandler
}

type Endpoint struct {
	jwtGroup *echo.Group
}
