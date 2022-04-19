package response

import "github.com/directoryxx/fiber-testing/domain"

type UserResponse struct {
	ID int
	Name string
	Username string
	RoleId int
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID: int(user.ID),
		Name: user.Name,
		Username: user.Username,
		RoleId: int(user.RoleID),
	}
}
