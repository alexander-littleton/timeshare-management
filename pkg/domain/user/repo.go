package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DbConn interface {
}

type Repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) Repo {
	return Repo{
		db: db,
	}
}

func (r Repo) CreateUser(ctx context.Context, user User) error {
	query := `INSERT INTO users (name, email) VALUES (@userName, @userEmail)`
	args := pgx.NamedArgs{
		"email":         user.Email,
		"password_hash": user.PasswordHash,
	}
	_, err := r.db.Exec(ctx, query, args)
	return err
}
