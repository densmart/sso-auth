package usecases

import (
	"github.com/densmart/sso-auth/internal/adapters/dto"
	"github.com/densmart/sso-auth/internal/domain/repo"
	"github.com/densmart/sso-auth/internal/utils"
	"github.com/densmart/sso-auth/pkg/paginator"
)

func CreateUser(repo repo.Repo, data dto.CreateUserDTO) (*dto.UserDTO, error) {
	// create password hash
	pwdHash, err := utils.GeneratePasswordHash(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = pwdHash

	if data.Is2fa == true {
		// generate initial 2fa token
		token2fa, err := utils.CreateOtpSecret()
		if err != nil {
			return nil, err
		}
		data.Token2fa = &token2fa
	}

	user, err := repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	response := dto.UserDTO{
		ID:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}
	return &response, nil
}

func RetrieveUser(repo repo.Repo, id uint64) (*dto.UserDTO, error) {
	user, err := repo.RetrieveUser(id)
	if err != nil {
		return nil, err
	}
	response := dto.UserDTO{
		ID:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}
	return &response, nil
}

func UpdateUser(repo repo.Repo, id uint64, data dto.UpdateUserDTO) (*dto.UserDTO, error) {
	user, err := repo.UpdateUser(id, data)
	if err != nil {
		return nil, err
	}
	response := dto.UserDTO{
		ID:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}
	return &response, nil
}

func DeleteUser(repo repo.Repo, id uint64) error {
	return repo.DeleteUser(id)
}

func SearchUsers(repo repo.Repo, filter dto.SearchUsersDTO) (*dto.UsersDTO, error) {
	pag := paginator.NewPaginator(filter.BaseSearchRequestDTO)
	offset := pag.GetOffset()
	limit := pag.GetLimit()
	filter.Offset = &offset
	filter.Limit = &limit

	users, _, err := repo.SearchUsers(filter)
	if err != nil {
		return nil, err
	}

	response := dto.UsersDTO{
		Pagination: pag.ToLinkHeader(),
	}

	for _, user := range users {
		response.Items = append(response.Items, dto.UserDTO{
			ID:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			IsActive:  user.IsActive,
			Is2fa:     user.Is2fa,
			RoleID:    user.RoleID,
		})
	}

	return &response, nil
}
