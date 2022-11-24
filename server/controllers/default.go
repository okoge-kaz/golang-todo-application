package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotImplemented renders error.html with 501 Not Implemented
func NotImplemented(ctx *gin.Context) {
	msg := fmt.Sprintf("%s access to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
	ctx.Header("Cache-Control", "no-cache")
	ctx.String(http.StatusNotImplemented, msg)
}
