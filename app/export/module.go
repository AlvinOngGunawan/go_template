package export

import (
	"Test_Go/app/user"
	"Test_Go/config"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type ExportModule struct {
	ExportService    ExportService
	ExportHandler    ExportHandler
	ExportRepository ExportRepository
	UserRepository   user.UserRepository
}

func NewExportModule(db *sqlx.DB, config config.AppConfig, redis *redis.Client, endpoint *echo.Echo, userRepository user.UserRepository) *ExportModule {
	exportRepository := NewExportRepository(db)
	exportService := NewExportService(exportRepository, userRepository)
	exportHandler := NewExportHandler(config, exportService)
	InitRoutes(endpoint, exportHandler, config)

	return &ExportModule{
		ExportService:    exportService,
		ExportHandler:    exportHandler,
		ExportRepository: exportRepository,
		UserRepository:   userRepository,
	}
}
