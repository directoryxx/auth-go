package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateDSNMySQL(t *testing.T) {
	dbName := "coba-db"
	dbUsername := "coba-user"
	dbPassword := "coba-pass"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	os.Setenv("DB_NAME",dbName)
	os.Setenv("DB_USERNAME",dbUsername)
	os.Setenv("DB_PASSWORD",dbPassword)
	os.Setenv("DB_HOST",dbHost)
	os.Setenv("DB_PORT",dbPort)
	dsn := GenerateDSNMySQL()
	manualGenerate := dbUsername+ ":"+ dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"-test"+"?charset=utf8mb4&parseTime=True"
	assert.Equal(t,dsn,manualGenerate)
	os.Setenv("TESTING","")
	dsn = GenerateDSNMySQL()
	manualGenerate = dbUsername+ ":"+ dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True"
	assert.Equal(t,dsn,manualGenerate)

}
