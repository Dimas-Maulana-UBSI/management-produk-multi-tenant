package controller

import (
	"fmt"
	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/service"
	"strconv"

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
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _  := strconv.Atoi(ctx.Query("page"))
	response,err := controller.ProdukServcie.GetAllProduk(ctx.UserContext(),limit,page)
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
	response,err := controller.ProdukServcie.CreateProduk(ctx.UserContext(),produk)
	helper.PanicIfError(err)
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: response,
	})
}

func(controller *ProdukControllerImpl)GetById(ctx *fiber.Ctx)error{
	idProduk,_ := strconv.Atoi(ctx.Params("idProduk"))
	fmt.Println(idProduk)
	response,err := controller.ProdukServcie.GetById(ctx.UserContext(),idProduk)
	if err != nil {
		return ctx.JSON(web.WebResponse{Status: 404,
		Message: err.Error(),
		Data: ""})
	}
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: response,
	})
}

func(controller *ProdukControllerImpl)Delete(ctx *fiber.Ctx)error{
	idProduk,_ := strconv.Atoi(ctx.Params("idProduk"))
	err := controller.ProdukServcie.Delete(ctx.UserContext(),idProduk)
	if err != nil {
		return ctx.JSON(web.WebResponse{
			Status: 404,
			Message: err.Error(),
			Data: "",
		})
	}
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: "",
	})
}

func(controller *ProdukControllerImpl)Update(ctx *fiber.Ctx)error{
	var produk web.ProdukRequest
	err := ctx.BodyParser(&produk)
	idProduk,err := strconv.Atoi(ctx.Params("idProduk"))
	response,err := controller.ProdukServcie.Update(ctx.UserContext(),idProduk,produk)
	if err != nil {
		return ctx.JSON(web.WebResponse{
			Status: 404,
			Message: err.Error(),
			Data: "",
		})
	}
	return ctx.JSON(web.WebResponse{
		Status: 200,
		Message: "success",
		Data: response,
	})
}