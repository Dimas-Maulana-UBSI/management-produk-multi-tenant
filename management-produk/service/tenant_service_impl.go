package service

import (
	"context"
	"database/sql"

	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/repository"
)

type TenantServiceImpl struct {
	tenantRepository repository.TenantRepository
	DB *sql.DB
}

func NewTenantInfoService(repo repository.TenantRepository, db *sql.DB) TenantService {
	return &TenantServiceImpl{
		tenantRepository: repo,
		DB:                db,
	}
}


func(repository *TenantServiceImpl) GetInfoTenant (ctx context.Context,apiKey string) (web.TenantInfoResponse, error) {
	db := repository.DB
	response,err := repository.tenantRepository.GetInfoTenant(ctx, db, apiKey)
	helper.PanicIfError(err)
	return helper.ToTenantInfoResponse(*response), nil
}