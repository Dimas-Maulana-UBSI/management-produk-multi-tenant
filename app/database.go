package app

import (
	"database/sql"
	"management-produk/helper"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewMasterDB() *sql.DB {
    dsn := "root@tcp(localhost:3306)/tenants?parseTime=true"

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    db.SetMaxIdleConns(5)
    db.SetMaxOpenConns(20)
    db.SetConnMaxLifetime(60 * time.Minute)
    db.SetConnMaxIdleTime(10 * time.Minute)

    if err := db.Ping(); err != nil {
        panic(err)
    }

    return db
}

func NewTenantDb() *sql.DB {
    dsn := "root@tcp(localhost:3306)/?parseTime=true"
    db,err := sql.Open("mysql",dsn)
    helper.PanicIfError(err)

    db.SetMaxIdleConns(20)
    db.SetMaxOpenConns(30)
    db.SetConnMaxLifetime(60 * time.Minute)
    db.SetConnMaxIdleTime(10 * time.Minute)

    if err := db.Ping(); err != nil {
        panic(err)
    }

    return db
}
