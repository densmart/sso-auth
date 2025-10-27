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
	SearchRoles(filter dto.SearchRolesDTO) ([]entities.Role, uint64, error)
}

type Users interface {
	CreateUser(data dto.CreateUserDTO) (entities.User, error)
	RetrieveUser(id uint64) (entities.User, error)
	UpdateUser(id uint64, data dto.UpdateUserDTO) (entities.User, error)
	DeleteUser(id uint64) error
	SearchUsers(filter dto.SearchUsersDTO) ([]entities.User, uint64, error)
}
