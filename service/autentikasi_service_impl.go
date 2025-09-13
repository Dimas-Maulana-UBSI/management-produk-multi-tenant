package service

import (
	"context"
	"database/sql"
	"errors"
	"management-produk/helper"
	"management-produk/model/domain"
	"management-produk/model/web"
	"management-produk/repository"
)

type AutentikasiServiceImpl struct {
	DB *sql.DB
	AutentikasiRepository repository.AutentikasiRepository
}

func NewAutentikasiService(db *sql.DB,repository repository.AutentikasiRepository) AutentikasiService {
	return &AutentikasiServiceImpl{
		DB: db,
		AutentikasiRepository: repository,
	}
}

func(service AutentikasiServiceImpl)Login(ctx context.Context,request web.LoginRequest)(web.LoginResponse,error){
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	loginRequest := domain.Tenant{
		Name: request.Name,
	}
	row,err := service.AutentikasiRepository.GetUser(ctx,tx,loginRequest)
	if err != nil {
		tx.Rollback()
		return web.LoginResponse{}, err
	}
	if request.Password != row.DBPassword{
		return web.LoginResponse{},errors.New("password salah")
	}

	return helper.ToLoginResponse(row),nil

}

func(service AutentikasiServiceImpl)Registrasi(ctx context.Context,request web.RegistrasiRequest)(web.RegistrasiResponse,error){
	tx,err := service.DB.Begin()
	helper.PanicIfError(err)
	tenantRequest := domain.Tenant{
		Name: request.Name,
		DBHost: "localhost",
		DBPort: 3306,
		ApiKey: helper.GenerateApiKey(10),
		DBName: ("db_"+request.Name),
		DBUser: request.Name,
		DBPassword: request.Password,
		Status: "active",
	}
	tenant, err := service.AutentikasiRepository.CreateUser(ctx, tx, tenantRequest)
	if err != nil {
		tx.Rollback()
		return web.RegistrasiResponse{}, err
	}
	_, err = service.AutentikasiRepository.CreateDb(ctx, tx, tenant)
	if err != nil {
		tx.Rollback()
		return web.RegistrasiResponse{}, err
	}
	tx.Commit()

	return helper.ToRegistrasiResponse(tenant),err
}