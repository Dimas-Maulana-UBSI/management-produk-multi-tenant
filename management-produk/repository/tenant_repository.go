package repository

import (
	"database/sql"
	"context"
	"management-produk/model/domain"
)

type TenantRepository interface {
	GetInfoTenant(ctx context.Context, db *sql.DB, apiKey string) (*domain.Tenant, error)
}