package entities

import (
	"time"
)

type User struct {
	BaseEntity
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Phone       *string   `db:"phone" json:"phone"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	Is2fa       bool      `db:"is_2fa" json:"is_2fa"`
	Token2fa    *string   `db:"token_2fa" json:"token_2fa"`
	LastLoginAt time.Time `db:"last_login_at" json:"last_login_at"`
	RoleID      uint64    `db:"role_id" json:"role_id"`
}
