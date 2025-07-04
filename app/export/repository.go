package export

import (
	"github.com/jmoiron/sqlx"
)

type ExportRepository struct{ DB *sqlx.DB }

func NewExportRepository(db *sqlx.DB) ExportRepository {
	return ExportRepository{DB: db}
}

func (d *ExportRepository) GetGoldPriceByKindID(kindID int64) (result float64, err error) {
	err = d.DB.Get(&result, "SELECT price FROM gold_prices WHERE item_kind_id = ? LIMIT 1", kindID)
	return
}

func (d *ExportRepository) GetGoldPriceByTypeID(typeID int64) (result float64, err error) {
	err = d.DB.Get(&result, "SELECT price FROM gold_prices WHERE item_type_id = ? LIMIT 1", typeID)
	return
}
