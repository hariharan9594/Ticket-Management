package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id             uint         `json:"id"`
	CreatedAt      time.Time    `json:"created_at"`
	UserName       string       `json:"user_name"`
	Email          string       `json:"email"`
	IsAdmin        sql.NullBool `json:"is_admin"`
	Password       string       `json:"password,omitempty" gorm:"-"`
	HashPassword   []byte       `json:"-"`
}