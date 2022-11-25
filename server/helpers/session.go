package helpers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionKey string = "user_id"

func SetSession(engine *gin.Engine) {
	store := cookie.NewStore([]byte("secret-key"))
	engine.Use(sessions.Sessions("user-session", store))
}

func GetUserID(ctx *gin.Context) int {
	userID := sessions.Default(ctx).Get(sessionKey)
	if userID == nil {
		ctx.Redirect(http.StatusFound, "http://localhost:3000/login")
		return -1 // not logged in
	}
	return userID.(int)
}

func SetUserID(ctx *gin.Context, userID uint) error {
	sessions.Default(ctx).Set(sessionKey, int(userID))
	err := sessions.Default(ctx).Save()

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return err // can't save session (正常にログインができなかった)
	}
	return nil // logged in(正常にログインできた)
}

func DeleteUserID(ctx *gin.Context) error {
	sessions.Default(ctx).Delete(sessionKey)
	err := sessions.Default(ctx).Save()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return err // can't save session (正常にログアウトできなかった)
	}
	sessions.Default(ctx).Clear()
	sessions.Default(ctx).Options(sessions.Options{MaxAge: -1})
	sessions.Default(ctx).Save()
	return nil // logged out(正常にログアウトできた)
}
