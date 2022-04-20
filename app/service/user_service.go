package service

import (
	"strconv"

	"github.com/directoryxx/auth-go/api/rest/request"
	"github.com/directoryxx/auth-go/api/rest/response"
	"github.com/directoryxx/auth-go/app/domain"
	"github.com/directoryxx/auth-go/app/helper"
	"github.com/directoryxx/auth-go/app/repository"
)

type UserService interface {
	Register(register *request.RegisterUserRequest) *response.UserResponse
	Login(login *request.LoginUserRequest) *response.UserResponse
	Create(user *request.UserRequest) *response.UserResponse
	Update(userReq *request.UserRequest, userid int) *response.UserResponse
	GetById(userid int) *response.UserResponse
	GetAll() *[]response.UserResponse
	Delete(userid int) bool
	RememberUuid(userId int, uuid string)
	Logout(uuid interface{})
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
		RoleRepository: roleRepo,
	}
}

func (us *UserServiceImpl) Register(user *request.RegisterUserRequest) *response.UserResponse {
	passwordGen, _ := helper.GeneratePassword(user.Password)
	userRoleId := us.RoleRepository.Find("name", "user")
	userCreate := &domain.User{
		Name:     user.Name,
		Username: user.Username,
		Password: passwordGen,
		RoleID:   uint(userRoleId.ID),
	}

	userCreated := us.UserRepository.Create(userCreate)
	response := &response.UserResponse{
		ID:       int(userCreated.ID),
		Name:     userCreated.Name,
		Username: userCreated.Username,
		RoleId:   int(userCreated.RoleID),
	}
	return response
}

func (us *UserServiceImpl) Login(user *request.LoginUserRequest) *response.UserResponse {
	userFind := us.UserRepository.Find("username", user.Username)

	comparePassword, _ := helper.ComparePassword(user.Password, userFind.Password)

	if comparePassword {
		response := &response.UserResponse{
			ID:       int(userFind.ID),
			Name:     userFind.Name,
			Username: userFind.Username,
			RoleId:   int(userFind.RoleID),
		}

		return response
	} else {
		response := &response.UserResponse{
			ID:       int(0),
			Name:     "",
			Username: "",
			RoleId:   int(0),
		}

		return response
	}

}

func (us *UserServiceImpl) RememberUuid(userId int, uuid string) {
	us.UserRepository.Set(uuid, strconv.Itoa(userId))
}

func (us *UserServiceImpl) Create(user *request.UserRequest) *response.UserResponse {
	userCreate := &domain.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		RoleID:   uint(user.RoleId),
	}

	userCreated := us.UserRepository.Create(userCreate)
	response := &response.UserResponse{
		ID:       int(userCreated.ID),
		Name:     userCreated.Name,
		Username: userCreated.Username,
		RoleId:   int(userCreated.RoleID),
	}
	return response
}

func (us *UserServiceImpl) Update(userReq *request.UserRequest, userid int) *response.UserResponse {
	userUpdate := &domain.User{
		Name:     userReq.Name,
		Username: userReq.Username,
		Password: userReq.Password,
		RoleID:   uint(userReq.RoleId),
	}

	userUpdated := us.UserRepository.Update(userUpdate, userid)
	response := &response.UserResponse{
		ID:       int(userUpdated.ID),
		Name:     userUpdated.Name,
		Username: userUpdated.Username,
		RoleId:   int(userUpdated.RoleID),
	}
	return response
}

func (us *UserServiceImpl) GetById(userid int) *response.UserResponse {
	user := us.UserRepository.FindById(userid)
	response := &response.UserResponse{
		ID:       int(user.ID),
		Name:     user.Name,
		Username: user.Username,
		RoleId:   int(user.RoleID),
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

func (us *UserServiceImpl) Logout(uuid interface{}) {
	uuidStr := uuid.(string)
	us.UserRepository.DeleteToken(uuidStr)
}
