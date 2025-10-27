package usecases

// create a test file to test the CreateRole usecase

import (
	"context"
	"testing"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
	"github.com/densmart/sso-auth/internal/utils"
	"github.com/densmart/sso-auth/pkg/configger"
	"github.com/densmart/sso-auth/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// define the input data
	inputData := dto.CreateRoleDTO{
		Name:        "Admin",
		Slug:        "admin",
		IsPermitted: true,
	}

	// define the expected output role
	expectedRole := dto.RoleDTO{
		Name:        "Admin",
		Slug:        "admin",
		IsPermitted: true,
	}

	// call the CreateRole usecase
	result, err := CreateRole(mockRepo, inputData)

	// assert no error occurred
	assert.NoError(t, err)

	// assert the result matches the expected output
	assert.Equal(t, uint64(1), result.ID)
	assert.Equal(t, expectedRole.Name, result.Name)
	assert.Equal(t, expectedRole.Slug, result.Slug)
	assert.Equal(t, expectedRole.IsPermitted, result.IsPermitted)
}

func TestRetrieveRole(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)
	// define the role ID to retrieve
	roleID := uint64(1)

	// define the expected output role
	expectedRole := dto.RoleDTO{
		ID:          1,
		CreatedAt:   "2025-05-25T10:11:12Z",
		Name:        "general manager",
		Slug:        "general-manager",
		IsPermitted: true,
	}
	// call the RetrieveRole usecase
	result, err := RetrieveRole(mockRepo, roleID)
	// assert no error occurred
	assert.NoError(t, err)
	// assert the result matches the expected output
	assert.Equal(t, expectedRole.ID, result.ID)
	assert.Equal(t, expectedRole.CreatedAt, result.CreatedAt)
	assert.Equal(t, expectedRole.Name, result.Name)
	assert.Equal(t, expectedRole.Slug, result.Slug)
	assert.Equal(t, expectedRole.IsPermitted, result.IsPermitted)
}

func TestUpdateRole(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// define the role ID to update
	roleID := uint64(1)

	// define the input data for update
	inputData := dto.UpdateRoleDTO{
		Name:        utils.Ptr("Updated Role"),
		Slug:        utils.Ptr("updated-role"),
		IsPermitted: utils.Ptr(false),
	}

	// define the expected output role
	expectedRole := dto.RoleDTO{
		ID:          1,
		Name:        "Updated Role",
		Slug:        "updated-role",
		IsPermitted: false,
	}

	// call the UpdateRole usecase
	result, err := UpdateRole(mockRepo, roleID, inputData)

	// assert no error occurred
	assert.NoError(t, err)

	// assert the result matches the expected output
	assert.Equal(t, expectedRole.ID, result.ID)
	assert.Equal(t, expectedRole.Name, result.Name)
	assert.Equal(t, expectedRole.Slug, result.Slug)
	assert.Equal(t, expectedRole.IsPermitted, result.IsPermitted)
}

func TestDeleteRole(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// define the role ID to delete
	roleID := uint64(1)

	// call the DeleteRole usecase
	err := DeleteRole(mockRepo, roleID)

	// assert no error occurred
	assert.NoError(t, err)
}

func TestSearchRoles(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	var filters = []dto.SearchRolesDTO{
		{
			ID: utils.Ptr(uint64(1)),
		},
		{
			Name: utils.Ptr("general"),
		},
		{
			Slug: utils.Ptr("scrum"),
		},
		{
			IsPermitted: utils.Ptr(true),
		},
	}

	for filterID, filter := range filters {
		var expectedRoles dto.RolesDTO
		switch filterID {
		case 0:
			expectedRoles = dto.RolesDTO{
				Items: []dto.RoleDTO{
					{
						ID:          1,
						Name:        "general manager",
						Slug:        "general-manager",
						IsPermitted: true,
					},
				},
			}
		case 1:
			expectedRoles = dto.RolesDTO{
				Items: []dto.RoleDTO{
					{
						ID:          1,
						Name:        "general manager",
						Slug:        "general-manager",
						IsPermitted: true,
					},
				},
			}
		case 2:
			expectedRoles = dto.RolesDTO{
				Items: []dto.RoleDTO{
					{
						ID:          2,
						Name:        "scrum master",
						Slug:        "scrum-master",
						IsPermitted: true,
					},
				},
			}
		case 3:
			expectedRoles = dto.RolesDTO{
				Items: []dto.RoleDTO{
					{
						ID:          2,
						Name:        "scrum master",
						Slug:        "scrum-master",
						IsPermitted: true,
					},
					{
						ID:          1,
						Name:        "general manager",
						Slug:        "general-manager",
						IsPermitted: true,
					},
				},
			}
		}

		// call the SearchRoles usecase
		result, err := SearchRoles(mockRepo, filter)

		// assert no error occurred
		assert.NoError(t, err)

		// assert the result matches the expected output
		assert.Equal(t, len(expectedRoles.Items), len(result.Items))
		for i, role := range result.Items {
			assert.Equal(t, expectedRoles.Items[i].ID, role.ID)
			assert.Equal(t, expectedRoles.Items[i].Name, role.Name)
			assert.Equal(t, expectedRoles.Items[i].Slug, role.Slug)
			assert.Equal(t, expectedRoles.Items[i].IsPermitted, role.IsPermitted)
		}
	}
}
