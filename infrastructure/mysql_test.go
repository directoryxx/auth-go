package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"github.com/directoryxx/fiber-testing/config"
	"testing"
)

func TestOpenDB(t *testing.T) {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/github.com/directoryxx/fiber-testing/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	db,err := OpenDBMysql(dsn)
	if assert.NoError(t, err) {
		assert.NotNil(t, db)
	}
}
