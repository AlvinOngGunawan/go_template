package server

func (s *ServerAttribute) InitRoutes() {
	// Public
	s.Server.POST("/register", s.Handler.userHandler.Register)
	s.Server.POST("/login", s.Handler.userHandler.Login)

	// Protected group
	s.Endpoint.jwtGroup.GET("/profile", s.Handler.userHandler.Profile)
}
