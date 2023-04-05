package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo/pkg/model"
)

type AuthMy struct {
	db *sqlx.DB
}

func NewAuthMy(db *sqlx.DB) *AuthMy {
	return &AuthMy{db: db}
}

func (r *AuthMy) CreateUser(user model.User) (int, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values(?, ?, ?);", usersTable)
	res, err := r.db.Exec(query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *AuthMy) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=? AND password_hash=?", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
