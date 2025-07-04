package server

import (
	"Test_Go/app/export"
	"Test_Go/app/user"
	"Test_Go/config"
	"github.com/bsm/redislock"
	"github.com/caarlos0/env/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

type ServerAttribute struct {
	Config    config.AppConfig
	Server    *echo.Echo
	DB        *sqlx.DB
	Redis     *redis.Client
	RedisLock *redislock.Client
	Module    Module
}

func (s *ServerAttribute) LoadConfig() {
	if err := env.Parse(&s.Config); err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
}

type Module struct {
	ExportModule *export.ExportModule
	UserModule   *user.UserModule
}
