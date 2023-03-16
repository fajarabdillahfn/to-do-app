package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	port := ":3030"

	router := gin.Default()
	routes(router)

	log.Println("start application on port " + port)
	router.Run(port)
}
