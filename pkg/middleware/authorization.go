package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (m *middleware) Authorize(roles []int) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, ok := ctx.Locals("userId").(uuid.UUID)
		if !ok {
			return &response.RoleUnauthorized
		}

		var user entity.User

		err := m.service.UserService.GetUserById(&user, userId)
		if err != nil {
			return err
		}

		for _, role := range roles {
			if user.RoleId == role {
				return ctx.Next()
			}
		}

		return &response.RoleUnauthorized
	}
}
