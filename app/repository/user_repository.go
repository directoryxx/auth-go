package repository

import (
	"context"
	"time"

	"github.com/directoryxx/auth-go/app/domain"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) *domain.User
	Update(user *domain.User, userid int) *domain.User
	FindById(userid int) *domain.User
	FindAll() []domain.User
	Find(column string, value string) *domain.User
	Delete(userid int) bool
	Set(key string, value string)
	DeleteToken(key string)
	Get(key string) (res string, err error)
}

type UserRepositoryImpl struct {
	DB      *gorm.DB
	Client  *redis.Client
	Context context.Context
}

func NewUserRepository(db *gorm.DB, redis *redis.Client, ctx context.Context) UserRepository {
	return &UserRepositoryImpl{
		DB:      db,
		Client:  redis,
		Context: ctx,
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

func (ur *UserRepositoryImpl) Find(column string, value string) *domain.User {
	user := &domain.User{}
	ur.DB.Model(&domain.User{}).Where(column+" = ?", value).First(user)
	return user
}

func (ur *UserRepositoryImpl) Delete(userid int) bool {
	ur.DB.Delete(&domain.Role{}, userid)
	return true
}

func (ur *UserRepositoryImpl) Set(key string, value string) {
	ur.Client.Set(ur.Context, key, value, time.Hour*7).Err()
}

func (ur *UserRepositoryImpl) Get(key string) (res string,err error) {
	return ur.Client.Get(ur.Context, key).Result()
}

func (ur *UserRepositoryImpl) DeleteToken(key string) {
	ur.Client.Del(ur.Context, key).Err()
}
