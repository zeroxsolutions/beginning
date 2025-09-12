package middleware

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// loggerKey is the context key for storing the logger
type loggerKey struct{}

// WithLogger returns a new context with the logger
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// LoggerFromContext returns the logger from the context, or the default logger if not found
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

func NewLoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		spanCtx := trace.SpanFromContext(c.Request.Context()).SpanContext()

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		l := logger.With(
			slog.String("http.method", c.Request.Method),
			slog.String("url.path", path),
			slog.String("client.ip", c.ClientIP()),
			slog.String("user_agent", c.Request.UserAgent()),
		)
		if spanCtx.IsValid() {
			l = l.With(
				slog.String("trace_id", spanCtx.TraceID().String()),
				slog.String("span_id", spanCtx.SpanID().String()),
			)
		}

		// put logger on request context so handlers can use LoggerFromContext
		ctx := WithLogger(c.Request.Context(), l)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		l.Info("http_request",
			slog.Int("http.status", c.Writer.Status()),
			slog.Int("http.response_bytes", c.Writer.Size()),
			slog.Duration("duration", time.Since(start)),
		)
	}
}
