package server

import "Test_Go/handler"

func (s *ServerAttribute) initHandler() {
	s.Handler = Handler{}
	s.Handler.userHandler = handler.NewUserHandler(s.Config, s.Services.userService)
	s.Handler.categoryHandler = handler.NewCategoryHandler(s.Config, s.Services.userService)
	return
}
