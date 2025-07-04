package export

import (
	"database/sql"
)

type InvoicesExport struct {
	ID                 sql.NullInt64   `db:"id"`
	NoFaktur           sql.NullString  `db:"no_faktur"`
	NoRef              sql.NullString  `db:"no_ref"`
	Status             sql.NullInt64   `db:"status"`
	SubTotal           sql.NullString  `db:"subtotal"`
	Discount           sql.NullFloat64 `db:"discount"`
	DiscountPurity     sql.NullInt64   `db:"discount_purity"`
	DiscountValue      sql.NullInt64   `db:"discount_value"`
	Total              sql.NullInt64   `db:"total"`
	LastPayment        sql.NullTime    `db:"last_payment"`
	OutstandingPayment sql.NullInt64   `db:"outstanding_payment"`
	FromWarehouseID    sql.NullInt64   `db:"from_warehouse_id"`
	PriceSegment       sql.NullInt64   `db:"price_segment"`
	CustomerID         sql.NullInt64   `db:"customer_id"`
	CustomerName       sql.NullString  `db:"customer_name"`
	CustomerPhone      sql.NullString  `db:"customer_phone"`
	CustomerAddress    sql.NullString  `db:"customer_address"`
	Note               sql.NullString  `db:"note"`
	ConfirmedAt        sql.NullTime    `db:"confirmed_at"`
	ConfirmedBy        sql.NullInt64   `db:"confirmed_by"`
	SoldAt             sql.NullTime    `db:"sold_at"`
	VoidAt             sql.NullTime    `db:"void_at"`
	VoidBy             sql.NullInt64   `db:"void_by"`
	VoidReason         sql.NullString  `db:"void_reason"`
	CreatedBy          sql.NullInt64   `db:"created_by"`
	UpdatedBy          sql.NullInt64   `db:"updated_by"`
	CreatedAt          sql.NullTime    `db:"created_at"`
	UpdatedAt          sql.NullTime    `db:"updated_at"`
	DivisionID         sql.NullInt64   `db:"division_id"`
	CurrentTotal       sql.NullInt64   `db:"current_total"`
	Prints             sql.NullInt64   `db:"prints"`
	InvoiceType        sql.NullString  `db:"invoice_type"`
	Fullname           sql.NullString  `db:"fullname"`
	WarehouseName      sql.NullString  `db:"warehouse_name"`
	OfficeName         sql.NullString  `db:"office_name"`
	DivisionName       sql.NullString  `db:"division_name"`
	ItemCount          sql.NullInt64   `db:"item_count"`
	TotalReturn        sql.NullInt64   `db:"total_return"`
}

type InvoiceDetailsExport struct {
	ID              sql.NullInt64  `db:"id"`
	NoFaktur        sql.NullString `db:"no_faktur"`
	NoRef           sql.NullString `db:"no_ref"`
	CustomerName    sql.NullString `db:"customer_name"`
	SoldAt          sql.NullTime   `db:"sold_at"`
	NoFakturPGI     sql.NullString `db:"no_faktur_pgi"`
	IMEISN          sql.NullString `db:"imei_sn"`
	PawnedAt        sql.NullString `db:"pawned_at"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	Source          sql.NullString `db:"source"`
	WarehouseName   sql.NullString `db:"warehouse_name"`
	OfficeName      sql.NullString `db:"office_name"`
	KindName        sql.NullString `db:"kind_name"`
	BrandName       sql.NullString `db:"brand_name"`
	TypeName        sql.NullString `db:"type_name"`
	Year            sql.NullString `db:"year"`
	SpecName        sql.NullString `db:"spec_name"`
	Batangan        sql.NullString `db:"batangan"`
	Grade           sql.NullString `db:"grade"`
	GradePGI        sql.NullString `db:"grade_pgi"`
	PriceAtPawn     sql.NullInt64  `db:"price_at_pawn"`
	GradeAPrice     sql.NullInt64  `db:"grade_a_price"`
	AdjGrade        sql.NullInt64  `db:"adj_grade"`
	AdjSpec         sql.NullInt64  `db:"adj_spec"`
	AdjBatangan     sql.NullInt64  `db:"adj_batangan"`
	AdjOther        sql.NullInt64  `db:"adj_other"`
	Adjustment      sql.NullInt64  `db:"adjustment"`
	AdjPriceSeg     sql.NullInt64  `db:"adjustment_price_segment"`
	DiscountValue   sql.NullInt64  `db:"discount_value"`
	Total           sql.NullInt64  `db:"total"`
	Capital         sql.NullInt64  `db:"capital"`
	PL              sql.NullInt64  `db:"pl"`
	Notes           sql.NullString `db:"notes"`
	AdjOtherNote    sql.NullString `db:"adj_other_note"`
	InvoiceReturnID sql.NullInt64  `db:"invoice_return_id"`
}

type InvoiceReturnsExport struct {
	ID        sql.NullInt64  `db:"id"`
	InvoiceID sql.NullInt64  `db:"invoice_id"`
	No        sql.NullString `db:"no"`
	Refund    sql.NullInt64  `db:"refund"`
	Amount    sql.NullInt64  `db:"amount"`
	CreatedBy sql.NullInt64  `db:"created_by"`
	UpdatedBy sql.NullInt64  `db:"updated_by"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	Prints    sql.NullInt64  `db:"prints"`
}

type InvoiceDetailAdjustmentExport struct {
	NoFaktur     sql.NullString `db:"no_faktur"`
	CustomerName sql.NullString `db:"customer_name"`
	NoFakturPGI  sql.NullString `db:"no_faktur_pgi"`
	Name         sql.NullString `db:"name"`
	Adjustment   sql.NullInt64  `db:"adjustment"`
}

type ReturnInvoiceExport struct {
	No              sql.NullString `db:"no"`
	NoFaktur        sql.NullString `db:"no_faktur"`
	CustomerName    sql.NullString `db:"customer_name"`
	SoldAt          sql.NullTime   `db:"sold_at"`
	NoFakturPGI     sql.NullString `db:"no_faktur_pgi"`
	NoRef           sql.NullString `db:"no_ref"`
	IMEISN          sql.NullString `db:"imei_sn"`
	PawnedAt        sql.NullString `db:"pawned_at"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	WarehouseName   sql.NullString `db:"warehouse_name"`
	OfficeName      sql.NullString `db:"office_name"`
	KindName        sql.NullString `db:"kind_name"`
	BrandName       sql.NullString `db:"brand_name"`
	TypeName        sql.NullString `db:"type_name"`
	Year            sql.NullString `db:"year"`
	SpecName        sql.NullString `db:"spec_name"`
	Batangan        sql.NullString `db:"batangan"`
	Grade           sql.NullString `db:"grade"`
	GradePGI        sql.NullString `db:"grade_pgi"`
	Total           sql.NullInt64  `db:"total"`
	CreatedAtReturn sql.NullTime   `db:"created_at_return"`
	FullName        sql.NullString `db:"fullname"`
	ReturnReason    sql.NullString `db:"return_reason"`
}

type InvoicesGoldExport struct {
	ID                 sql.NullInt64   `db:"id"`
	NoFaktur           sql.NullString  `db:"no_faktur"`
	NoRef              sql.NullString  `db:"no_ref"`
	Status             sql.NullInt64   `db:"status"`
	SubTotal           sql.NullString  `db:"subtotal"`
	Discount           sql.NullFloat64 `db:"discount"`
	DiscountPurity     sql.NullInt64   `db:"discount_purity"`
	DiscountValue      sql.NullInt64   `db:"discount_value"`
	Total              sql.NullInt64   `db:"total"`
	LastPayment        sql.NullTime    `db:"last_payment"`
	OutstandingPayment sql.NullInt64   `db:"outstanding_payment"`
	FromWarehouseID    sql.NullInt64   `db:"from_warehouse_id"`
	PriceSegment       sql.NullInt64   `db:"price_segment"`
	CustomerID         sql.NullInt64   `db:"customer_id"`
	CustomerName       sql.NullString  `db:"customer_name"`
	CustomerPhone      sql.NullString  `db:"customer_phone"`
	CustomerAddress    sql.NullString  `db:"customer_address"`
	Note               sql.NullString  `db:"note"`
	ConfirmedAt        sql.NullTime    `db:"confirmed_at"`
	ConfirmedBy        sql.NullInt64   `db:"confirmed_by"`
	SoldAt             sql.NullTime    `db:"sold_at"`
	VoidAt             sql.NullTime    `db:"void_at"`
	VoidBy             sql.NullInt64   `db:"void_by"`
	VoidReason         sql.NullString  `db:"void_reason"`
	CreatedBy          sql.NullInt64   `db:"created_by"`
	UpdatedBy          sql.NullInt64   `db:"updated_by"`
	CreatedAt          sql.NullTime    `db:"created_at"`
	UpdatedAt          sql.NullTime    `db:"updated_at"`
	DivisionID         sql.NullInt64   `db:"division_id"`
	CurrentTotal       sql.NullInt64   `db:"current_total"`
	Prints             sql.NullInt64   `db:"prints"`
	InvoiceType        sql.NullString  `db:"invoice_type"`
	Fullname           sql.NullString  `db:"fullname"`
	WarehouseName      sql.NullString  `db:"warehouse_name"`
	OfficeName         sql.NullString  `db:"office_name"`
	DivisionName       sql.NullString  `db:"division_name"`
	ItemCount          sql.NullInt64   `db:"item_count"`
	TotalReturn        sql.NullInt64   `db:"total_return"`
}

type InvoiceDetailsGoldExport struct {
	ID              sql.NullInt64   `db:"id"`
	NoFaktur        sql.NullString  `db:"no_faktur"`
	NoRef           sql.NullString  `db:"no_ref"`
	CustomerName    sql.NullString  `db:"customer_name"`
	SoldAt          sql.NullTime    `db:"sold_at"`
	NoFakturPGI     sql.NullString  `db:"no_faktur_pgi"`
	IMEISN          sql.NullString  `db:"imei_sn"`
	PawnedAt        sql.NullString  `db:"pawned_at"`
	CreatedAt       sql.NullTime    `db:"created_at"`
	Source          sql.NullString  `db:"source"`
	WarehouseName   sql.NullString  `db:"warehouse_name"`
	OfficeName      sql.NullString  `db:"office_name"`
	KindName        sql.NullString  `db:"kind_name"`
	BrandName       sql.NullString  `db:"brand_name"`
	TypeName        sql.NullString  `db:"type_name"`
	Purity          sql.NullInt64   `db:"purity"`
	DryWeight       sql.NullFloat64 `db:"dry_weight"`
	WeightReduction sql.NullFloat64 `db:"weight_reduction"`
	NetWeight       sql.NullString  `db:"net_weight"`
	GoldMintMark    sql.NullString  `db:"gold_mint_mark_name"`
	GoldType        sql.NullString  `db:"name"`
	PieceCount      sql.NullInt64   `db:"piece_count"`
	DiscountValue   sql.NullInt64   `db:"discount_value"`
	Total           sql.NullInt64   `db:"total"`
	Capital         sql.NullInt64   `db:"capital"`
	PL              sql.NullInt64   `db:"pl"`
	InvoiceReturnID sql.NullInt64   `db:"invoice_return_id"`
	ItemKindID      sql.NullInt64   `db:"item_kind_id"`
	TypeID          sql.NullInt64   `db:"type_id"`
}

type ReturnInvoiceGoldExport struct {
	No              sql.NullString  `db:"no"`
	NoFaktur        sql.NullString  `db:"no_faktur"`
	CustomerName    sql.NullString  `db:"customer_name"`
	SoldAt          sql.NullTime    `db:"sold_at"`
	NoFakturPGI     sql.NullString  `db:"no_faktur_pgi"`
	NoRef           sql.NullString  `db:"no_ref"`
	PawnedAt        sql.NullString  `db:"pawned_at"`
	CreatedAt       sql.NullTime    `db:"created_at"`
	WarehouseName   sql.NullString  `db:"warehouse_name"`
	OfficeName      sql.NullString  `db:"office_name"`
	KindName        sql.NullString  `db:"kind_name"`
	BrandName       sql.NullString  `db:"brand_name"`
	TypeName        sql.NullString  `db:"type_name"`
	Purity          sql.NullInt64   `db:"purity"`
	DryWeight       sql.NullFloat64 `db:"dry_weight"`
	WeightReduction sql.NullFloat64 `db:"weight_reduction"`
	NetWeight       sql.NullString  `db:"net_weight"`
	GoldMintMark    sql.NullString  `db:"gold_min_mark_name"`
	GoldType        sql.NullString  `db:"type_name"`
	PieceCount      sql.NullInt64   `db:"piece_count"`
	Total           sql.NullInt64   `db:"total"`
	CreatedAtReturn sql.NullTime    `db:"created_at_return"`
	FullName        sql.NullString  `db:"fullname"`
	ReturnReason    sql.NullString  `db:"return_reason"`
}

type ReturnInvoiceReport struct {
	ID           sql.NullInt64  `db:"id"`
	No           sql.NullString `db:"no"`
	NoFaktur     sql.NullString `db:"no_faktur"`
	NoRef        sql.NullString `db:"no_ref"`
	CustomerName sql.NullString `db:"customer_name"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	FullName     sql.NullString `db:"fullname"`
	Amount       sql.NullInt64  `db:"amount"`
	Refund       sql.NullInt64  `db:"refund"`
}

type CustomerTransactionReport struct {
	CreatedAt    sql.NullTime   `db:"created_at"`
	CustomerName sql.NullString `db:"customer_name"`
	Type         sql.NullInt64  `db:"type"`
	Amount       sql.NullInt64  `db:"amount"`
	CreditChange sql.NullInt64  `db:"credit_change"`
	CreditAfter  sql.NullInt64  `db:"credit_after"`
	DebtChange   sql.NullInt64  `db:"debt_change"`
	DebtAfter    sql.NullInt64  `db:"debt_after"`
	FullName     sql.NullString `db:"fullname"`
	Params       sql.NullString `db:"params"`
}

type DeliveryBatchReport struct {
	ID        sql.NullInt64  `db:"id"`
	Date      sql.NullTime   `db:"date"`
	Branch    sql.NullString `db:"branch"`
	Source    sql.NullString `db:"source"`
	ItemCount sql.NullInt64  `db:"item_count"`
}

type DeliveryBatchItemReport struct {
	Date               sql.NullTime   `db:"date"`
	Branch             sql.NullString `db:"branch"`
	Source             sql.NullString `db:"source"`
	NoFakturPGI        sql.NullString `db:"no_faktur_pgi"`
	IMEISN             sql.NullString `db:"imei_sn"`
	KindName           sql.NullString `db:"kind_name"`
	BrandName          sql.NullString `db:"brand_name"`
	TypeName           sql.NullString `db:"type_name"`
	Year               sql.NullString `db:"year"`
	PawnedAt           sql.NullString `db:"pawned_at"`
	Status             sql.NullInt64  `db:"status"`
	GradePGI           sql.NullString `db:"grade_pgi"`
	BatanganPGI        sql.NullString `db:"batangan_pgi"`
	PriceAtPawn        sql.NullInt64  `db:"price_at_pawn"`
	BasePrice          sql.NullInt64  `db:"base_price"`
	FinalPrice         sql.NullInt64  `db:"final_price"`
	Capital            sql.NullInt64  `db:"capital"`
	Grade              sql.NullString `db:"grade"`
	SpecName           sql.NullString `db:"spec_name"`
	Batangan           sql.NullString `db:"batangan"`
	WarehouseName      sql.NullString `db:"warehouse_name"`
	MissingAccessories sql.NullString `db:"missing_accessories"`
	NotOriAccessories  sql.NullString `db:"not_ori_accessories"`
	FullName           sql.NullString `db:"fullname"`
	ApprovedAt         sql.NullTime   `db:"approved_at"`
	Description        sql.NullString `db:"description"`
}
