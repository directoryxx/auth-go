package repository

import (
	"github.com/directoryxx/auth-go/app/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *domain.Role) *domain.Role
	Update(role *domain.Role, roleid int) *domain.Role
	FindById(roleid int) *domain.Role
	FindAll() []domain.Role
	Find(column string, value string) *domain.Role
	Delete(roleid int) bool
}

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		DB: db,
	}
}

func (r *RoleRepositoryImpl) Create(role *domain.Role) *domain.Role {
	r.DB.Create(&role)
	return role
}

func (r *RoleRepositoryImpl) Update(role *domain.Role, roleid int) *domain.Role {
	r.DB.Model(role).Where("id = ?", roleid).Updates(role)
	return role
}

func (r *RoleRepositoryImpl) FindById(roleid int) *domain.Role {
	role := &domain.Role{}
	r.DB.Model(&domain.Role{}).Where("id = ?", roleid).First(role)
	return role
}

func (r *RoleRepositoryImpl) Find(column string, value string) *domain.Role {
	role := &domain.Role{}
	r.DB.Model(&domain.Role{}).Where(column+" = ?", value).First(role)
	return role
}

func (r *RoleRepositoryImpl) FindAll() []domain.Role {
	var Role []domain.Role
	r.DB.Model(&domain.Role{}).Find(&Role)
	return Role
}

func (r *RoleRepositoryImpl) Delete(roleid int) bool {
	r.DB.Delete(&domain.Role{}, roleid)
	return true
}
