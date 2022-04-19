package response

import "github.com/directoryxx/fiber-testing/domain"

type RoleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


func ToRoleResponse(role domain.Role) RoleResponse {
	return RoleResponse{
		ID: int(role.ID),
		Name: role.Name,
	}
}