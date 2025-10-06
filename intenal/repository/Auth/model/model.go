package model

import (
	"database/sql"
	"time"
)

type Auth struct {
	Id        int64        `db: "id"`
	Info      *AuthInfo    `db""`
	Role      int64        `db:"role"`
	CreateAt  time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type AuthInfo struct {
	Name  string `db: "Name"`
	Email string `db: "email"`
}
