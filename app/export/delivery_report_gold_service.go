package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelDeliveryGoldReport(startDate, endDate string, userID int, typeReport string) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Penerimaan Barang Emas %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	headers := constanta.HeaderDeliveryBatchReport

	if typeReport == "EXT" {
		filename = fmt.Sprintf("Laporan Penerimaan Barang Emas External %s - %s - %s.xlsx", startDate, endDate, time.Now().String())
		query, args, err := sqlx.In(constanta.QueryExportDeliveryBatchGoldEXTReport, constanta.GoldDivision, userID)
		if err != nil {
			return "", err
		}
		listID, errs := s.generateExcelDeliveryBatchGoldFromDB(headers, query, args, "Batch Pengiriman")
		if errs != nil {
			return "", errs
		}

		//item
		headers = constanta.HeaderDeliveryItemGoldEXTReport
		query, args, err = sqlx.In(constanta.QueryExportDeliveryItemGoldEXTReport, constanta.GoldDivision, userID, startDate, endDate, listID)
		if err != nil {
			return "", err
		}
		errs = s.generateExcelDeliveryBatchItemGoldEXTFromDB(headers, query, args, "List Barang Pengiriman")
		if errs != nil {
			return "", errs
		}
	} else {
		query, args, err := sqlx.In(constanta.QueryExportDeliveryBatchGoldReport, constanta.GoldDivision, userID, startDate, endDate)
		if err != nil {
			return "", err
		}
		listID, errs := s.generateExcelDeliveryBatchGoldFromDB(headers, query, args, "Batch Pengiriman")
		if errs != nil {
			return "", errs
		}

		//item
		headers = constanta.HeaderDeliveryItemGoldReport
		query, args, err = sqlx.In(constanta.QueryExportDeliveryItemGoldReport, constanta.GoldDivision, userID, listID)
		if err != nil {
			return "", err
		}
		errs = s.generateExcelDeliveryBatchItemGoldFromDB(headers, query, args, "List Barang Pengiriman")
		if errs != nil {
			return "", errs
		}
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelDeliveryBatchGoldFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (listID []int64, err error) {
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

		result := []interface{}{i.Date.Time.Format("02-01-2006"), i.Branch.String, i.Source.String, i.ItemCount.Int64}
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

func (s *ExportService) generateExcelDeliveryBatchItemGoldEXTFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i DeliveryBatchItemGoldEXTReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.Date.Time.Format("02-01-2006"), i.Branch.String, i.Source.String,
			i.NoFakturPGI.String, i.KindName.String, i.BrandName.String, i.TypeName.String,
			i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.NetWeight.String,
			i.GoldMintMark.String, i.GoldType.String, i.PieceCount.Int64, s.resolveStatusDeliveryGold(i.Status.Int64),
			s.calculateBasePrice(i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.ItemKindID.Int64, i.ItemTypeID.Int64),
			i.WarehouseName.String, i.IsCap.String,
			i.FullName.String, i.ApprovedAt.Time.Format("02-01-2006"),
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

func (s *ExportService) generateExcelDeliveryBatchItemGoldFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i DeliveryBatchItemGoldReport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.Date.Time.Format("02-01-2006"), i.Branch.String, i.Source.String,
			i.NoFakturPGI.String, i.KindName.String, i.BrandName.String, i.TypeName.String,
			i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.NetWeight.String,
			i.GoldMintMark.String, i.GoldType.String, i.PieceCount.Int64, i.PawnedAt.Time.Format("02-01-2006"), s.resolveStatusDeliveryGold(i.Status.Int64),
			i.GradePGI.String, i.PriceAtPawn.Int64,
			s.calculateBasePrice(i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.ItemKindID.Int64, i.ItemTypeID.Int64),
			i.WarehouseName.String, i.IsCap.String,
			i.FullName.String, i.ApprovedAt.Time.Format("02-01-2006"),
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

func (s *ExportService) resolveStatusDeliveryGold(status int64) string {
	if status == 6 || status == 66 {
		return "Kirim Balik PGI"
	} else if status == 5 || status == 55 {
		return "Hilang"
	}
	return "Normal"
}
