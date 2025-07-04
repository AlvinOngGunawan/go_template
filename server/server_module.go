package server

import (
	"Test_Go/app/export"
	"Test_Go/app/user"
)

func (s *ServerAttribute) initModule() {
	s.Module = Module{}

	s.Module.UserModule = user.NewUserModule(s.DB, s.Config, s.Redis, s.Server)
	s.Module.ExportModule = export.NewExportModule(s.DB, s.Config, s.Redis, s.Server, s.Module.UserModule.UserRepository)
	return
}
