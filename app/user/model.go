package user

import "database/sql"

type User struct {
	ID                 sql.NullInt64  `db:"id"`
	Username           sql.NullString `db:"username"`
	Password           sql.NullString `db:"password"`
	FullName           sql.NullString `db:"fullname"`
	IsActive           sql.NullInt64  `db:"is_active"`
	IsAfterHourAllowed sql.NullInt64  `db:"is_afterhour_allowed"`
	RoleID             sql.NullInt64  `db:"role_id"`
	CreatedBy          sql.NullInt64  `db:"created_by"`
	UpdatedBy          sql.NullInt64  `db:"updated_by"`
	CreatedAt          sql.NullTime   `db:"created_at"`
	UpdatedAt          sql.NullTime   `db:"updated_at"`
	LoginTime          sql.NullTime   `db:"login_time"`
	Tries              sql.NullInt64  `db:"tries"`
}
