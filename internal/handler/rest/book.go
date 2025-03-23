package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) SearchBooks(ctx *fiber.Ctx) error {
	var bookSearch model.BookSearch
	bookSearch.Page = ctx.QueryInt("page", 1)
	bookSearch.PageSize = ctx.QueryInt("size", 9)
	bookSearch.SearchParam = ctx.Query("search", "%")

	books, err := r.service.BookService.SearchBooks(bookSearch)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", books)
	return nil
}

func (r *Rest) CreateBook(ctx *fiber.Ctx) error {
	var create model.CreateBook
	if err := ctx.BodyParser(&create); err != nil {
		return err
	}

	if err := r.service.BookService.CreateBook(&create); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}
