package api

import (
	"fmt"
	"net/http"

	"github.com/fajarabdillahfn/to-do-app/common/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) ActivityGetAll(ctx *gin.Context) {
	var activities []*model.Activity

	h.DB.WithContext(ctx).Find(&activities)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (h *handler) ActivityGetById(ctx *gin.Context) {
	id := ctx.Param("id")

	var activity *model.Activity

	res := h.DB.WithContext(ctx).First(&activity, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (h *handler) ActivityCreate(ctx *gin.Context) {
	requestData := model.AddActivityRequest{}

	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:  "Bad Request",
			Message: "invalid input",
		})
		return
	}

	if requestData.Title == "" {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
		return
	}

	var activity *model.Activity = &model.Activity{
		Title: requestData.Title,
		Email: requestData.Email,
	}

	h.DB.WithContext(ctx).Create(&activity)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (h *handler) ActivityUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var activity *model.Activity

	res := h.DB.WithContext(ctx).First(&activity, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	requestData := model.UpdateActivityRequest{}

	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
		})
		return
	}

	activity.Title = requestData.Title

	h.DB.WithContext(ctx).Save(&activity)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (h *handler) ActivityDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var activity *model.Activity

	res := h.DB.WithContext(ctx).First(&activity, id)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	h.DB.WithContext(ctx).Delete(&activity)

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Status:  "Success",
		Message: "Success",
	})
}
