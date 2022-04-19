package service

import (
	"github.com/directoryxx/fiber-testing/api/rest/request"
	"github.com/directoryxx/fiber-testing/api/rest/response"
	"github.com/directoryxx/fiber-testing/domain"
	"github.com/directoryxx/fiber-testing/repository"
)

type UserService interface {
	Create(user *request.UserRequest) *response.UserResponse
	Update(userReq *request.UserRequest,userid int) *response.UserResponse
	GetById(userid int) *response.UserResponse
	GetAll() *[]response.UserResponse
	Delete(userid int) bool
}

type UserServiceImpl struct{
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

func (us *UserServiceImpl) Create(user *request.UserRequest) *response.UserResponse {
	userCreate := &domain.User{
		Name: user.Name,
		Username: user.Username,
		Password: user.Password,
		RoleID: uint(user.RoleId),
	}

	userCreated := us.UserRepository.Create(userCreate)
	response := &response.UserResponse{
		ID: int(userCreated.ID),
		Name: userCreated.Name,
		Username: userCreated.Username,
		RoleId: int(userCreated.RoleID),
	}
	return response
}

func (us *UserServiceImpl) Update(userReq *request.UserRequest, userid int) *response.UserResponse {
	userUpdate := &domain.User{
		Name: userReq.Name,
		Username: userReq.Username,
		Password: userReq.Password,
		RoleID: uint(userReq.RoleId),
	}

	userUpdated := us.UserRepository.Update(userUpdate,userid)
	response := &response.UserResponse{
		ID: int(userUpdated.ID),
		Name: userUpdated.Name,
		Username: userUpdated.Username,
		RoleId: int(userUpdated.RoleID),
	}
	return response
}

func (us *UserServiceImpl) GetById(userid int) *response.UserResponse {
	user := us.UserRepository.FindById(userid)
	response := &response.UserResponse{
		ID: int(user.ID),
		Name: user.Name,
		Username: user.Username,
		RoleId: int(user.RoleID),
	}
	return response
}

func (us *UserServiceImpl) GetAll() *[]response.UserResponse {
	user := us.UserRepository.FindAll()
	var userResponses []response.UserResponse
	for _, user := range user {
		userResponses = append(userResponses, response.ToUserResponse(user))
	}
	return &userResponses
}

func (us *UserServiceImpl) Delete(userid int) bool {
	user := us.UserRepository.Delete(userid)
	return user
}
