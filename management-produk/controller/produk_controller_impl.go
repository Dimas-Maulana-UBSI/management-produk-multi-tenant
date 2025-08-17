package controller

import (
	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/service"

	"github.com/gofiber/fiber/v2"
)

type ProdukControllerImpl struct {
	ProdukServcie service.ProdukServcie
}

func NewProdukController(service service.ProdukServcie)ProdukController{
	return &ProdukControllerImpl{
		ProdukServcie: service,
	}
}

func(controller *ProdukControllerImpl)GetProduk(ctx *fiber.Ctx)error{
	response,err := controller.ProdukServcie.GetAllProduk(ctx.Context())
	helper.PanicIfError(err)
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: response,
	})
}

func(controller *ProdukControllerImpl)CreateProduk(ctx *fiber.Ctx)error{
	var produk web.ProdukRequest
	err := ctx.BodyParser(&produk)
	helper.PanicIfError(err)
	response,err := controller.ProdukServcie.CreateProduk(ctx.Context(),produk)
	helper.PanicIfError(err)
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: response,
	})
}