package repo

import (
	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
)

type Roles interface {
	CreateRole(data dto.CreateRoleDTO) (entities.Role, error)
	RetrieveRole(id uint64) (entities.Role, error)
	UpdateRole(id uint64, data dto.UpdateRoleDTO) (entities.Role, error)
	DeleteRole(id uint64) error
	SearchRoles(filter dto.SearchRoleDTO) ([]entities.Role, int64, error)
}
