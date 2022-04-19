package seeder

import (
	"github.com/directoryxx/auth-go/app/domain"
	"github.com/directoryxx/auth-go/app/repository"
	"github.com/directoryxx/auth-go/config"
	"github.com/directoryxx/auth-go/infrastructure"
)

func SeederRole() {
	dsn := config.GenerateDSNMySQL()
	database, _ := infrastructure.OpenDBMysql(dsn)
	repoRole := repository.NewRoleRepository(database)

	// Add new struct if want to add more data
	var roleData = []domain.Role{
		{
			Name: "admin",
		},
		{
			Name: "user",
		},
	}

	for _, role := range roleData {
		repoRole.Create(&role)
	}
}
