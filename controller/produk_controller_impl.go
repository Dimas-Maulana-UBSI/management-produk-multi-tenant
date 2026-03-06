package controller

import (
	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProdukControllerImpl struct {
	ProdukServcie service.ProdukServcie
}

func NewProdukController(service service.ProdukServcie) ProdukController {
	return &ProdukControllerImpl{
		ProdukServcie: service,
	}
}

func (controller *ProdukControllerImpl) GetProduk(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit <= 0 {
		limit = 0 // biarkan service yang handle default
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page <= 0 {
		page = 0 // biarkan service yang handle default
	}

	response, err := controller.ProdukServcie.GetAllProduk(ctx.UserContext(), limit, page)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Status(200).JSON(web.WebResponse{
		Status:  200,
		Message: "success",
		Data:    response,
	})
}

func (controller *ProdukControllerImpl) CreateProduk(ctx *fiber.Ctx) error {
	var produk web.ProdukRequest
	if err := ctx.BodyParser(&produk); err != nil {
		return helper.RespondFiberError(ctx, helper.BadRequest(err.Error()))
	}

	response, err := controller.ProdukServcie.CreateProduk(ctx.UserContext(), produk)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Status(200).JSON(web.WebResponse{
		Status:  200,
		Message: "success",
		Data:    response,
	})
}

func (controller *ProdukControllerImpl) GetById(ctx *fiber.Ctx) error {
	idProduk, err := strconv.Atoi(ctx.Params("idProduk"))
	if err != nil {
		return helper.RespondFiberError(ctx, helper.BadRequest("idProduk harus berupa angka"))
	}

	response, err := controller.ProdukServcie.GetById(ctx.UserContext(), idProduk)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Status(200).JSON(web.WebResponse{
		Status:  200,
		Message: "success",
		Data:    response,
	})
}

func (controller *ProdukControllerImpl) Delete(ctx *fiber.Ctx) error {
	idProduk, err := strconv.Atoi(ctx.Params("idProduk"))
	if err != nil {
		return helper.RespondFiberError(ctx, helper.BadRequest("idProduk harus berupa angka"))
	}

	err = controller.ProdukServcie.Delete(ctx.UserContext(), idProduk)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Status(200).JSON(web.WebResponse{
		Status:  200,
		Message: "success",
		Data:    "",
	})
}

func (controller *ProdukControllerImpl) Update(ctx *fiber.Ctx) error {
	var produk web.ProdukRequest
	if err := ctx.BodyParser(&produk); err != nil {
		return helper.RespondFiberError(ctx, helper.BadRequest(err.Error()))
	}

	idProduk, err := strconv.Atoi(ctx.Params("idProduk"))
	if err != nil {
		return helper.RespondFiberError(ctx, helper.BadRequest("idProduk harus berupa angka"))
	}

	response, err := controller.ProdukServcie.Update(ctx.UserContext(), idProduk, produk)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Status(200).JSON(web.WebResponse{
		Status:  200,
		Message: "success",
		Data:    response,
	})
}
