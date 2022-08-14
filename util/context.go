package util

import (
	"context"
	"math/rand"

	"github.com/gin-gonic/gin"
	"waetherproject.com/utilities"
)

func SetContext(c *gin.Context) context.Context {
	logId := rand.Int()
	ctx := context.WithValue(c, utilities.LogId, logId)
	return ctx
}
