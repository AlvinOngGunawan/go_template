package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelSalesReport(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Sales %s - %s - %s.xlsx", startDate, endDate, time.Now().String())
	var admin bool
	userOnDB, err2 := s.userRepo.GetUserByID(userID)
	if err2 != nil {
		return "", err2
	}
	if userOnDB.RoleID.Int64 == 1 {
		admin = true
	}

	//sales invoice
	headers := constanta.HeaderSalesReportSalesInvoice

	query, args, err := sqlx.In(constanta.QueryExportInvoice, constanta.GoldDivision, []interface{}{constanta.OnGoingStatus, constanta.PaidStatus}, userID, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	listID, errs := s.generateExcelSalesReportInvoiceFromDB(headers, query, args, "Sales Invoices", filename)
	if errs != nil {
		return "", errs
	}

	//sales invoice detail
	headers = constanta.HeaderSalesReportSalesinvoiceDetailNonAdmin
	if admin {
		headers = constanta.HeaderSalesReportSalesInvoiceDetailAdmin
	}

	query, args, err = sqlx.In(constanta.QueryExportInvoiceDetail, listID)
	if err != nil {
		return "", err
	}
	var listReturnID []int64
	listID, listReturnID, err = s.generateExcelSalesReportInvoiceDetailFromDB(headers, query, args, "Sales Items", admin)
	if err != nil {
		return "", err
	}

	//sales invoice detail adjustment
	headers = constanta.HeaderSalesReportInvoiceDetailAdjustment

	query, args, err = sqlx.In(constanta.QueryExportInvoiceDetailAdjustment, listID)
	if err != nil {
		return "", err
	}
	err = s.generateExcelSalesReportInvoiceDetailAdjustmentFromDB(headers, query, args, "Sales Items Adjustments")
	if err != nil {
		return "", err
	}

	//return invoice
	headers = constanta.HeaderSalesReportReturnInvoice

	query, args, err = sqlx.In(constanta.QueryExportInvoiceReturn, listReturnID)
	if err != nil {
		return "", err
	}
	err = s.generateExcelSalesReportInvoiceReturnFromDB(headers, query, args, "Retur Items")
	if err != nil {
		return "", err
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelSalesReportInvoiceFromDB(headers []interface{}, query string, param []interface{}, sheetName string, filename string) (listID []int64, err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		s.f.DeleteSheet("Sheet1")
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return nil, err
		}
	}

	streamWriter, err := s.f.NewStreamWriter(sheetName)
	if err != nil {
		return nil, err
	}
	// write header
	if err = streamWriter.SetRow("A1", headers); err != nil {
		return nil, err
	}

	rows, err := s.repo.DB.Queryx(query, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows == nil {
		return nil, errors.New("Data Not Found")
	}

	rowIndex := 2

	for rows.Next() {
		var i InvoicesExport
		if err = rows.StructScan(&i); err != nil {
			return nil, err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("2006-01-02"), i.Fullname.String, s.resolveStatus(i.Status.Int64),
			i.WarehouseName.String, i.OfficeName.String, s.resolvePriceSegmentName(i.PriceSegment.Int64), i.ItemCount.Int64, i.SubTotal.String, i.Discount.Float64,
			i.DiscountValue.Int64, i.Total.Int64, i.TotalReturn.Int64, i.OutstandingPayment.Int64, i.LastPayment.Time.Format("2006-01-02"), i.Note.String, i.DivisionName.String}
		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return nil, err
		}
		listID = append(listID, i.ID.Int64)
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return nil, err
	}

	return listID, err
}

func (s *ExportService) generateExcelSalesReportInvoiceDetailFromDB(headers []interface{}, query string, param []interface{}, sheetName string, admin bool) (listID []int64, listIDreturned []int64, err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		s.f.DeleteSheet("Sheet1")
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return
		}
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
		err = errors.New("Data Not Found")
		return
	}

	rowIndex := 2

	for rows.Next() {
		var i InvoiceDetailsExport
		if err = rows.StructScan(&i); err != nil {
			return
		}

		cell := fmt.Sprintf("A%d", rowIndex)
		var result []interface{}

		if admin {
			result = []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("2006-01-02"), i.NoFakturPGI.String, i.IMEISN.String,
				i.PawnedAt.String, i.CreatedAt.Time.Format("2006-01-02"), i.Source.String, i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
				i.TypeName.String, i.Year.String, i.SpecName.String, i.Batangan.String, i.Grade.String, i.GradePGI.String, i.PriceAtPawn.Int64, i.GradeAPrice.Int64, i.AdjGrade.Int64,
				i.AdjSpec.Int64, i.AdjBatangan.Int64, i.AdjOther.Int64, i.Adjustment.Int64, i.AdjPriceSeg.Int64, i.DiscountValue.Int64, i.Total.Int64, i.Capital.Int64, i.PL.Int64,
				i.Notes.String, i.AdjOtherNote.String}
		} else {
			result = []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("2006-01-02"), i.NoFakturPGI.String, i.IMEISN.String,
				i.PawnedAt.String, i.CreatedAt.Time.Format("2006-01-02"), i.Source.String, i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
				i.TypeName.String, i.Year.String, i.SpecName.String, i.Batangan.String, i.Grade.String, i.GradePGI.String, i.Adjustment.Int64, i.AdjPriceSeg.Int64, i.DiscountValue.Int64, i.Total.Int64,
				i.Notes.String}
		}

		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return
		}
		listID = append(listID, i.ID.Int64)
		if i.InvoiceReturnID.Int64 != 0 {
			listIDreturned = append(listIDreturned, i.InvoiceReturnID.Int64)
		}
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return
	}

	return
}

func (s *ExportService) generateExcelSalesReportInvoiceDetailAdjustmentFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		s.f.DeleteSheet("Sheet1")
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return err
		}
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
		var i InvoiceDetailAdjustmentExport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		cell := fmt.Sprintf("A%d", rowIndex)
		var result []interface{}

		result = []interface{}{i.NoFaktur.String, i.CustomerName.String, i.NoFakturPGI.String, i.Name.String, i.Adjustment.Int64}

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

func (s *ExportService) generateExcelSalesReportInvoiceReturnFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		s.f.DeleteSheet("Sheet1")
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return err
		}
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
		var i ReturnInvoiceExport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		cell := fmt.Sprintf("A%d", rowIndex)
		var result []interface{}

		result = []interface{}{i.No.String, i.NoFaktur.String, i.CustomerName.String, i.SoldAt.Time.Format("2006-01-02"), i.NoFakturPGI.String, i.NoRef.String, i.IMEISN.String,
			i.PawnedAt.String, i.CreatedAt.Time.Format("2006-01-02"), i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
			i.TypeName.String, i.Year.String, i.SpecName.String, i.Batangan.String, i.Grade.String, i.GradePGI.String, i.Total.Int64, i.CreatedAtReturn.Time.Format("2006-01-02"),
			i.FullName.String, i.ReturnReason.String}

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

func (s *ExportService) resolvePriceSegmentName(priceSegment int64) string {
	switch priceSegment {
	case 1:
		return "Grosir Offline"
	case 2:
		return "Grosir Online"
	case 3:
		return "Retail"
	default:
		return "-"
	}
}

func (s *ExportService) resolveStatus(status int64) string {
	switch status {
	case 0:
		return "Void"
	case 1:
		return "Draft"
	case 2:
		return "Lunas"
	case 3:
		return "Paid"
	default:
		return "-"
	}
}
