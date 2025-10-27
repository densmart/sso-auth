package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
	"github.com/densmart/sso-auth/internal/utils"
	"github.com/densmart/sso-auth/pkg/configger"
	"github.com/densmart/sso-auth/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	userPassword := "secure-user-password"

	// define the input data
	// define the input data
	inputData := dto.CreateUserDTO{
		Email:     "admin@admin.com",
		FirstName: "Admin",
		LastName:  "User",
		Phone:     utils.Ptr("1234567890"),
		IsActive:  true,
		Is2fa:     false,
		RoleID:    1,
		Password:  userPassword,
	}

	// define the expected output user
	expectedUser := dto.UserDTO{
		Email:     "admin@admin.com",
		FirstName: "Admin",
		LastName:  "User",
		Phone:     utils.Ptr("1234567890"),
		IsActive:  true,
		Is2fa:     false,
		RoleID:    1,
	}

	// call the CreateUser usecase
	result, err := CreateUser(mockRepo, inputData)

	// assert no error occurred
	assert.NoError(t, err)

	// check password hash
	assert.NoError(t, err)

	// assert the result matches the expected output
	assert.Equal(t, uint64(1), result.ID)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.FirstName, result.FirstName)
	assert.Equal(t, expectedUser.LastName, result.LastName)
	assert.Equal(t, expectedUser.Phone, result.Phone)
	assert.Equal(t, expectedUser.IsActive, result.IsActive)
	assert.Equal(t, expectedUser.Is2fa, result.Is2fa)
	assert.Equal(t, expectedUser.RoleID, result.RoleID)
}

func TestRetrieveUser(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// call the RetrieveUser usecase
	result, err := RetrieveUser(mockRepo, 1)

	// assert no error occurred
	assert.NoError(t, err)

	// define the expected output user
	expectedUser := dto.UserDTO{
		Email:     "john@doe.com",
		FirstName: "John",
		LastName:  "Doe",
		Phone:     utils.Ptr("+1234567890"),
		IsActive:  true,
		Is2fa:     false,
		RoleID:    1,
	}

	// assert the result matches the expected output
	assert.Equal(t, uint64(1), result.ID)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.FirstName, result.FirstName)
	assert.Equal(t, expectedUser.LastName, result.LastName)
	assert.Equal(t, expectedUser.Phone, result.Phone)
	assert.Equal(t, expectedUser.IsActive, result.IsActive)
	assert.Equal(t, expectedUser.Is2fa, result.Is2fa)
	assert.Equal(t, expectedUser.RoleID, result.RoleID)
}

func TestUpdateUser(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// define the input data
	inputData := dto.UpdateUserDTO{
		Email:     utils.Ptr("johnny@doer.com"),
		FirstName: utils.Ptr("Johnny"),
		LastName:  utils.Ptr("Doer"),
		Phone:     utils.Ptr("+0987654321"),
		IsActive:  utils.Ptr(false),
		RoleID:    utils.Ptr(uint64(2)),
	}

	// define the expected output user
	expectedUser := dto.UserDTO{
		Email:     "johnny@doer.com",
		FirstName: "Johnny",
		LastName:  "Doer",
		Phone:     utils.Ptr("+0987654321"),
		IsActive:  false,
		Is2fa:     false,
		RoleID:    2,
	}

	// call the UpdateUser usecase
	result, err := UpdateUser(mockRepo, 1, inputData)

	// assert no error occurred
	assert.NoError(t, err)

	// assert the result matches the expected output
	assert.Equal(t, uint64(1), result.ID)
	assert.Equal(t, expectedUser.Email, result.Email)
	assert.Equal(t, expectedUser.FirstName, result.FirstName)
	assert.Equal(t, expectedUser.LastName, result.LastName)
	assert.Equal(t, expectedUser.Phone, result.Phone)
	assert.Equal(t, expectedUser.IsActive, result.IsActive)
	assert.Equal(t, expectedUser.Is2fa, result.Is2fa)
	assert.Equal(t, expectedUser.RoleID, result.RoleID)
}

func TestDeleteUser(t *testing.T) {
	// initialize config
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	// call the DeleteUser usecase
	err := mockRepo.DeleteUser(1)

	// assert no error occurred
	assert.NoError(t, err)
}

func TestSearchUsers(t *testing.T) {
	configger.InitConfig("../../../config", "config", "yaml")
	logger.InitLogger()
	// create a mock repo
	ctx := context.Background()
	mockRepo, rErr := repo.NewRepo(ctx, "mockdb")
	// assert no error occurred
	assert.NoError(t, rErr)

	var filters = []dto.SearchUsersDTO{
		{
			ID: utils.Ptr(uint64(1)),
		},
		{
			Email: utils.Ptr("will"),
		},
		{
			FirstName: utils.Ptr("Jo"),
		},
		{
			LastName: utils.Ptr("rrow"),
		},
		{
			Phone: utils.Ptr("+1234567890"),
		},
		{
			IsActive: utils.Ptr(true),
		},
		{
			Is2fa: utils.Ptr(false),
		},
		{
			RoleID: utils.Ptr(uint64(1)),
		},
	}

	for filterID, filter := range filters {
		var expectedUsers dto.UsersDTO
		switch filterID {
		case 0:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 1:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        2,
						CreatedAt: time.Date(2025, 5, 26, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "will@turner.com",
						FirstName: "Will",
						LastName:  "Turner",
						Phone:     nil,
						IsActive:  true,
						Is2fa:     true,
						RoleID:    2,
					},
				},
			}
		case 2:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 3:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        3,
						CreatedAt: time.Date(2025, 5, 27, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "jack@sparrow.com",
						FirstName: "Jack",
						LastName:  "Sparrow",
						Phone:     nil,
						IsActive:  false,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 4:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 5:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        2,
						CreatedAt: time.Date(2025, 5, 26, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "will@turner.com",
						FirstName: "Will",
						LastName:  "Turner",
						Phone:     nil,
						IsActive:  true,
						Is2fa:     true,
						RoleID:    2,
					},
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 6:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        3,
						CreatedAt: time.Date(2025, 5, 27, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "jack@sparrow.com",
						FirstName: "Jack",
						LastName:  "Sparrow",
						Phone:     nil,
						IsActive:  false,
						Is2fa:     false,
						RoleID:    1,
					},
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		case 7:
			expectedUsers = dto.UsersDTO{
				Items: []dto.UserDTO{
					{
						ID:        3,
						CreatedAt: time.Date(2025, 5, 27, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "jack@sparrow.com",
						FirstName: "Jack",
						LastName:  "Sparrow",
						Phone:     nil,
						IsActive:  false,
						Is2fa:     false,
						RoleID:    1,
					},
					{
						ID:        1,
						CreatedAt: time.Date(2025, 5, 25, 10, 11, 12, 13, time.UTC).Format(time.RFC3339),
						Email:     "john@doe.com",
						FirstName: "John",
						LastName:  "Doe",
						Phone:     utils.Ptr("+1234567890"),
						IsActive:  true,
						Is2fa:     false,
						RoleID:    1,
					},
				},
			}
		}
		// call the SearchRoles usecase
		result, err := SearchUsers(mockRepo, filter)

		// assert no error occurred
		assert.NoError(t, err)

		// assert the result matches the expected output
		assert.Equal(t, len(expectedUsers.Items), len(result.Items))
		for i, user := range result.Items {
			assert.Equal(t, expectedUsers.Items[i].ID, user.ID)
			assert.Equal(t, expectedUsers.Items[i].Email, user.Email)
			assert.Equal(t, expectedUsers.Items[i].FirstName, user.FirstName)
			assert.Equal(t, expectedUsers.Items[i].LastName, user.LastName)
			assert.Equal(t, expectedUsers.Items[i].Phone, user.Phone)
			assert.Equal(t, expectedUsers.Items[i].IsActive, user.IsActive)
			assert.Equal(t, expectedUsers.Items[i].Is2fa, user.Is2fa)
			assert.Equal(t, expectedUsers.Items[i].RoleID, user.RoleID)
		}
	}
}
