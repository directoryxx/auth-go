package service

import (
	"github.com/directoryxx/fiber-testing/api/rest/request"
	"github.com/directoryxx/fiber-testing/api/rest/response"
	"github.com/directoryxx/fiber-testing/domain"
	"github.com/directoryxx/fiber-testing/repository"
)

type RoleService interface {
	Create(role *request.RoleRequest) *response.RoleResponse
	Update(roleReq *request.RoleRequest,roleId int) *response.RoleResponse
	GetById(roleId int) *response.RoleResponse
	GetAll() *[]response.RoleResponse
	Delete(roleId int) bool
}

type RoleServiceImpl struct{
	RoleRepository repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepo,
	}
}

func (svc *RoleServiceImpl) Create(role *request.RoleRequest) *response.RoleResponse {
	roleCreate := &domain.Role{
		Name: role.Name,
	}

	roleCreated := svc.RoleRepository.Create(roleCreate)
	response := &response.RoleResponse{
		ID: int(roleCreated.ID),
		Name: roleCreated.Name,
	}
	return response
}

func (svc *RoleServiceImpl) Update(roleReq *request.RoleRequest,roleId int) *response.RoleResponse {
	roleUpdate := &domain.Role{
		Name: roleReq.Name,
	}

	roleUpdated := svc.RoleRepository.Update(roleUpdate, roleId)
	response := &response.RoleResponse{
		ID: roleId,
		Name: roleUpdated.Name,
	}
	return response
}

func (svc *RoleServiceImpl) GetById(roleId int) *response.RoleResponse {
	role := svc.RoleRepository.FindById(roleId)
	response := &response.RoleResponse{
		ID: int(role.ID),
		Name: role.Name,
	}
	return response
}

func (svc *RoleServiceImpl) GetAll() *[]response.RoleResponse {
	role := svc.RoleRepository.FindAll()
	var roleResponses []response.RoleResponse
	for _, role := range role {
		roleResponses = append(roleResponses, response.ToRoleResponse(role))
	}
	return &roleResponses
}

func (svc *RoleServiceImpl) Delete(roleId int) bool {
	role := svc.RoleRepository.Delete(roleId)
	return role
}

