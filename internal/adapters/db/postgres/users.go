package postgres

import (
	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
)

func (db *PgDB) CreateUser(data dto.CreateUserDTO) (entities.User, error) {
	return entities.User{}, nil
}

func (db *PgDB) RetrieveUser(id uint64) (entities.User, error) {
	return entities.User{}, nil
}

func (db *PgDB) UpdateUser(id uint64, data dto.UpdateUserDTO) (entities.User, error) {
	return entities.User{}, nil
}

func (db *PgDB) DeleteUser(id uint64) error {
	return nil
}

func (db *PgDB) SearchUsers(filter dto.SearchUsersDTO) ([]entities.User, uint64, error) {
	return []entities.User{}, 0, nil
}
