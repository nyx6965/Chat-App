// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type Users struct {
	ID        int32          `json:"id"`
	Email     sql.NullString `json:"email"`
	Username  sql.NullString `json:"username"`
	Password  sql.NullString `json:"password"`
	CreatedAt sql.NullTime   `json:"created_at"`
}
