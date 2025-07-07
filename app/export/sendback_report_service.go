package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelSendback(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Kirim Inventaris %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	headers := constanta.HeaderSendbackReport

	query, args, err := sqlx.In(constanta.QueryExportSendbackReport, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	listID, errs := s.generateExcelSendbackFromDB(headers, query, args, "Sendback Invoice")
	if errs != nil {
		return "", errs
	}

	headers = constanta.HeaderSendbackDetailReport
	query, args, err = sqlx.In(constanta.QueryExportSendbackDetailReport, listID)
	if err != nil {
		return "", err
	}
	errs = s.generateExcelSendbackDetailsFromDB(headers, query, args, "Sendback Items")
	if errs != nil {
		return "", errs
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelSendbackFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (listID []int64, err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return
		}
		s.f.DeleteSheet("Sheet1")
	}

	streamWriter, err := s.f.NewStreamWriter(sheetName)
	if err != nil {
		return
	}
	// write header
	if err = streamWriter.SetRow("A1", headers); err != nil {
		return
	}

	rows, err := s.repo.DB.Queryx(query, param...)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows == nil {
		return nil, errors.New("Data Not Found")
	}

	rowIndex := 2

	for rows.Next() {
		var i SendbackReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.SendDate.Time, i.ConfirmedAt.Time, i.Fullname.String, i.Status.String,
			i.Warehouse.String, i.DetailsCount.Int64, i.Notes.String}
		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return
		}
		listID = append(listID, i.ID.Int64)
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return
	}

	return
}

func (s *ExportService) generateExcelSendbackDetailsFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return
		}
		s.f.DeleteSheet("Sheet1")
	}

	streamWriter, err := s.f.NewStreamWriter(sheetName)
	if err != nil {
		return
	}
	// write header
	if err = streamWriter.SetRow("A1", headers); err != nil {
		return
	}

	rows, err := s.repo.DB.Queryx(query, param...)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows == nil {
		return errors.New("Data Not Found")
	}

	rowIndex := 2

	for rows.Next() {
		var i SendbackDetailReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.NoFakturPGI.String, i.CreatedAt.Time.Format("2006-02-01"),
			i.KindName.String, i.IMEISN.String, i.BrandName.String, i.TypeName.String, i.SpecName.String, i.Year.String,
			i.Batangan.String, i.Grade.String, i.GradePGI.String, i.FinalPriceAfterAdj.Int64, i.WarehouseName.String, i.AdjName.String}
		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return
		}
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return
	}

	return
}
