package service

import (
	"context"
	"database/sql"
	"management-produk/helper"
	"management-produk/model/domain"
	"management-produk/model/web"
	"management-produk/repository"
	"os"
	"strconv"
)

type AutentikasiServiceImpl struct {
	DB                    *sql.DB
	AutentikasiRepository repository.AutentikasiRepository
}

func NewAutentikasiService(db *sql.DB, repository repository.AutentikasiRepository) AutentikasiService {
	return &AutentikasiServiceImpl{
		DB:                    db,
		AutentikasiRepository: repository,
	}
}

func (service AutentikasiServiceImpl) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, helper.Internal("gagal memulai transaksi", err)
	}
	defer tx.Rollback()

	loginRequest := domain.Tenant{
		Name: request.Name,
	}
	row, err := service.AutentikasiRepository.GetUser(ctx, tx, loginRequest)
	if err != nil {
		return web.LoginResponse{}, helper.ToAppError(err)
	}
	if request.Password != row.DBPassword {
		return web.LoginResponse{}, helper.Unauthorized("password salah")
	}

	return helper.ToLoginResponse(row), nil

}

func (service AutentikasiServiceImpl) Registrasi(ctx context.Context, request web.RegistrasiRequest) (web.RegistrasiResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.RegistrasiResponse{}, helper.Internal("gagal memulai transaksi", err)
	}
	defer tx.Rollback()
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	tenantRequest := domain.Tenant{
		Name:       request.Name,
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		ApiKey:     helper.GenerateApiKey(10),
		DBName:     ("db_" + request.Name),
		DBUser:     request.Name,
		DBPassword: request.Password,
		Status:     "active",
	}
	tenant, err := service.AutentikasiRepository.CreateUser(ctx, tx, tenantRequest)
	if err != nil {
		return web.RegistrasiResponse{}, helper.ToAppError(err)
	}
	_, err = service.AutentikasiRepository.CreateDb(ctx, tx, tenant)
	if err != nil {
		return web.RegistrasiResponse{}, helper.ToAppError(err)
	}
	if err := tx.Commit(); err != nil {
		return web.RegistrasiResponse{}, helper.Internal("gagal menyelesaikan transaksi", err)
	}

	return helper.ToRegistrasiResponse(tenant), nil
}
