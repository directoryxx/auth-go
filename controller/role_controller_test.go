package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io"
	"net/http/httptest"
	"os"
	"path"
	"github.com/directoryxx/fiber-testing/api/rest/request"
	"github.com/directoryxx/fiber-testing/config"
	"github.com/directoryxx/fiber-testing/infrastructure"
	"github.com/directoryxx/fiber-testing/repository"
	"github.com/directoryxx/fiber-testing/service"
	"strconv"
	"testing"
)

type Suite struct {
	suite.Suite
	RoleRepo   repository.RoleRepository
	RoleSvc   service.RoleService
	DB *gorm.DB
	app *fiber.App
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/github.com/directoryxx/fiber-testing/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	database,_ := infrastructure.OpenDBMysql(dsn)
	s.RoleRepo = repository.NewRoleRepository(database)
	s.RoleSvc = service.NewRoleService(s.RoleRepo)
	s.DB = database
	s.app = fiber.New()
	s.app.Group("/api")
	role := NewRoleController(s.RoleSvc,s.app)
	role.RoleRouter()
}

func (s *Suite) TestRoleController_createRole() {
	values := map[string]string{"name": "test"}
	json_data, err := json.Marshal(values)
	assert.NoError(s.T(), err)
	req := httptest.NewRequest("POST", "http://localhost:3000/role", bytes.NewBuffer(json_data))
	req.Header.Set("Content-type", "application/json")
	resp, _ := s.app.Test(req)
	assert.Equal(s.T(), "200 OK", resp.Status)
	defer resp.Body.Close()
	defer req.Body.Close()
}

func (s *Suite) TestRoleController_deleteRole() {
	req := &request.RoleRequest{
		Name: "coba",
	}

	role := s.RoleSvc.Create(req)
	roleId := strconv.Itoa(role.ID)
	reqDelete := httptest.NewRequest("DELETE", "http://localhost:3000/role/"+roleId, nil)
	resp, _ := s.app.Test(reqDelete)
	assert.Equal(s.T(), "200 OK", resp.Status)
	defer resp.Body.Close()
	defer reqDelete.Body.Close()
}

func (s *Suite) TestRoleController_findAllRole() {
	req := httptest.NewRequest("GET", "http://localhost:3000/role", nil)
	resp, _ := s.app.Test(req)
	assert.Equal(s.T(), "200 OK", resp.Status)
}

func (s *Suite) TestRoleController_findByIdRole() {
	req := &request.RoleRequest{
		Name: "coba",
	}

	role := s.RoleSvc.Create(req)
	roleId := strconv.Itoa(role.ID)
	reqGET := httptest.NewRequest("GET", "http://localhost:3000/role/"+roleId, nil)
	resp, _ := s.app.Test(reqGET)

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	assert.NoError(s.T(), err)
	fmt.Println("test")
	fmt.Println(string(b))
	assert.Equal(s.T(), "200 OK", resp.Status)
	defer resp.Body.Close()

	req = &request.RoleRequest{
		Name: "coba",
	}

	roleId = strconv.Itoa(1000000)
	reqGET = httptest.NewRequest("GET", "http://localhost:3000/role/"+roleId, nil)
	resp, _ = s.app.Test(reqGET)
	assert.Equal(s.T(), "404 Not Found", resp.Status)
	defer resp.Body.Close()
}

func (s *Suite) TestRoleController_updateRole() {
	req := &request.RoleRequest{
		Name: "coba",
	}

	role := s.RoleSvc.Create(req)
	roleId := strconv.Itoa(role.ID)


	values := map[string]string{"name": "test"}
	json_data, err := json.Marshal(values)
	assert.NoError(s.T(), err)
	reqGET := httptest.NewRequest("PUT", "http://localhost:3000/role/"+roleId, bytes.NewBuffer(json_data))
	reqGET.Header.Set("Content-type","application/json")
	resp, _ := s.app.Test(reqGET)
	assert.Equal(s.T(), "200 OK", resp.Status)
	defer resp.Body.Close()
}

func (s *Suite) TearDownSuite() {
	s.app.Shutdown()
}
