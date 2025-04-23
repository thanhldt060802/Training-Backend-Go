package dto

import "training-project/internal/model"

type UserResDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Role     string `json:"role"`
	// Hidden field: Password
}

type UserCreateReqDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	// Hidden field: Role
}

type UserUpdateReqDTO struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Address  *string `json:"address"`
	// Hidden field: Username, Role
}

type UserDeleteReqDTO struct {
	Id int64 `json:"id"`
	// Hidden field: Name, Email, Username, Password, Address, Role
}

func ToUserResDTO(user *model.User) *UserResDTO {
	return &UserResDTO{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Address:  user.Address,
		Role:     user.Role,
	}
}

func ToUserResDTOs(users []model.User) []UserResDTO {
	userResDTOs := make([]UserResDTO, len(users))
	for i, user := range users {
		userResDTOs[i] = *ToUserResDTO(&user)
	}
	return userResDTOs
}
