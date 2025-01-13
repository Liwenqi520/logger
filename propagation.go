package logger

import (
	"context"
	"net/http"

	"brewing-sass/gopkg/logger/propagation/extract"
	"brewing-sass/gopkg/logger/propagation/inject"
	"github.com/gin-gonic/gin"
)

// HTTPInject inject spanContext
func HttpInject(ctx context.Context, request *http.Request) error {
	return inject.HttpInject(ctx, request)
}

// GinMiddleware extract spanContext
func GinMiddleware(service string) gin.HandlerFunc {
	return extract.GinMiddleware(service)
}
