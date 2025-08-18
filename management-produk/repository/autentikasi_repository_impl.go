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
	// 1. Buat database tenant
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", tenant.DBName)
	_, err := tx.ExecContext(ctx, createDBSQL)
	helper.PanicIfError(err)

	// 2. Buat tabel di database tenant
	createTableSQL := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s.produk (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		harga int(11) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`, tenant.DBName)

	_, err = tx.ExecContext(ctx, createTableSQL)
	helper.PanicIfError(err)

	return tenant, nil
}

func(repository *AutentikasiRepositoryImpl)CreateUser(ctx context.Context,tx *sql.Tx,tenant domain.Tenant)(domain.Tenant,error){
	sql := "insert into tenants (name,api_key,db_host,db_port,db_name,db_user,db_password,status) values (?,?,?,?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,sql,tenant.Name,tenant.ApiKey,tenant.DBHost,tenant.DBPort,tenant.DBName,tenant.DBUser,tenant.DBPassword,tenant.Status)
	helper.PanicIfError(err)
	return tenant,nil
}


func(repository *AutentikasiRepositoryImpl)GetUser(ctx context.Context,tx *sql.Tx,tenant domain.Tenant)(domain.Tenant,error){
	sql := "select name,api_key,db_host,db_port,db_name,db_password from tenants where name = ?"
	row,err := tx.QueryContext(ctx,sql,tenant.Name)
	helper.PanicIfError(err)
	defer row.Close()
	if row.Next(){
		err := row.Scan(&tenant.Name,&tenant.ApiKey,&tenant.DBHost,&tenant.DBPort,&tenant.DBName,&tenant.DBPassword)
		if err != nil {
			return tenant,err
		}
	}else{
		return tenant, fmt.Errorf("tenant not found")
	}
	return tenant,nil
}
