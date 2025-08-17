package controller

import (
	"fmt"
	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/service"

	"github.com/gofiber/fiber/v2"
)

type AutentikasiControllerImpl struct {
	AutentikasiService service.AutentikasiService
}

func NewAutentikasiController(service service.AutentikasiService)AutentikasiController{
	return &AutentikasiControllerImpl{
		AutentikasiService: service,
	}
}

func (c *AutentikasiControllerImpl) LoginView(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"Title": "Login Page",
	})
}

func(controller *AutentikasiControllerImpl)Login(ctx *fiber.Ctx)error{
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")
	request := web.LoginRequest{
		Name: name,
		Password: password,
	}
	response,err := controller.AutentikasiService.Login(ctx.Context(),request)
	helper.PanicIfError(err)
	fmt.Println(response)
	return ctx.Render("home",fiber.Map{
		"name":response.Name,
		"api_key" : response.ApiKey,
	})
}

func(controller *AutentikasiControllerImpl)RegisterView(ctx *fiber.Ctx)error{
	return ctx.Render("register",fiber.Map{})
}

func(controller *AutentikasiControllerImpl)Register(ctx *fiber.Ctx)error{
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")
	request := web.RegistrasiRequest{
		Name: name,
		Password: password,
	}
	response,err:= controller.AutentikasiService.Registrasi(ctx.Context(),request)
	helper.PanicIfError(err)
	fmt.Println(response)
	return ctx.SendString("registrasi berhasil")
}