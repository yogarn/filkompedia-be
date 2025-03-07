package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) Register(ctx *fiber.Ctx) (err error) {
	registerReq := &model.RegisterReq{}
	if err := ctx.BodyParser(registerReq); err != nil {
		return err
	}

	user, err := r.service.AuthService.Register(registerReq)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusCreated, user)
	return nil
}
