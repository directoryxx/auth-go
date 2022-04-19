package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"github.com/directoryxx/fiber-testing/api/rest/request"
	"github.com/directoryxx/fiber-testing/api/rest/response"
	"github.com/directoryxx/fiber-testing/helper"
	"github.com/directoryxx/fiber-testing/service"
)

type RoleController interface {
	createRole() fiber.Handler
	updateRole() fiber.Handler
	deleteRole() fiber.Handler
	findByIdRole() fiber.Handler
	findAllRole() fiber.Handler
	RoleRouter()
}

type RoleControllerImpl struct{
	Ctx *fiber.Ctx
	Service service.RoleService
	Router fiber.Router
}

func NewRoleController(svc service.RoleService, app fiber.Router) RoleController {
	return &RoleControllerImpl{
		Service: svc,
		Router: app,
	}
}

func (r *RoleControllerImpl) RoleRouter()  {
	group := r.Router.Group("role")
	group.Get("/", r.findAllRole())
	group.Get("/:id", r.findByIdRole())
	group.Put("/:id", r.updateRole())
	group.Delete("/:id", r.deleteRole())
	group.Post("/", r.createRole())
}

func (r *RoleControllerImpl) createRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var role *request.RoleRequest
		errRequest := c.BodyParser(&role)
		helper.PanicIfError(errRequest)

		create := r.Service.Create(role)

		return c.JSON(&response.DefaultSuccess{
			Data: create,
			Status: http.StatusOK,
		})
	}
}

func (r *RoleControllerImpl) updateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		var roleReq *request.RoleRequest
		errRequest := c.BodyParser(&roleReq)
		helper.PanicIfError(errRequest)
		role := r.Service.Update(roleReq,id)

		return c.JSON(&response.DefaultSuccess{
			Data: role,
			Status: http.StatusOK,
		})

	}
}

func (r *RoleControllerImpl) deleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		r.Service.Delete(id)

		return c.JSON(&response.DefaultSuccess{
			Data: "true",
			Status: http.StatusOK,
		})
	}
}

func (r *RoleControllerImpl) findByIdRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		role := r.Service.GetById(id)

		if (role.ID == 0){
			c.Status(http.StatusNotFound)
			return c.JSON(&response.DefaultSuccess{
				Data: nil,
				Status: http.StatusNotFound,
			})
		}

		return c.JSON(&response.DefaultSuccess{
			Data: role,
			Status: http.StatusOK,
		})
	}
}

func (r *RoleControllerImpl) findAllRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleAll := r.Service.GetAll()
		return c.JSON(&response.DefaultSuccess{
			Data: roleAll,
			Status: http.StatusOK,
		})
	}
}
