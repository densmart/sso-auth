package entities

const (
	SuperUserRoleSlug = "superuser"
)

type Role struct {
	BaseEntity
	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	IsPermitted bool   `db:"is_permitted" json:"is_permitted"`
}
