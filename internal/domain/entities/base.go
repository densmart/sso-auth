package entities

import "time"

type BaseEntity struct {
	Id        uint64    `db:"id" json:"id"`                 // table general primary key bigint, autoincrement
	CreatedAt time.Time `db:"created_at" json:"created_at"` // record create datetime
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"` // record update datetime
}
