package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
)

// user login check
func LoginCheck(ctx *gin.Context) {
	switch helpers.GetUserID(ctx) {
	case 0: // not login
		ctx.Redirect(http.StatusFound, "http://localhost:3000/login") // フロントエンド側にリダイレクトさせる
		ctx.Abort()
	default: // login
		ctx.Next()
	}
}
