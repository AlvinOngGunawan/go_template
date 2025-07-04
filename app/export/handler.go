package export

import (
	"Test_Go/config"
	error2 "Test_Go/error"
	"Test_Go/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ExportHandler struct {
	svc    ExportService
	config config.AppConfig
}

func NewExportHandler(config config.AppConfig, s ExportService) ExportHandler {
	return ExportHandler{svc: s, config: config}
}

func (h *ExportHandler) GenerateExcelSalesReport(c echo.Context) error {
	var req In
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, nil, err.Error())
	}
	userID := utils.GetUserIDFromToken(c)
	file, err := h.svc.GenerateExcelSalesReport(req.FromDate, req.ToDate, userID)
	if err != nil {
		return error2.JSON(c, http.StatusInternalServerError, nil, err.Error())
	}
	return error2.JSON(c, http.StatusOK, file, "")
}

func (h *ExportHandler) GenerateExcelSalesReportGold(c echo.Context) error {
	var req In
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, nil, err.Error())
	}
	userID := utils.GetUserIDFromToken(c)
	file, err := h.svc.GenerateExcelSalesReportGold(req.FromDate, req.ToDate, userID)
	if err != nil {
		return error2.JSON(c, http.StatusInternalServerError, nil, err.Error())
	}
	return error2.JSON(c, http.StatusOK, file, "")
}

func (h *ExportHandler) GenerateExcelReturnInvoice(c echo.Context) error {
	var req In
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, nil, err.Error())
	}
	userID := utils.GetUserIDFromToken(c)
	file, err := h.svc.GenerateExcelReturnInvoice(req.FromDate, req.ToDate, userID)
	if err != nil {
		return error2.JSON(c, http.StatusInternalServerError, nil, err.Error())
	}
	return error2.JSON(c, http.StatusOK, file, "")
}

func (h *ExportHandler) GenerateExcelCustomerTransaction(c echo.Context) error {
	var req In
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, nil, err.Error())
	}
	userID := utils.GetUserIDFromToken(c)
	file, err := h.svc.GenerateExcelCustomerTransaction(req.FromDate, req.ToDate, userID)
	if err != nil {
		return error2.JSON(c, http.StatusInternalServerError, nil, err.Error())
	}
	return error2.JSON(c, http.StatusOK, file, "")
}

func (h *ExportHandler) GenerateExcelDeliveryBatch(c echo.Context) error {
	var req In
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, nil, err.Error())
	}
	userID := utils.GetUserIDFromToken(c)
	file, err := h.svc.GenerateExcelDeliveryReport(req.FromDate, req.ToDate, userID, req.TypeReport)
	if err != nil {
		return error2.JSON(c, http.StatusInternalServerError, nil, err.Error())
	}
	return error2.JSON(c, http.StatusOK, file, "")
}
