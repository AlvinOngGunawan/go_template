package export

import (
	"Test_Go/app/export/constanta"
	"Test_Go/utils"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelDeliveryReport(startDate, endDate string, userID int, typeReport string) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Penerimaan Barang %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	//customer transaction
	headers := constanta.HeaderDeliveryBatchReport

	if typeReport == "EXT" {
		filename = fmt.Sprintf("Laporan Penerimaan Barang External %s - %s - %s.xlsx", startDate, endDate, time.Now().String())
		query, args, err := sqlx.In(constanta.QueryExportDeliveryBatchEXTReport, constanta.GoldDivision, userID, userID)
		if err != nil {
			return "", err
		}
		listID, errs := s.generateExcelDeliveryBatchFromDB(headers, query, args, "Batch Pengiriman")
		if errs != nil {
			return "", errs
		}

		//item
		headers = constanta.HeaderDeliveryItemEXTReport
		query, args, err = sqlx.In(constanta.QueryExportDeliveryItemEXTReport, constanta.GoldDivision, userID, startDate, endDate, listID)
		if err != nil {
			return "", err
		}
		errs = s.generateExcelDeliveryBatchItemEXTFromDB(headers, query, args, "List Barang Pengiriman")
		if errs != nil {
			return "", errs
		}
	} else {
		query, args, err := sqlx.In(constanta.QueryExportDeliveryBatchReport, constanta.GoldDivision, userID, userID, startDate, endDate)
		if err != nil {
			return "", err
		}
		listID, errs := s.generateExcelDeliveryBatchFromDB(headers, query, args, "Batch Pengiriman")
		if errs != nil {
			return "", errs
		}

		//item
		headers = constanta.HeaderDeliveryItemReport
		query, args, err = sqlx.In(constanta.QueryExportDeliveryItemReport, constanta.GoldDivision, userID, listID)
		if err != nil {
			return "", err
		}
		errs = s.generateExcelDeliveryBatchItemFromDB(headers, query, args, "List Barang Pengiriman")
		if errs != nil {
			return "", errs
		}
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelDeliveryBatchFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (listID []int64, err error) {
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
		var i DeliveryBatchReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.Date.Time.Format("2006-01-02"), i.Branch.String, i.Source.String, i.ItemCount.Int64}
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

func (s *ExportService) generateExcelDeliveryBatchItemEXTFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i DeliveryBatchItemReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.Date.Time.Format("2006-01-02"), i.Branch.String, i.Source.String,
			i.NoFakturPGI.String, i.IMEISN.String, i.KindName.String, i.BrandName.String, i.TypeName.String,
			i.Year.String, i.PawnedAt.String, s.resolveStatusDelivery(i.Status.Int64), i.GradePGI.String, i.BatanganPGI.String,
			i.PriceAtPawn.Int64, i.BasePrice.Int64, i.FinalPrice.Int64, i.Capital.Int64, i.Grade.String, i.SpecName.String,
			i.Batangan.String, i.WarehouseName.String,
			utils.DecodeAccessoriesArrayToString(i.MissingAccessories.String, "-"),
			utils.DecodeAccessoriesArrayToString(i.NotOriAccessories.String, "-"),
			i.FullName.String, i.ApprovedAt.Time.Format("2006-01-02"),
			i.ApprovedAt.Time.Format("15:04:05"), i.Description.String}
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

func (s *ExportService) generateExcelDeliveryBatchItemFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i DeliveryBatchItemReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.Date.Time.Format("2006-01-02"), i.Branch.String, i.Source.String,
			i.NoFakturPGI.String, i.IMEISN.String, i.KindName.String, i.BrandName.String, i.TypeName.String,
			i.Year.String, i.PawnedAt.String, s.resolveStatusDelivery(i.Status.Int64), i.GradePGI.String, i.BatanganPGI.String,
			i.PriceAtPawn.Int64, i.BasePrice.Int64, i.FinalPrice.Int64, i.Grade.String, i.SpecName.String,
			i.Batangan.String, i.WarehouseName.String,
			utils.DecodeAccessoriesArrayToString(i.MissingAccessories.String, "-"),
			utils.DecodeAccessoriesArrayToString(i.NotOriAccessories.String, "-"),
			i.FullName.String, i.ApprovedAt.Time.Format("2006-01-02"),
			i.ApprovedAt.Time.Format("15:04:05")}
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

func (s *ExportService) resolveStatusDelivery(status int64) string {
	if status == 6 || status == 66 {
		return "Kirim Balik PGI"
	} else if status == 5 || status == 55 {
		return "Hilang"
	}
	return "Normal"
}
