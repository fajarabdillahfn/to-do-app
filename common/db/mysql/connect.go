package mysql

import (
	"fmt"
	"os"

	"github.com/fajarabdillahfn/to-do-app/common/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	if port == "" {
		port = "3306"
	}

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	openConnection := mysql.Open(sqlInfo)

	DB, err := gorm.Open(openConnection, &gorm.Config{})
	DB = DB.Set("gorm:auto_preload", true)
	if err != nil {
		panic("failed to connect database")
	}

	if err := autoMigrate(DB); err != nil {
		panic("failed to migrate: " + err.Error())
	}

	return DB
}

func autoMigrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(&model.Activity{}, &model.Todo{})

	if err != nil {
		return err
	}

	return nil
}
