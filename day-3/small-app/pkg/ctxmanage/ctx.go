package ctxmanage

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"small-app/middlewares"
)

func GetTraceIdOfRequest(c *gin.Context) string {
	ctx := c.Request.Context()

	// ok is false if the type assertion was not successful
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		slog.Error("trace id not present in the context")
		traceId = "Unknown"
	}
	return traceId
}
