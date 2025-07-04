package user

import "github.com/jmoiron/sqlx"

type UserRepository struct{ DB *sqlx.DB }

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r UserRepository) GetUserByID(id int) (user User, err error) {
	err = r.DB.Get(&user, "SELECT id, username, password, fullname, is_active, role_id FROM users WHERE id = ?", id)
	return
}
