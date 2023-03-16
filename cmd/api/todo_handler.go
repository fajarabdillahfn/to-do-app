package api

import (
	"fmt"
	"net/http"

	"github.com/fajarabdillahfn/to-do-app/common/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) TodoGetAll(ctx *gin.Context) {
	activityId := ctx.Query("activity_group_id")

	var todos []*model.Todo

	if activityId != "" {
		h.DB = h.DB.Where("activity_group_id = ?", activityId)
	}

	h.DB.WithContext(ctx).Find(&todos)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (h *handler) TodoGetById(ctx *gin.Context) {
	id := ctx.Param("id")

	var todo *model.Todo

	res := h.DB.WithContext(ctx).First(&todo, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func (h *handler) TodoCreate(ctx *gin.Context) {
	requestData := model.AddTodoRequest{}

	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
		return
	}

	var todo *model.Todo = &model.Todo{
		Title:           requestData.Title,
		ActivityGroupId: requestData.ActivityGroupId,
		IsActive:        requestData.IsActive,
	}

	h.DB.WithContext(ctx).Create(&todo)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func (h *handler) TodoUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var todo *model.Todo

	res := h.DB.WithContext(ctx).First(&todo, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		})
		return
	}

	requestData := model.UpdateTodoRequest{}

	ctx.BindJSON(&requestData)

	todo.Title = requestData.Title
	todo.ActivityGroupId = requestData.ActivityGroupId
	todo.Priority = requestData.Priority
	todo.IsActive = requestData.IsActive

	h.DB.WithContext(ctx).Save(&todo)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func (h *handler) TodoDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var todo *model.Todo

	res := h.DB.WithContext(ctx).First(&todo, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		})
		return
	}

	h.DB.WithContext(ctx).Delete(&todo)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
	})
}
