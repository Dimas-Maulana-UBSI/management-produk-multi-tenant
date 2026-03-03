package web

type ProdukResponse struct {
	Id int
	Name string
	Harga int

}

type ProdukPaginationResponse struct {
	Data []ProdukResponse
	Page int
	Limit int
	TotalData int
	TotalPage int
}