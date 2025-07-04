package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelReturnInvoice(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()
	var admin bool

	filename = fmt.Sprintf("Laporan Retur %s - %s - %s.xlsx", startDate, endDate, time.Now().String())
	userOnDB, err2 := s.userRepo.GetUserByID(userID)
	if err2 != nil {
		return "", err2
	}
	if userOnDB.RoleID.Int64 == 1 {
		admin = true
	}

	_ = admin

	//sales invoice gold
	headers := constanta.HeaderReturnInvoiceReport

	query, args, err := sqlx.In(constanta.QueryExportReturnInvoiceReport, userID, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	listReturnID, errs := s.generateExcelReturnInvoiceFromDB(headers, query, args, "Retur Invoices")
	if errs != nil {
		return "", errs
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

func (s *ExportService) generateExcelReturnInvoiceFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (listID []int64, err error) {
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		_, err = s.f.NewSheet(sheetName)
		if err != nil {
			return nil, err
		}
		s.f.DeleteSheet("Sheet1")
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
		var i ReturnInvoiceReport
		if err = rows.StructScan(&i); err != nil {
			return nil, err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.No.String, i.NoFaktur.String, i.NoRef.String, i.CustomerName.String,
			i.CreatedAt.Time.Format("2006-01-02"), i.FullName.String, i.Amount.Int64, i.Refund.Int64}
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
