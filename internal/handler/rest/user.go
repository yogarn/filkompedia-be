package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) GetMe(ctx *fiber.Ctx) (err error) {
	userId, ok := ctx.Locals("userId").(uuid.UUID)
	if !ok {
		return &response.Unauthorized
	}

	var profile model.Profile
	if err := r.service.UserService.GetProfile(&profile, userId); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", profile)
	return nil
}

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

	response.Success(ctx, http.StatusOK, "success", profile)
	return nil
}

func (r *Rest) GetAllUserProfile(ctx *fiber.Ctx) (err error) {
	var userProfile model.ProfilesReq
	userProfile.Page = ctx.QueryInt("page", 1)
	userProfile.PageSize = ctx.QueryInt("size", 9)

	var profiles []model.Profile
	if err := r.service.UserService.GetProfiles(&profiles, userProfile); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", profiles)
	return nil
}
