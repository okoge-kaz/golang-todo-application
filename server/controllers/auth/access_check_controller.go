package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
	"github.com/okoge-kaz/golang-todo-application/server/models"
)

const frontEndURL string = "http://localhost:3000"

// user login check
func LoginCheck(ctx *gin.Context) {
	switch helpers.GetUserID(ctx) {
	case 0: // not login
		ctx.Redirect(http.StatusFound, frontEndURL+"/login") // フロントエンド側にリダイレクトさせる
		ctx.Abort()
	default: // login
		ctx.Next()
	}
}

// task owner check
func TaskOwnerCheck(ctx *gin.Context) {
	// get task_id
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if ownershipCheck(ctx, taskID) {
		ctx.Next()
	} else {
		ctx.Redirect(http.StatusFound, frontEndURL+"/task") // フロントエンド側にリダイレクトさせる
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}

// ownership check
func ownershipCheck(ctx *gin.Context, taskID int) bool {
	// get user_id from session
	userID := helpers.GetUserID(ctx)
	if userID == -1 {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return false
	}

	// check ownership of task
	_, err := models.GetOwnership(userID, taskID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return false
	}

	return true
}
