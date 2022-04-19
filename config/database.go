package config

import "os"

func GenerateDSNMySQL() string {
	dbName := ""
	testing := os.Getenv("TESTING")
	if testing == "true" {
		dbName = os.Getenv("DB_NAME")+"-test"
	} else {
		dbName = os.Getenv("DB_NAME")
	}

	dbDsn := os.Getenv("DB_USERNAME")+ ":"+ os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+dbName+"?charset=utf8mb4&parseTime=True"

	return dbDsn
}
