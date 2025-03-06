package response

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, code int, data interface{}) {
	ctx.Status(code).JSON(data)
}

func Error(ctx *fiber.Ctx, code int, message string, err error) {
	response := ErrorResponse{
		Message: message,
		Code:    code,
	}

	ctx.Status(code).JSON(response)
}
