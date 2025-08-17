package repository

import (
	"context"
	"database/sql"
	"errors"
	"management-produk/helper"
	"management-produk/model/domain"
)
type TenantRepositoryImpl struct {
}

func NewTenantInfoRepository() TenantRepository{
	return &TenantRepositoryImpl{}
}

func (repository *TenantRepositoryImpl) GetInfoTenant(ctx context.Context, db *sql.DB, apiKey string) (*domain.Tenant, error) {
	sql := "SELECT id, name, api_key, db_host, db_port, db_name, db_user, db_password, status FROM tenants WHERE api_key = ?"
	row, err := db.QueryContext(ctx, sql, apiKey)
	defer row.Close()
	helper.PanicIfError(err)
	tenant := &domain.Tenant{}
	if row.Next() {
		err = row.Scan(&tenant.Id, &tenant.Name, &tenant.ApiKey, &tenant.DBHost, &tenant.DBPort, &tenant.DBName, &tenant.DBUser, &tenant.DBPassword, &tenant.Status)
		helper.PanicIfError(err)
		return tenant, nil
	}
	return nil, errors.New("tenant tidak ditemukan")
}