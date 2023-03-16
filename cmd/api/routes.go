package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func routes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})

	/*
		=======================
		ACTIVITY GROUPS
		=======================
	*/
	activityRoutes := router.Group("/activity-groups")

	// GET ALL
	activityRoutes.GET("/", h.ActivityGetAll)

	// GET BY ID
	activityRoutes.GET("/:id", h.ActivityGetById)

	// CREATE
	activityRoutes.POST("/", h.ActivityCreate)

	// UPDATE
	activityRoutes.PATCH("/:id", h.ActivityUpdate)

	// DELETE
	activityRoutes.DELETE("/:id", h.ActivityDelete)

	/*
		=======================
		TODO GROUPS
		=======================
	*/
	todoRoutes := router.Group("/todo-items")

	// GET ALL
	todoRoutes.GET("/", h.TodoGetAll)

	// GET BY ID
	todoRoutes.GET("/:id", h.TodoGetById)

	// CREATE
	todoRoutes.POST("/", h.TodoCreate)

	// UPDATE
	todoRoutes.PATCH("/:id", h.TodoUpdate)

	// DELETE
	todoRoutes.DELETE("/:id", h.TodoDelete)
}
