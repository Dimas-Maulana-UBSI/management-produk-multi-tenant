package repository

import (
	"context"
	"database/sql"
	"management-produk/helper"
	"management-produk/model/domain"
)

type TenantRepositoryImpl struct {
}

func NewTenantInfoRepository() TenantRepository {
	return &TenantRepositoryImpl{}
}

func (repository *TenantRepositoryImpl) GetInfoTenant(ctx context.Context, db *sql.DB, apiKey string) (*domain.Tenant, error) {
	sql := "SELECT id, name, api_key, db_host, db_port, db_name, db_user, db_password, status FROM tenants WHERE api_key = ?"
	rows, err := db.QueryContext(ctx, sql, apiKey)
	if err != nil {
		return nil, helper.Internal("query tenant failed", err)
	}
	defer rows.Close()

	tenant := &domain.Tenant{}
	if rows.Next() {
		err = rows.Scan(&tenant.Id, &tenant.Name, &tenant.ApiKey, &tenant.DBHost, &tenant.DBPort, &tenant.DBName, &tenant.DBUser, &tenant.DBPassword, &tenant.Status)
		if err != nil {
			return nil, helper.Internal("scan tenant failed", err)
		}
		return tenant, nil
	}
	return nil, helper.NotFound("tenant tidak ditemukan")
}
