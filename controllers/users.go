package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safakkizkin/httputil"
	"safakkizkin/models"
)

func GetOneUserByMail(ctx *gin.Context) {
	var user models.User
	user.Mail = ctx.Params.ByName("mail")
	err := models.GetUser(&user)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetUsers(ctx *gin.Context) {
	var users []models.User
	err := models.GetAllUsers(&users)
	if err != nil {
		httputil.NewError(ctx, http.StatusNoContent, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func AddNewUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	err := models.AddNewUser(&user)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	mail := ctx.Params.ByName("mail")
	var user models.User
	err := models.DeleteUser(&user, mail)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
