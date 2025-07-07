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

func (s *ExportService) GenerateExcelSalesReportGold(startDate, endDate string, userID int) (filename string, err error) {
	s.f = excelize.NewFile()
	defer s.f.Close()
	var admin bool

	filename = fmt.Sprintf("Laporan Sales Emas %s - %s - %s.xlsx", startDate, endDate, time.Now().String())
	userOnDB, err2 := s.userRepo.GetUserByID(userID)
	if err2 != nil {
		return "", err2
	}
	if userOnDB.RoleID.Int64 == 1 {
		admin = true
	}

	//sales invoice gold
	headers := constanta.HeaderSalesReportInvoiceGold

	query, args, err := sqlx.In(constanta.QueryExportInvoiceGold, constanta.GoldDivision, []interface{}{constanta.OnGoingStatus, constanta.PaidStatus}, userID, userID, startDate, endDate)
	if err != nil {
		return "", err
	}
	listID, errs := s.generateExcelSalesReportInvoiceGoldFromDB(headers, query, args, "Sales Invoices", filename)
	if errs != nil {
		return "", errs
	}

	//sales invoice detail gold
	headers = constanta.HeaderSalesReportInvoiceGoldDetailNonAdmin
	if admin {
		headers = constanta.HeaderSalesReportInvoiceGoldDetail
	}
	query, args, err = sqlx.In(constanta.QueryExportInvoiceGoldDetail, listID)
	if err != nil {
		return "", err
	}
	listID, errs = s.generateExcelSalesReportInvoiceDetailGoldFromDB(headers, query, args, "Sales Items", admin)
	if errs != nil {
		return "", errs
	}

	//sales invoice return gold
	headers = constanta.HeaderSalesReportReturnInvoiceGold
	query, args, err = sqlx.In(constanta.QueryExportReturnInvoiceGold, listID)
	if err != nil {
		return "", err
	}
	err = s.generateExcelSalesReportInvoiceReturnGoldFromDB(headers, query, args, "Retur Items")
	if err != nil {
		return "", err
	}

	if err = s.f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, err
}

func (s *ExportService) generateExcelSalesReportInvoiceGoldFromDB(headers []interface{}, query string, param []interface{}, sheetName string, filename string) (listID []int64, err error) {
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
		var i InvoicesGoldExport
		if err = rows.StructScan(&i); err != nil {
			return nil, err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		result := []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("02-01-2006"), i.Fullname.String, s.resolveStatus(i.Status.Int64),
			i.WarehouseName.String, i.OfficeName.String, s.resolvePriceSegmentName(i.PriceSegment.Int64), i.ItemCount.Int64, i.SubTotal.String, i.Discount.Float64,
			i.DiscountValue.Int64, i.Total.Int64, i.TotalReturn.Int64, i.OutstandingPayment.Int64, i.LastPayment.Time.Format("02-01-2006"), i.Note.String, i.DivisionName.String,
			s.resolveInvoiceType(i.InvoiceType.String)}
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

func (s *ExportService) generateExcelSalesReportInvoiceDetailGoldFromDB(headers []interface{}, query string, param []interface{}, sheetName string, admin bool) (listIDReturned []int64, err error) {
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
		var i InvoiceDetailsGoldExport
		if err = rows.StructScan(&i); err != nil {
			return nil, err
		}

		cell := fmt.Sprintf("A%d", rowIndex)

		var result []interface{}

		if admin {
			result = []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("02-01-2006"), i.NoFakturPGI.String, i.IMEISN.String,
				i.PawnedAt.String, i.CreatedAt.Time.Format("02-01-2006"), i.Source.String, i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
				i.TypeName.String, i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.NetWeight.String, i.GoldMintMark.String, i.GoldType.String, i.PieceCount.Int64,
				s.calculateBasePrice(i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.ItemKindID.Int64, i.TypeID.Int64),
				i.DiscountValue.Int64, i.Total.Int64, i.Capital.Int64, i.PL.Int64}
		} else {
			result = []interface{}{i.NoFaktur.String, i.NoRef.String, i.CustomerName.String, i.SoldAt.Time.Format("02-01-2006"), i.NoFakturPGI.String, i.IMEISN.String,
				i.PawnedAt.String, i.CreatedAt.Time.Format("02-01-2006"), i.Source.String, i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
				i.TypeName.String, i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.NetWeight.String, i.GoldMintMark.String, i.GoldType.String, i.PieceCount.Int64,
				s.calculateBasePrice(i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.ItemKindID.Int64, i.TypeID.Int64),
				i.DiscountValue.Int64, i.Total.Int64}
		}

		err = streamWriter.SetRow(cell, result)
		if err != nil {
			return nil, err
		}
		listIDReturned = append(listIDReturned, i.InvoiceReturnID.Int64)
		rowIndex++
	}

	if err = streamWriter.Flush(); err != nil {
		return nil, err
	}

	return listIDReturned, err
}

func (s *ExportService) generateExcelSalesReportInvoiceReturnGoldFromDB(headers []interface{}, query string, param []interface{}, sheetName string) (err error) {
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
		var i ReturnInvoiceGoldExport
		if err = rows.StructScan(&i); err != nil {
			return err
		}

		cell := fmt.Sprintf("A%d", rowIndex)
		var result []interface{}

		result = []interface{}{i.No.String, i.NoFaktur.String, i.CustomerName.String, i.SoldAt.Time.Format("02-01-2006"), i.NoFakturPGI.String, i.NoRef.String,
			i.PawnedAt.String, i.CreatedAt.Time.Format("02-01-2006"), i.WarehouseName.String, i.OfficeName.String, i.KindName.String, i.BrandName.String,
			i.TypeName.String, i.Purity.Int64, i.DryWeight.Float64, i.WeightReduction.Float64, i.NetWeight.String, i.GoldMintMark.String, i.GoldType.String, i.PieceCount.Int64,
			i.Total.Int64, i.CreatedAtReturn.Time.Format("02-01-2006"),
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

func (s *ExportService) resolveInvoiceType(invoiceType string) string {
	switch invoiceType {
	case "AUCTION":
		return "Lelang"
	case "ORIGINAL_PRICE":
		return "Harga Asli"
	case "MELTED_SALE":
		return "Jual Lebur"
	default:
		return ""
	}
}

func (s *ExportService) calculateBasePrice(purity int64, dryWeight float64, weightReduction float64, itemKindID int64, itemTypeID int64) int64 {
	goldPrice, err := s.repo.GetGoldPriceByKindID(itemKindID)
	if err != nil {
		err = errors.New("Gold Price Not Found")
		return 0
	}
	if !((purity == 0) && (dryWeight == 0) && (weightReduction == 0)) {
		price := utils.CalculateGoldPrice(purity, dryWeight, weightReduction, goldPrice)
		return int64(price)
	}

	price, err := s.repo.GetGoldPriceByTypeID(itemTypeID)
	if err != nil {
		err = errors.New("Gold Price Not Found")
		return 0
	}
	return int64(price)
}
