package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Username string
	Password string
	DBName   string
}

func NewMyDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@/%s", cfg.Username, cfg.Password, cfg.DBName))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
