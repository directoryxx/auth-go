package controller

import (
	"net/http"

	"github.com/directoryxx/auth-go/api/rest/request"
	"github.com/directoryxx/auth-go/api/rest/response"
	"github.com/directoryxx/auth-go/app/helper"
	"github.com/directoryxx/auth-go/app/service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	createRole() fiber.Handler
	updateRole() fiber.Handler
	deleteRole() fiber.Handler
	findByIdRole() fiber.Handler
	findAllRole() fiber.Handler
	UserRouter()
}

type UserControllerImpl struct {
	Ctx     *fiber.Ctx
	Service service.UserService
	Router  fiber.Router
}

func NewUserController(svc service.UserService, app fiber.Router) UserController {
	return &UserControllerImpl{
		Service: svc,
		Router:  app,
	}
}

func (r *UserControllerImpl) UserRouter() {
	r.Router.Post("/register", r.register())
	r.Router.Post("/login", r.login())
	// group := r.Router.Group("user")
	// group.Get("/", r.findAllRole())
	// group.Get("/:id", r.findByIdRole())
	// group.Put("/:id", r.updateRole())
	// group.Delete("/:id", r.deleteRole())
	// group.Post("/", r.createRole())
}

func (r *UserControllerImpl) register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registerReq *request.RegisterUserRequest
		errRequest := c.BodyParser(&registerReq)
		helper.PanicIfError(errRequest)

		registerUser := r.Service.Register(registerReq)
		return c.JSON(&response.DefaultSuccess{
			Data:   registerUser,
			Status: http.StatusOK,
		})
	}
}

func (r *UserControllerImpl) login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginReq *request.LoginUserRequest
		errRequest := c.BodyParser(&loginReq)
		helper.PanicIfError(errRequest)

		loginUser := r.Service.Login(loginReq)

		if loginUser.ID == 0 {
			return c.JSON(&response.LoginResponse{
				Message: "Username/Password Salah",
				Status:  http.StatusOK,
				Data:    nil,
			})
		}

		return c.JSON(&response.LoginResponse{
			Data:    loginUser,
			Status:  http.StatusOK,
			Message: "Berhasil login",
		})
	}
}

func (r *UserControllerImpl) createRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var role *request.RoleRequest
		// errRequest := c.BodyParser(&role)
		// helper.PanicIfError(errRequest)

		// create := r.Service.Create(role)

		// return c.JSON(&response.DefaultSuccess{
		// 	Data:   create,
		// 	Status: http.StatusOK,
		// })
		return c.JSON("ok")
	}
}

func (r *UserControllerImpl) updateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var role *request.RoleRequest
		// errRequest := c.BodyParser(&role)
		// helper.PanicIfError(errRequest)

		// create := r.Service.Create(role)

		// return c.JSON(&response.DefaultSuccess{
		// 	Data:   create,
		// 	Status: http.StatusOK,
		// })
		return c.JSON("ok")
	}
}

func (r *UserControllerImpl) deleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var role *request.RoleRequest
		// errRequest := c.BodyParser(&role)
		// helper.PanicIfError(errRequest)

		// create := r.Service.Create(role)

		// return c.JSON(&response.DefaultSuccess{
		// 	Data:   create,
		// 	Status: http.StatusOK,
		// })
		return c.JSON("ok")
	}
}

func (r *UserControllerImpl) findByIdRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var role *request.RoleRequest
		// errRequest := c.BodyParser(&role)
		// helper.PanicIfError(errRequest)

		// create := r.Service.Create(role)

		// return c.JSON(&response.DefaultSuccess{
		// 	Data:   create,
		// 	Status: http.StatusOK,
		// })
		return c.JSON("ok")
	}
}

func (r *UserControllerImpl) findAllRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var role *request.RoleRequest
		// errRequest := c.BodyParser(&role)
		// helper.PanicIfError(errRequest)

		// create := r.Service.Create(role)

		// return c.JSON(&response.DefaultSuccess{
		// 	Data:   create,
		// 	Status: http.StatusOK,
		// })
		return c.JSON("ok")
	}
}