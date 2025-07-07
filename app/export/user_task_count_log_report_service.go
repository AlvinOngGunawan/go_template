package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelUserTaskCountLogs(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Tugas Admin %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	headers := constanta.HeaderUserTaskCountLogsReport

	query, args, err := sqlx.In(constanta.QueryExportUserTaskCountLogReport, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	errs := s.generateExcelUserTaskCountLogFromDB(headers, query, args, "Laporan Tugas Admin")
	if errs != nil {
		return "", errs
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelUserTaskCountLogFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i UserTaskCountLogsReport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.LogDate.Time.Format("02-01-2006"), i.FullName.String, i.ApproveCount.Int64,
			i.InvoiceDetailCount.Int64, i.ReturnCount.Int64, i.ReqeustResetCount.Int64, i.AdjustmentCount.Int64, i.CetakBarcode.Int64,
			i.PindahGudang.Int64, i.UploadFoto.Int64, i.InputAksesoris.Int64}
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
