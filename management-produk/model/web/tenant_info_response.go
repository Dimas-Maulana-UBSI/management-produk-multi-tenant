package web

type TenantInfoResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DBName    string `json:"db_name"`
	APIKey    string `json:"api_key"`
}
