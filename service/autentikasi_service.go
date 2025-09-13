package service

import (
	"context"
	"management-produk/model/web"
)

type AutentikasiService interface {
	Registrasi(ctx context.Context,request web.RegistrasiRequest)(web.RegistrasiResponse,error)
	Login(ctx context.Context,request web.LoginRequest)(web.LoginResponse,error)
}