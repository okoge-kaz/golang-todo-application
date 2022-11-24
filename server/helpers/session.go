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
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return -1 // not logged in
	}
	return userID.(int)
}

func SetUserID(ctx *gin.Context, userID int) error {
	sessions.Default(ctx).Set(sessionKey, userID)
	err := sessions.Default(ctx).Save()

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return err // can't save session (正常にログインができなかった)
	}
	return nil // logged in(正常にログインできた)
}
