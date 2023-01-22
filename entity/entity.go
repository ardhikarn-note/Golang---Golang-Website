package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	if p.Stock < 3 {
		return "Stock is running out"
	}
	return "In Stock"
}
