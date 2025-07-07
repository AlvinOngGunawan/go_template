package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelInventoryMovementReport(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Perpindahan Barang %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	headers := constanta.HeaderInventoryMovementReport

	query, args, err := sqlx.In(constanta.QueryExportInventoryMovementReportExport, userID, startDate, endDate, userID, userID)
	if err != nil {
		return "", err
	}
	listID, errs := s.generateExcelInventoryMovementFromDB(headers, query, args, "Movement Invoices")
	if errs != nil {
		return "", errs
	}

	headers = constanta.HeaderInventoryMovementItemReport
	query, args, err = sqlx.In(constanta.QueryInventoryMovementItemReportExport, listID)
	if err != nil {
		return "", err
	}
	errs = s.generateExcelInventoryMovementItemFromDB(headers, query, args, "Movement Items")
	if errs != nil {
		return "", errs
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelInventoryMovementFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (listID []int64, err error) {
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
		var i InventoryMovementReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.NoRef.String, i.MoveDate.Time.Format("02-01-2006"),
			i.ConfirmedAt.Time.Format("02-01-2006"), i.FullName.String, i.Status.String, i.FromName.String,
			i.FromOfficeName.String, i.ToName.String, i.ToOfficeName.String, i.DetailsCount.Int64, i.Note.String}
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

func (s *ExportService) generateExcelInventoryMovementItemFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i InventoryMovementItemReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.NoRef.String, i.NoFakturPGI.String, i.IMEISN.String,
			i.PawnedAt.String, i.CreatedAt.Time.Format("02-01-2006"), i.WarehouseFromName.String,
			i.OfficeFromName.String, i.WarehouseToName.String, i.OfficeToName.String, i.KindName.String, i.BrandName.String, i.TypeName.String,
			i.Year.String, i.SpecName.String, i.Batangan.String, i.Grade.String, i.GradePGI.String, i.PriceAtPawn.Int64, i.Capital.Int64}
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
