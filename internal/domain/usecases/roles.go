package usecases

import (
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
	"github.com/densmart/sso-auth/pkg/paginator"
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

func RetrieveRole(repo repo.Repo, id uint64) (*dto.RoleDTO, error) {
	role, err := repo.RetrieveRole(id)
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

func UpdateRole(repo repo.Repo, id uint64, data dto.UpdateRoleDTO) (*dto.RoleDTO, error) {
	role, err := repo.UpdateRole(id, data)
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

func DeleteRole(repo repo.Repo, id uint64) error {
	return repo.DeleteRole(id)
}

func SearchRoles(repo repo.Repo, filter dto.SearchRolesDTO) (*dto.RolesDTO, error) {
	pag := paginator.NewPaginator(filter.BaseSearchRequestDTO)
	offset := pag.GetOffset()
	limit := pag.GetLimit()
	filter.Offset = &offset
	filter.Limit = &limit

	roles, _, err := repo.SearchRoles(filter)
	if err != nil {
		return nil, err
	}

	response := dto.RolesDTO{
		Pagination: pag.ToLinkHeader(),
	}
	for _, role := range roles {
		roleDTO := dto.RoleDTO{
			ID:          role.Id,
			CreatedAt:   role.CreatedAt.Format(time.RFC3339),
			Name:        role.Name,
			Slug:        role.Slug,
			IsPermitted: role.IsPermitted,
		}
		response.Items = append(response.Items, roleDTO)
	}

	return &response, nil
}
