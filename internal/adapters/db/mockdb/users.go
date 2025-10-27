package mockdb

import (
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/entities"
	"github.com/densmart/sso-auth/internal/utils"
)

var mockUsersData = []entities.User{
	{
		BaseEntity: entities.BaseEntity{
			Id:        1,
			CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 25, 12, 6, 5, 4, time.UTC),
		},
		Email:       "john@doe.com",
		Password:    "some-hashed-password",
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       utils.Ptr("+1234567890"),
		IsActive:    true,
		Is2fa:       false,
		Token2fa:    nil,
		LastLoginAt: time.Date(2025, 5, 26, 9, 0, 0, 0, time.UTC),
		RoleID:      1,
	},
	{
		BaseEntity: entities.BaseEntity{
			Id:        2,
			CreatedAt: time.Date(2025, 5, 26, 10, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 26, 12, 6, 5, 4, time.UTC),
		},
		Email:       "will@turner.com",
		Password:    "another-hashed-password",
		FirstName:   "Will",
		LastName:    "Turner",
		Phone:       nil,
		IsActive:    true,
		Is2fa:       true,
		Token2fa:    utils.Ptr("123456"),
		LastLoginAt: time.Date(2025, 5, 27, 9, 0, 0, 0, time.UTC),
		RoleID:      2,
	},
	{
		BaseEntity: entities.BaseEntity{
			Id:        3,
			CreatedAt: time.Date(2025, 5, 27, 10, 11, 12, 13, time.UTC),
			UpdatedAt: time.Date(2025, 5, 27, 12, 6, 5, 4, time.UTC),
		},
		Email:       "jack@sparrow.com",
		Password:    "weird-hashed-password",
		FirstName:   "Jack",
		LastName:    "Sparrow",
		Phone:       nil,
		IsActive:    false,
		Is2fa:       false,
		Token2fa:    nil,
		LastLoginAt: time.Date(2025, 5, 28, 9, 0, 0, 0, time.UTC),
		RoleID:      1,
	},
}

func (db *MockDB) CreateUser(data dto.CreateUserDTO) (entities.User, error) {
	be := entities.BaseEntity{
		Id:        1,
		CreatedAt: defaultCreatedAt,
		UpdatedAt: defaultCreatedAt,
	}
	return entities.User{
		BaseEntity: be,
		Email:      data.Email,
		Password:   data.Password,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Phone:      data.Phone,
		IsActive:   data.IsActive,
		Is2fa:      data.Is2fa,
		Token2fa:   data.Token2fa,
		RoleID:     data.RoleID,
	}, nil
}

func (db *MockDB) RetrieveUser(id uint64) (entities.User, error) {
	for _, user := range mockUsersData {
		if user.Id == id {
			return user, nil
		}
	}
	return entities.User{}, errors.New("user not found")
}

func (db *MockDB) UpdateUser(id uint64, data dto.UpdateUserDTO) (entities.User, error) {
	for _, user := range mockUsersData {
		if user.Id == id {
			be := entities.BaseEntity{
				Id:        1,
				CreatedAt: defaultCreatedAt,
				UpdatedAt: time.Now().UTC(),
			}

			var result = entities.User{
				BaseEntity: be,
			}
			if data.Email != nil {
				result.Email = *data.Email
			}
			if data.FirstName != nil {
				result.FirstName = *data.FirstName
			}
			if data.LastName != nil {
				result.LastName = *data.LastName
			}
			if data.Phone != nil {
				result.Phone = data.Phone
			}
			if data.IsActive != nil {
				result.IsActive = *data.IsActive
			}
			if data.Is2fa != nil {
				result.Is2fa = *data.Is2fa
			}
			if data.RoleID != nil {
				result.RoleID = *data.RoleID
			}
			if data.LastLoginAt != nil {
				result.LastLoginAt = *data.LastLoginAt
			}
			if data.Password != nil {
				result.Password = *data.Password
			}
			return result, nil
		}
	}
	return entities.User{}, errors.New("user not found")
}

func (db *MockDB) DeleteUser(id uint64) error {
	for _, user := range mockUsersData {
		if user.Id == id {
			return nil
		}
	}
	return errors.New("user not found")
}

func (db *MockDB) SearchUsers(filter dto.SearchUsersDTO) ([]entities.User, uint64, error) {
	var result []entities.User
	for _, user := range mockUsersData {
		match := true
		if filter.ID != nil && user.Id != *filter.ID {
			match = false
		}
		if filter.Email != nil && !strings.Contains(user.Email, *filter.Email) {
			match = false
		}
		if filter.FirstName != nil && !strings.Contains(user.FirstName, *filter.FirstName) {
			match = false
		}
		if filter.LastName != nil && !strings.Contains(user.LastName, *filter.LastName) {
			match = false
		}
		if filter.Phone != nil {
			if user.Phone == nil || !strings.Contains(*user.Phone, *filter.Phone) {
				match = false
			}
		}
		if filter.IsActive != nil && user.IsActive != *filter.IsActive {
			match = false
		}
		if filter.Is2fa != nil && user.Is2fa != *filter.Is2fa {
			match = false
		}
		if filter.RoleID != nil && user.RoleID != *filter.RoleID {
			match = false
		}
		if match {
			result = append(result, user)
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
		case "email":
			if strings.ToLower(sortOrder) == "desc" {
				return strings.ToLower(result[i].Email) > strings.ToLower(result[j].Email)
			}
			return strings.ToLower(result[i].Email) < strings.ToLower(result[j].Email)
		case "first_name":
			if strings.ToLower(sortOrder) == "desc" {
				return strings.ToLower(result[i].FirstName) > strings.ToLower(result[j].FirstName)
			}
			return strings.ToLower(result[i].FirstName) < strings.ToLower(result[j].FirstName)
		case "last_name":
			if strings.ToLower(sortOrder) == "desc" {
				return strings.ToLower(result[i].LastName) > strings.ToLower(result[j].LastName)
			}
			return strings.ToLower(result[i].LastName) < strings.ToLower(result[j].LastName)
		case "created_at":
			if strings.ToLower(sortOrder) == "desc" {
				return result[i].CreatedAt.After(result[j].CreatedAt)
			}
			return result[i].CreatedAt.Before(result[j].CreatedAt)
		case "is_active":
			if strings.ToLower(sortOrder) == "desc" {
				return !result[i].IsActive && result[j].IsActive
			}
			return result[i].IsActive && !result[j].IsActive
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
			result = []entities.User{}
		}
	}

	return result, total, nil
}
