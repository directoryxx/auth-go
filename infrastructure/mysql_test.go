package infrastructure

import (
	"os"
	"path"
	"testing"

	"github.com/directoryxx/auth-go/config"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/github.com/directoryxx/auth-go/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	db, err := OpenDBMysql(dsn)
	if assert.NoError(t, err) {
		assert.NotNil(t, db)
	}
}
