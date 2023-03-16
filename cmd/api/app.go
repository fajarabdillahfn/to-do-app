package api

import (
	"log"

	cDB "github.com/fajarabdillahfn/to-do-app/common/db/mysql"
	"github.com/gin-gonic/gin"
)

func Start() {
	port := ":3030"

	router := gin.Default()
	dbTodo := cDB.OpenDB()
	routes(router, dbTodo)

	log.Println("start application on port " + port)
	router.Run(port)
}
