package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (m *middleware) LogrusMiddleware(ctx *fiber.Ctx) error {
	start := ctx.Context().Time()

	err := ctx.Next()
	statusCode, _ := response.GetErrorInfo(err)

	entry := m.logger.WithFields(logrus.Fields{
		"method":    ctx.Method(),
		"path":      ctx.Path(),
		"status":    statusCode,
		"latency":   ctx.Context().Time().Sub(start),
		"ip":        ctx.IP(),
		"userAgent": ctx.Get("User-Agent"),
	})

	if statusCode >= 200 && statusCode < 300 {
		entry.Info("incoming request")
	} else {
		entry.Error("error request")
	}

	return err
}
