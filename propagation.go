package logger

import (
	"context"
	"net/http"

	// "propagation/extract"
	// "propagation/inject"
	"logger/propagation"
	"github.com/gin-gonic/gin"
)

// HTTPInject inject spanContext
func HttpInject(ctx context.Context, request *http.Request) error {
	return propagation.inject.HttpInject(ctx, request)
}

// GinMiddleware extract spanContext
func GinMiddleware(service string) gin.HandlerFunc {
	return propagation.extract.GinMiddleware(service)
}
