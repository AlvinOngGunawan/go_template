package export

import (
	"Test_Go/app/export/constanta"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"time"
)

func (s *ExportService) GenerateExcelCustomerTransaction(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()

	filename = fmt.Sprintf("Laporan Transaksi Customer %s - %s - %s.xlsx", startDate, endDate, time.Now().String())

	//customer transaction
	headers := constanta.HeaderCustomerTransactionReport

	query, args, err := sqlx.In(constanta.QueryExportCustomerTransactionReport, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	errs := s.generateExcelTransactionCustomerFromDB(headers, query, args, "Transaksi Customer")
	if errs != nil {
		return "", errs
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelTransactionCustomerFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
	typeOfTransaction := []string{"", "Pembelian", "Retur", "Pembayaran", "Pelunasan Hutang"}

	for rows.Next() {
		var i CustomerTransactionReport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		if i.Type.Int64 == 1 || i.Type.Int64 == 4 {
			i.DebtChange.Int64 *= -1
		} else {
			i.CreditChange.Int64 *= -1
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.CreatedAt.Time.Format("2006-01-02"), i.CustomerName.String, typeOfTransaction[i.Type.Int64], i.Amount.Int64,
			i.CreditChange.Int64, i.CreditAfter.Int64, i.DebtChange.Int64, i.DebtAfter.Int64, i.FullName.String, i.Params.String}
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
