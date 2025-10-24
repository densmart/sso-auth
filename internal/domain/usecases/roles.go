package usecases

import (
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
)

func CreateRole(repo repo.Repo, data dto.CreateRoleDTO) (*dto.RoleDTO, error) {
	role, err := repo.CreateRole(data)
	if err != nil {
		return nil, err
	}
	response := dto.RoleDTO{
		ID:          role.Id,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		Name:        role.Name,
		Slug:        role.Slug,
		IsPermitted: role.IsPermitted,
	}
	return &response, nil
}
