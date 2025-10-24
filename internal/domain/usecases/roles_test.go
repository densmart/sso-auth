package usecases

// create a test file to test the CreateRole usecase

import (
	"context"
	"testing"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
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
