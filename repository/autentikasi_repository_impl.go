package repository

import (
	"context"
	"database/sql"
	"fmt"
	"management-produk/helper"
	"management-produk/model/domain"
)

type AutentikasiRepositoryImpl struct {
}

func NewAutentikasiRepository() AutentikasiRepository {
	return &AutentikasiRepositoryImpl{}
}
func (repository *AutentikasiRepositoryImpl) CreateDb(ctx context.Context, tx *sql.Tx, tenant domain.Tenant) (domain.Tenant, error) {

	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", tenant.DBName)
	_, err := tx.ExecContext(ctx, createDBSQL)
	if err != nil {
		return domain.Tenant{}, helper.Internal("create database failed", err)
	}

	createTableSQL := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.produk (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		harga int(11) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`, tenant.DBName)

	_, err = tx.ExecContext(ctx, createTableSQL)
	if err != nil {
		return domain.Tenant{}, helper.Internal("create table failed", err)
	}

	return tenant, nil
}

func (repository *AutentikasiRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, tenant domain.Tenant) (domain.Tenant, error) {
	sql := "insert into tenants (name,api_key,db_host,db_port,db_name,db_user,db_password,status) values (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, sql, tenant.Name, tenant.ApiKey, tenant.DBHost, tenant.DBPort, tenant.DBName, tenant.DBUser, tenant.DBPassword, tenant.Status)
	if err != nil {
		return domain.Tenant{}, helper.Internal("insert tenant failed", err)
	}
	return tenant, nil
}

func (repository *AutentikasiRepositoryImpl) GetUser(ctx context.Context, tx *sql.Tx, tenant domain.Tenant) (domain.Tenant, error) {
	sql := "select name,api_key,db_host,db_port,db_name,db_password from tenants where name = ?"
	rows, err := tx.QueryContext(ctx, sql, tenant.Name)
	if err != nil {
		return domain.Tenant{}, helper.Internal("query tenant failed", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&tenant.Name, &tenant.ApiKey, &tenant.DBHost, &tenant.DBPort, &tenant.DBName, &tenant.DBPassword)
		if err != nil {
			return domain.Tenant{}, helper.Internal("scan tenant failed", err)
		}
		return tenant, nil
	}
	return domain.Tenant{}, helper.NotFound("tenant not found")
}
