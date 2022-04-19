package service

import (
	"fmt"
	"github.com/directoryxx/fiber-testing/api/rest/request"
	"github.com/directoryxx/fiber-testing/config"
	"github.com/directoryxx/fiber-testing/domain"
	"github.com/directoryxx/fiber-testing/infrastructure"
	"github.com/directoryxx/fiber-testing/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"path"
	"testing"
)

type Suite struct {
	suite.Suite
	RoleRepo   repository.RoleRepository
	DB *gorm.DB
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/github.com/directoryxx/fiber-testing/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	fmt.Println(dsn)
	database,_ := infrastructure.OpenDBMysql(dsn)
	s.RoleRepo = repository.NewRoleRepository(database)
	s.DB = database
	s.DB.Exec("DELETE FROM roles")
}

func (s *Suite) TestNewRoleService() {
	roleSvc := NewRoleService(s.RoleRepo)
	assert.NotNil(s.T(),roleSvc)
}

func (s *Suite) TestRoleService_Create() {
	requestRole := &request.RoleRequest{
		Name: "coba",
	}
	roleSvc := NewRoleService(s.RoleRepo)
	create := roleSvc.Create(requestRole)
	assert.NotNil(s.T(),create.ID)
}

func (s *Suite) TestRoleService_Delete() {
	roleLast := &domain.Role{}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	roleSvc := NewRoleService(s.RoleRepo)
	delete := roleSvc.Delete(int(roleLast.ID))
	assert.True(s.T(), delete)
}

func (s *Suite) TestRoleService_GetAll() {
	requestRole := &request.RoleRequest{
		Name: "coba",
	}
	roleSvc := NewRoleService(s.RoleRepo)
	create := roleSvc.Create(requestRole)
	assert.NotNil(s.T(),create.ID)
	getAll := roleSvc.GetAll()
	assert.NotNil(s.T(), getAll)
}

func (s *Suite) TestRoleService_GetById() {
	roleLast := &domain.Role{}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	roleSvc := NewRoleService(s.RoleRepo)
	role := roleSvc.GetById(int(roleLast.ID))
	assert.Equal(s.T(), int(role.ID),int(roleLast.ID))
}

func (s *Suite) TestRoleService_Update() {
	roleLast := &domain.Role{}
	requestRole := &request.RoleRequest{
		Name: "cobaRubah",
	}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	roleSvc := NewRoleService(s.RoleRepo)
	role := roleSvc.Update(requestRole,int(roleLast.ID))
	assert.Equal(s.T(), role.Name,requestRole.Name)
}

func (s *Suite) TearDownSuite() {
	s.DB.Exec("DELETE FROM roles")
}