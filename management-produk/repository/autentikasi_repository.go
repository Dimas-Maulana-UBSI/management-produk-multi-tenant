package repository

import (
	"context"
	"database/sql"
	"management-produk/model/domain"
)

type AutentikasiRepository interface {
	CreateDb(ctx context.Context,tx *sql.Tx, tenant domain.Tenant) (domain.Tenant,error)
	CreateUser(ctx context.Context,tx *sql.Tx, tenant domain.Tenant)(domain.Tenant,error)
	GetUser(ctx context.Context,tx *sql.Tx,tenant domain.Tenant)(domain.Tenant,error)
}