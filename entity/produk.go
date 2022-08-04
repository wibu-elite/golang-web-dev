package entity

type Produk struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Produk) StockStatus() string {

	var status string
	if p.Stock > 10 && p.Stock < 50 {
		status = "Stock Terbatas"
	} else if p.Stock < 10 {
		status = "Stock Hampir Habis"
	} else {
		status = "Stock Masih Banyak"
	}
	return status
}

