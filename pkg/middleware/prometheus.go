package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"

	monitoring "github.com/yogarn/filkompedia-be/pkg/prometheus"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func PromMiddleware(reg monitoring.Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		now := time.Now()

		err := c.Next()

		// somehow this deep copy fix data race issues with fasthttp
		// do not remove
		method := string(append([]byte{}, c.Method()...))
		path := string(append([]byte{}, c.Path()...))

		respStatus := c.Response().StatusCode()
		duration := time.Since(now).Seconds()

		if err != nil {
			respStatus = fiber.StatusInternalServerError
			var errorRequest *response.ErrorResponse
			if errors.As(err, &errorRequest) {
				respStatus = errorRequest.Code
			}

			var fiberError *fiber.Error
			if errors.As(err, &fiberError) {
				respStatus = fiberError.Code
			}

			var validationError validator.ValidationErrors
			if errors.As(err, &validationError) {
				respStatus = fiber.StatusBadRequest
			}
		}

		reg.Duration.WithLabelValues(strconv.Itoa(respStatus), method, path).Observe(duration)
		reg.RequestTotal.With(prometheus.Labels{"response_code": strconv.Itoa(respStatus), "method": method}).Inc()
		reg.DurationSummary.Observe(duration)

		return err
	}
}
