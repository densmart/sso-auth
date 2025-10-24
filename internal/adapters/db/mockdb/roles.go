package mockdb

import (
	"strings"
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
)

var dbData = []entities.Role{
	{
		BaseEntity: entities.BaseEntity{
			Id:        1,
			CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 25, 12, 6, 5, 4, time.UTC),
		},
		Name:        "general manager",
		Slug:        "general-manager",
		IsPermitted: true,
	},
	{
		BaseEntity: entities.BaseEntity{
			Id:        2,
			CreatedAt: time.Date(2025, 5, 25, 12, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 25, 15, 6, 5, 4, time.UTC),
		},
		Name:        "scrum master",
		Slug:        "scrum-master",
		IsPermitted: true,
	},
	{
		BaseEntity: entities.BaseEntity{
			Id:        3,
			CreatedAt: time.Date(2025, 5, 25, 16, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 25, 18, 6, 5, 4, time.UTC),
		},
		Name:        "developer",
		Slug:        "developer",
		IsPermitted: false,
	},
}

func (db *MockDB) CreateRole(data dto.CreateRoleDTO) (entities.Role, error) {
	be := entities.BaseEntity{
		Id:        1,
		CreatedAt: defaultCreatedAt,
		UpdatedAt: defaultCreatedAt,
	}
	return entities.Role{
		BaseEntity:  be,
		Name:        data.Name,
		Slug:        data.Slug,
		IsPermitted: data.IsPermitted,
	}, nil
}

func (db *MockDB) RetrieveRole(id uint64) (entities.Role, error) {
	be := entities.BaseEntity{
		Id:        1,
		CreatedAt: defaultCreatedAt,
		UpdatedAt: time.Now().UTC(),
	}
	return entities.Role{
		BaseEntity:  be,
		Name:        "Mock Role",
		Slug:        "mock-role",
		IsPermitted: true,
	}, nil
}

func (db *MockDB) UpdateRole(id uint64, data dto.UpdateRoleDTO) (entities.Role, error) {
	be := entities.BaseEntity{
		Id:        1,
		CreatedAt: defaultCreatedAt,
		UpdatedAt: time.Now().UTC(),
	}

	var result = entities.Role{
		BaseEntity: be,
	}
	if data.Name != nil {
		result.Name = *data.Name
	}
	if data.Slug != nil {
		result.Slug = *data.Slug
	}
	if data.IsPermitted != nil {
		result.IsPermitted = *data.IsPermitted
	}

	return result, nil
}

func (db *MockDB) DeleteRole(id uint64) error {
	return nil
}

func (db *MockDB) SearchRoles(filter dto.SearchRoleDTO) ([]entities.Role, int64, error) {
	var result []entities.Role
	for _, role := range dbData {
		if filter.Name != nil && !strings.Contains(role.Name, *filter.Name) {
			continue
		}
		if filter.Slug != nil && !strings.Contains(role.Slug, *filter.Slug) {
			continue
		}
		if filter.IsPermitted != nil && role.IsPermitted != *filter.IsPermitted {
			continue
		}
		result = append(result, role)
	}

	return result, 0, nil
}
