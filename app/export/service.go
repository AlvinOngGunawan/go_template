package export

import (
	"Test_Go/app/user"
	"github.com/xuri/excelize/v2"
)

type ExportService struct {
	repo     ExportRepository
	f        *excelize.File
	userRepo user.UserRepository
}

func NewExportService(r ExportRepository, u user.UserRepository) ExportService {
	return ExportService{repo: r, userRepo: u}
}
