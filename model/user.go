package model

import (
	"database/sql"
)

type User struct {
	ID       sql.NullInt64
	Email    sql.NullString
	Password sql.NullString
}
