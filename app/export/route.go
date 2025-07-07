package export

import (
	"Test_Go/config"
	"Test_Go/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(srv *echo.Echo, handler ExportHandler, config config.AppConfig) {
	userGroup := srv.Group("/export")
	userGroup.Use(middleware.JWTMiddleware(config.JWTSecret))
	userGroup.GET("/sales-report", handler.GenerateExcelSalesReport)
	userGroup.GET("/sales-report-gold", handler.GenerateExcelSalesReportGold)
	userGroup.GET("/return-invoice-report", handler.GenerateExcelReturnInvoice)
	userGroup.GET("/customer-transaction-report", handler.GenerateExcelCustomerTransaction)
	userGroup.GET("/delivery-report", handler.GenerateExcelDeliveryBatch)
	userGroup.GET("/delivery-report-gold", handler.GenerateExcelDeliveryGoldBatch)
	userGroup.GET("/inventory-report", handler.GenerateExcelInventoryMovement)
	userGroup.GET("/catalog-customer-login-log", handler.GenerateExcelCatalogCustomerLoginLog)
	userGroup.GET("/user-task-count-log", handler.GenerateExcelUserTaskCountLogs)
	userGroup.GET("/inventory-return-report", handler.GenerateExcelInventoryReturn)
	userGroup.GET("/sendback-report", handler.GenerateExcelSendbackReport)
}
