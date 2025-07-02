package server

import (
	"Test_Go/service/user"
)

func (s *ServerAttribute) initServices() {
	s.Services = Services{}
	s.Services.userService = user.NewUserService(s.Repository.userRepository)
}
