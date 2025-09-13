// model/domain/tenant.go
package domain

type Tenant struct {
    Id         int64
    Name       string
    ApiKey     string
    DBHost     string
    DBPort     int
    DBName     string
    DBUser     string
    DBPassword string
    Status     string
}
