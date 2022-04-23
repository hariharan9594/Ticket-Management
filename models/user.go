package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           uint         `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UserName     string       `json:"user_name"`
	Email        string       `json:"email"`
	IsAdmin      sql.NullBool `json:"is_admin"`
	Password     string       `json:"password,omitempty" gorm:"-"`
	HashPassword []byte       `json:"-"`
}

type Ticket struct {
	Desk        string `json:"Desk"`
	Ticket_Type string `json:"ticket_type"`
	Ticket_ID   int    `json:"ticket_id"`
	Subject     string `json:"subject"`
	U_id        int    `json:"u_id"`
}
