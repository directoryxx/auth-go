//go:build !test
// +build !test

package main

import (
	"os"

	"github.com/directoryxx/auth-go/app/controller"
	"github.com/directoryxx/auth-go/app/repository"
	"github.com/directoryxx/auth-go/app/service"
	"github.com/directoryxx/auth-go/config"
	"github.com/directoryxx/auth-go/helper"
	"github.com/directoryxx/auth-go/infrastructure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	app := SetupInit()
	if os.Getenv("TESTING") != "true" {
		app.Listen(":3000") //excluded
	}

}

func SetupInit() *fiber.App {
	errLoadEnv := godotenv.Load()
	config.GetConfiguration(errLoadEnv)

	dsn := config.GenerateDSNMySQL()
	database, err := infrastructure.OpenDBMysql(dsn)
	// redis := infrastructure.OpenRedis()
	helper.PanicIfError(err)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}[${ip}]:${port} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// root
	root := app.Group("/api")

	// Role
	repoRole := repository.NewRoleRepository(database)
	svcRole := service.NewRoleService(repoRole)
	role := controller.NewRoleController(svcRole, root)
	role.RoleRouter()

	return app
}
