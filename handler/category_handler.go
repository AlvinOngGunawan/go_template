package handler

import (
	"Test_Go/config"
	"Test_Go/service/user"
)

type CategoryHandler struct {
	svc    user.UserService
	config config.AppConfig
}

func NewCategoryHandler(config config.AppConfig, s user.UserService) CategoryHandler {
	return CategoryHandler{svc: s, config: config}
}
