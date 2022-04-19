package repository

import (
	"github.com/directoryxx/auth-go/app/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) *domain.User
	Update(user *domain.User, userid int) *domain.User
	FindById(userid int) *domain.User
	FindAll() []domain.User
	Delete(userid int) bool
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (ur *UserRepositoryImpl) Create(user *domain.User) *domain.User {
	ur.DB.Create(&user)
	return user
}

func (ur *UserRepositoryImpl) Update(user *domain.User, userid int) *domain.User {
	ur.DB.Model(user).Where("id = ?", userid).Updates(user)
	return user
}

func (ur *UserRepositoryImpl) FindById(userid int) *domain.User {
	user := &domain.User{}
	ur.DB.Model(&domain.User{}).Where("id = ?", userid).First(user)
	return user
}

func (ur *UserRepositoryImpl) FindAll() []domain.User {
	var User []domain.User
	ur.DB.Model(&domain.User{}).Find(&User)
	return User
}

func (ur *UserRepositoryImpl) Delete(userid int) bool {
	ur.DB.Delete(&domain.Role{}, userid)
	return true
}
