package postgres

import (
	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
)

func (db *PgDB) CreateRole(data dto.CreateRoleDTO) (entities.Role, error) {
	return entities.Role{}, nil
}

func (db *PgDB) RetrieveRole(id uint64) (entities.Role, error) {
	return entities.Role{}, nil
}

func (db *PgDB) UpdateRole(id uint64, data dto.UpdateRoleDTO) (entities.Role, error) {
	return entities.Role{}, nil
}

func (db *PgDB) DeleteRole(id uint64) error {
	return nil
}

func (db *PgDB) SearchRoles(filter dto.SearchRoleDTO) ([]entities.Role, int64, error) {
	return []entities.Role{}, 0, nil
}
