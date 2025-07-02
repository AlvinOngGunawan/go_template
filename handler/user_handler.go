package handler

import (
	"Test_Go/config"
	"Test_Go/dto/in"
	error2 "Test_Go/error"
	"Test_Go/service/user"
	"Test_Go/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	svc    user.UserService
	config config.AppConfig
}

func NewUserHandler(config config.AppConfig, s user.UserService) UserHandler {
	return UserHandler{svc: s, config: config}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req in.User
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, error2.Error{Error: err}, err)
	}
	if err := h.svc.Register(req.Email, req.Password); err != nil {
		return error2.JSON(c, http.StatusInternalServerError, error2.Error{Error: err}, err)
	}
	return error2.JSON(c, http.StatusCreated, req, nil)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req in.User
	if err := c.Bind(&req); err != nil {
		return error2.JSON(c, http.StatusBadRequest, error2.Error{Error: err}, err)
	}
	user, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		return error2.JSON(c, http.StatusUnauthorized, error2.Error{Error: err}, err)
	}
	token, _ := utils.GenerateJWT(user.ID.Int64, h.config.JWTSecret)
	return error2.JSON(c, http.StatusOK, token, nil)
}

func (h *UserHandler) Profile(c echo.Context) error {
	userID := utils.GetUserIDFromToken(c) // implement helper if needed
	return error2.JSON(c, http.StatusOK, echo.Map{"user_id": userID}, nil)
}
