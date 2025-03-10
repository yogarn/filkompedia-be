package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) GetUserProfile(ctx *fiber.Ctx) (err error) {
	param := ctx.Params("userId")
	userId, err := uuid.Parse(param)
	if err != nil {
		return err
	}

	var profile model.Profile
	if err := r.service.UserService.GetProfile(&profile, userId); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, profile)
	return nil
}

func (r *Rest) GetAllUserProfile(ctx *fiber.Ctx) (err error) {
	var profilesReq model.ProfilesReq
	if err := ctx.BodyParser(&profilesReq); err != nil {
		return err
	}

	var profiles []model.Profile
	if err := r.service.UserService.GetProfiles(&profiles, profilesReq); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, profiles)
	return nil
}
