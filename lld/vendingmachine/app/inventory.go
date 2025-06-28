package app

type Inventory struct {
	products map[*Product]int
}

func (in *Inventory) AddProduct(product *Product, quantity int) {
	in.products[product] = quantity
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (in *Inventory) RemoveProduct(product *Product) {
	if quantity, ok := in.products[product]; ok {
		in.products[product] = quantity - 1
	}
}

func (in *Inventory) IsAvailable(product *Product) bool {
	if quantity, ok := in.products[product]; ok {
		return quantity > 0
	}
	return false
}
