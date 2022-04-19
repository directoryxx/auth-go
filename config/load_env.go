package config

import "os"

func GetConfiguration(err error)  {
	if err != nil {
		os.Setenv("DB_NAME", "test_database")
		os.Setenv("DB_HOST", "127.0.0.1")
		os.Setenv("DB_PORT", "3306")
		os.Setenv("DB_USERNAME", "root")
		os.Setenv("DB_PASSWORD", "rootpass")
	}
}