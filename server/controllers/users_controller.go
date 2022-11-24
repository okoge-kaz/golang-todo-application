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

	// check duplicate user name
	switch models.CheckDuplicateUserName(user.Name) {
	case true: // duplicate
		ctx.String(http.StatusConflict, "Conflict (duplicate user name)")
		// TODO: フロントエンド側で status code 409 に対応した処理を実装すること
		return
	default: // not duplicate
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
}

func ShowUser(ctx *gin.Context) {
	// get user
	userID := helpers.GetUserID(ctx)

	user, err := models.GetUserByID(userID)
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
	// get user_id
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request (invalid user id)")
		return
	}

	// get user
	user, err := models.GetUserByID(userID)

	// get password
	password := ctx.PostForm("password")
	newPassword := ctx.PostForm("new_password")

	// check password
	switch helpers.CheckPasswordHash(user, password) {
	case true:
		// change password
		err = models.ChangeUserPassword(&user, newPassword)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	default:
		ctx.String(http.StatusUnauthorized, "Unauthorized")
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

	// get user
	user, err := models.GetUserByID(userID)

	// get new user name
	newUserName := ctx.PostForm("new_user_name")

	// change name
	err = models.ChangeUserName(&user, newUserName)
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
