package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"safakkizkin/httputil"
	"safakkizkin/models"
)

func GetTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var task models.Task
	err := models.GetTask(&task, id)
	if err != nil || task.Model.ID <= 0 {
		httputil.NewError(ctx, http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func GetTasks(ctx *gin.Context) {
	var tasks []models.Task
	err := models.GetAllTasks(&tasks)
	if err != nil {
		httputil.NewError(ctx, http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var task models.Task

	if task.StartDate.After(task.EndDate) {
		err := errors.New("Task: Start date can not be after end date.")
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	err := models.DeleteTask(&task, id)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func AddNewTask(ctx *gin.Context) {
	var task models.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err = models.AddNewTask(&task)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}
