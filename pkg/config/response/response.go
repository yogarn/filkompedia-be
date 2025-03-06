package response

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Code    int         `json:"code"`
}

func Success(ctx *fiber.Ctx, code int, data interface{}) {
	ctx.Status(code).JSON(data)
}

func Error(ctx *fiber.Ctx, code int, message string, err error) {
	response := ErrorResponse{
		Message: message,
		Error:   err.Error(),
		Code:    code,
	}

	ctx.Status(code).JSON(response)
}
