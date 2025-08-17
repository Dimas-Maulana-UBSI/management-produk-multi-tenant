package helper
import (
	"management-produk/model/domain"
	"management-produk/model/web"
)

func ToTenantInfoResponse(tenant domain.Tenant) web.TenantInfoResponse {
	return web.TenantInfoResponse{
		ID:          int(tenant.Id),
		Name:        tenant.Name,
		DBName:     tenant.DBName,
		APIKey:     tenant.ApiKey,
	}
}

func ToRegistrasiResponse(tenant domain.Tenant)web.RegistrasiResponse{
	return web.RegistrasiResponse{
		Name: tenant.Name,
		ApiKey: tenant.ApiKey,
		Status: tenant.Status,
	}
}

func ToLoginResponse(tenant domain.Tenant)web.LoginResponse{
	return web.LoginResponse{
		Name: tenant.Name,
		ApiKey: tenant.ApiKey,
	}
}

func ToProdukResponse(produk domain.Produk)web.ProdukResponse{
	return web.ProdukResponse{
		Id: produk.Id,
		Name: produk.Nama,
		Harga: produk.Harga,
	}
}