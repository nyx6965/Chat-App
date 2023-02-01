// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  username, email, password
) VALUES (
   ?, ?, ?
)
`

type CreateUserParams struct {
	Username sql.NullString `json:"username"`
	Email    sql.NullString `json:"email"`
	Password sql.NullString `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
}

const loginUser = `-- name: LoginUser :one
SELECT id, email, username, password, created_at FROM users
WHERE username = ?
`

func (q *Queries) LoginUser(ctx context.Context, username sql.NullString) (Users, error) {
	row := q.db.QueryRowContext(ctx, loginUser, username)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
