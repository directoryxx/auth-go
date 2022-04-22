package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/directoryxx/auth-go/api/rest/request"
	"github.com/directoryxx/auth-go/api/rest/response"
	"github.com/directoryxx/auth-go/app/helper"
	"github.com/directoryxx/auth-go/app/middleware"
	"github.com/directoryxx/auth-go/app/service"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserController interface {
	createUser() fiber.Handler
	updateRole() fiber.Handler
	deleteRole() fiber.Handler
	findByIdRole() fiber.Handler
	findAllUser() fiber.Handler
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
	// JWT Middleware
	r.Router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	r.Router.Post("/logout", middleware.JWTProtected(r.Service), r.logout())
	group := r.Router.Group("user")
	group.Get("/", middleware.JWTProtected(r.Service), r.findAllUser())
	// group.Get("/:id", r.findByIdRole())
	// group.Put("/:id", r.updateRole())
	// group.Delete("/:id", r.deleteRole())
	group.Post("/", r.createUser())
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
		id := uuid.New()

		loginUser := r.Service.Login(loginReq)

		if loginUser.ID == 0 {
			return c.JSON(&response.LoginResponse{
				Message: "Username/Password Salah",
				Status:  http.StatusOK,
				Data:    nil,
			})
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name": loginUser.Name,
			"uuid": id.String(),
			"exp":  time.Now().Add(time.Hour * 7).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		r.Service.RememberUuid(loginUser.ID, id.String())

		return c.JSON(&response.LoginResponse{
			Data:    loginUser,
			Status:  http.StatusOK,
			Token:   t,
			Message: "Berhasil login",
		})
	}
}

func (r *UserControllerImpl) logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		r.Service.Logout(claims["uuid"])
		return c.JSON("success")
	}
}

func (r *UserControllerImpl) createUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *request.UserRequest
		errRequest := c.BodyParser(&user)
		helper.PanicIfError(errRequest)

		create := r.Service.Create(user)

		return c.JSON(&response.DefaultSuccess{
			Data:   create,
			Status: http.StatusOK,
		})
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

func (r *UserControllerImpl) findAllUser() fiber.Handler {
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
