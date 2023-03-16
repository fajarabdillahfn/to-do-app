package api

import (
	cDB "github.com/fajarabdillahfn/to-do-app/common/db/mysql"
	"gorm.io/gorm"
)

var (
	dbTodo *gorm.DB
)

func init() {
	dbTodo = cDB.OpenDB()
}