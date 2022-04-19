package infrastructure

import (
	"github.com/directoryxx/auth-go/app/domain"
	"github.com/directoryxx/auth-go/app/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDBMysql(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	helper.PanicIfError(err)

	autoMigrate(db)

	return db, err
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.User{})
}
