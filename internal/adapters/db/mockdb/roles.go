package mockdb

import (
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
)

var mockRolesData = []entities.Role{
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
	for _, role := range mockRolesData {
		if role.Id == id {
			return role, nil
		}
	}
	return entities.Role{}, errors.New("role not found")
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
	for _, role := range mockRolesData {
		if role.Id == id {
			return nil
		}
	}
	return errors.New("role not found")
}

func (db *MockDB) SearchRoles(filter dto.SearchRolesDTO) ([]entities.Role, uint64, error) {
	var result []entities.Role
	for _, role := range mockRolesData {
		match := false
		if filter.ID != nil && role.Id == *filter.ID {
			match = true
		}
		if filter.Name != nil && strings.Contains(role.Name, *filter.Name) {
			match = true
		}
		if filter.Slug != nil && strings.Contains(role.Slug, *filter.Slug) {
			match = true
		}
		if filter.IsPermitted != nil && role.IsPermitted == *filter.IsPermitted {
			match = true
		}
		if !filter.CreatedAtFrom.IsZero() {
			if role.CreatedAt.After(filter.CreatedAtFrom) {
				match = true
			}
		}
		if !filter.CreatedAtTo.IsZero() {
			if role.CreatedAt.Before(filter.CreatedAtTo) {
				match = true
			}
		}
		if match {
			result = append(result, role)
		}
	}
	total := uint64(len(result))

	// ordering
	sortField := "id"
	sortOrder := "desc"

	if filter.OrderBy != nil {
		sortField = *filter.OrderBy
		if sortField[0:1] == "-" {
			sortField = sortField[1:]
		} else {
			sortOrder = "asc"
		}
	}

	sort.Slice(result, func(i, j int) bool {
		switch strings.ToLower(sortField) {
		case "id":
			if strings.ToLower(sortOrder) == "desc" {
				return result[i].Id > result[j].Id
			}
			return result[i].Id < result[j].Id
		case "name":
			if strings.ToLower(sortOrder) == "desc" {
				return strings.ToLower(result[i].Name) > strings.ToLower(result[j].Name)
			}
			return strings.ToLower(result[i].Name) < strings.ToLower(result[j].Name)
		case "slug":
			if strings.ToLower(sortOrder) == "desc" {
				return strings.ToLower(result[i].Slug) > strings.ToLower(result[j].Slug)
			}
			return strings.ToLower(result[i].Slug) < strings.ToLower(result[j].Slug)
		case "created_at":
			if strings.ToLower(sortOrder) == "desc" {
				return result[i].CreatedAt.After(result[j].CreatedAt)
			}
			return result[i].CreatedAt.Before(result[j].CreatedAt)
		default:
			// По умолчанию сортируем по ID
			return result[i].Id < result[j].Id
		}
	})

	if filter.Limit != nil {
		limit := *filter.Limit
		if limit < uint(len(result)) {
			result = result[:limit]
		}
	}
	if filter.Offset != nil {
		offset := *filter.Offset
		if offset < uint(len(result)) {
			result = result[offset:]
		} else {
			result = []entities.Role{}
		}
	}

	return result, total, nil
}
