package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get ID from query parameter
func GetIDFromQueryParameter(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return 0, err
	}
	return id, nil
}
