package server

import (
	"Test_Go/middleware"
)

func (s *ServerAttribute) initEndpoint() {
	s.Endpoint = Endpoint{}

	//API Group
	s.Endpoint.jwtGroup = s.Server.Group("/api")
	s.Endpoint.jwtGroup.Use(middleware.JWTMiddleware(s.Config.JWTSecret))
}
