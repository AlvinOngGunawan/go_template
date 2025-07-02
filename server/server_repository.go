package server

import "Test_Go/repository"

func (s *ServerAttribute) initRepository() {
	s.Repository = Repository{}
	s.Repository.userRepository = repository.NewUserRepository(s.DB)
}
