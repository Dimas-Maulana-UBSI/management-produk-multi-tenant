package service

import (
	"context"
	"management-produk/model/web"
)

type TenantService interface {
	// Define methods that the TenantService should implement
	GetInfoTenant(ctx context.Context,apiKey string) (web.TenantInfoResponse, error)
}