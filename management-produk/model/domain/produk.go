package domain

import "time"

type Produk struct {
	Id         int
	Nama       string
	Harga      int
	Created_at time.Time
}