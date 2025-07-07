package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelInventoryReturn(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Kirim Balik PGI %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	headers := constanta.HeaderInventoryReturnReport

	query, args, err := sqlx.In(constanta.QueryExportInventoryReturnsReport, startDate, endDate, userID)
	if err != nil {
		return "", err
	}
	errs := s.generateExcelInventoryReturnFromDB(headers, query, args, "Laporan Kirim Balik PGI")
	if errs != nil {
		return "", errs
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelInventoryReturnFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return err
		}
		s.f.DeleteSheet("Sheet1")
	}

	streamWriter, err := s.f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}
	// write header
	if err = streamWriter.SetRow("A1", headers); err != nil {
		return err
	}

	rows, err := s.repo.DB.Queryx(query, param...)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows == nil {
		return errors.New("Data Not Found")
	}

	rowIndex := 2

	for rows.Next() {
		var i InventoryReturnsReport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFakturPGI.String, i.Status.String, i.CreatedAt.Time.Format("02-01-2006"), i.CreatedAt.Time.Format("15:04:05"),
			i.RequestNote.String, i.ProcessedAt.Time.Format("02-01-2006"), i.ProcessedAt.Time.Format("15:04:05"), i.ProcessedName.String,
			i.CancelAt.Time.Format("02-01-2006"), i.CancelNote.String, i.CancelName.String}
		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return err
		}
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return err
	}

	return err
}
