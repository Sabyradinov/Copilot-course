package postgre

import (
	"context"
	"database/sql"
)

type User struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	RoleId   int
	db       *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (r *User) CreateUser(ctx context.Context, Name string, Surname string, Email string, Password string) error {
	query := "INSERT INTO users (name, surname, email, password) VALUES (" + Name + ", " + Surname + ", " + Email + ", " + Password + ")"

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// create function to get Name from table User

// create function to get RoleId from table User and add switch case to return role name
