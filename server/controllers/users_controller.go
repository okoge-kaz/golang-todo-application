package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
	"github.com/okoge-kaz/golang-todo-application/server/models"
)

func CreateUser(ctx *gin.Context) {
	// get user
	var user entities.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user)")
		return
	}

	// create user
	err = models.CreateUser(&user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// set session
	err = helpers.SetUserID(ctx, user.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// render
	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func ShowUser(ctx *gin.Context) {
	// get user
	userID, err := helpers.GetIDFromQueryParameter(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user id)")
		return
	}

	user, err := models.GetUser(userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// render
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func ChangeUserPassword(ctx *gin.Context) {
	// get user
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user id)")
		return
	}

	var user entities.User
	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user)")
		return
	}

	// change password
	err = models.ChangeUserPassword(userID, user.Password)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// render
	ctx.Status(http.StatusNoContent)
}

func ChangeUserName(ctx *gin.Context) {
	// get user
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user id)")
		return
	}

	var user entities.User
	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user)")
		return
	}

	// change name
	err = models.ChangeUserName(userID, user.Name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// render
	ctx.Status(http.StatusNoContent)
}

func DeleteUser(ctx *gin.Context) {
	// get user
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user id)")
		return
	}

	// delete user
	err = models.DeleteUser(userID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// delete session
	helpers.DeleteUserID(ctx)

	// render
	ctx.Status(http.StatusNoContent)
}
