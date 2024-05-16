package postgre

import (
	"database/sql"
)

type Postgre struct {
	// DB is the connection to the database
	DB *sql.DB
}

func NewPostgre() (*Postgre, error) {
	return &Postgre{}, nil
}
